package main

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/goph/emperror"
	"github.com/goph/fxt"
	fxsql "github.com/goph/fxt/database/sql"
	fxdebug "github.com/goph/fxt/debug"
	fxerrors "github.com/goph/fxt/errors"
	fxgrpc "github.com/goph/fxt/grpc"
	fxlog "github.com/goph/fxt/log"
	fxprometheus "github.com/goph/fxt/metrics/prometheus"
	fxtracing "github.com/goph/fxt/tracing"
	"github.com/goph/healthz"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"go.uber.org/fx"
)

func main() {
	status := healthz.NewStatusChecker(healthz.Healthy)
	var ext struct {
		Config       *Config
		Closer       fxt.Closer
		Logger       log.Logger
		ErrorHandler emperror.Handler

		DebugErr fxdebug.Err
		GrpcErr  fxgrpc.Err
	}

	app := fx.New(
		fx.NopLogger,
		fxt.Bootstrap,
		fx.Provide(
			NewConfig,

			// Log and error handling
			NewLoggerConfig,
			fxlog.NewLogger,
			fxerrors.NewHandler,

			// Debug server
			NewDebugConfig,
			fxdebug.NewServer,
			fxdebug.NewHealthCollector,
		),
		fx.Invoke(func(collector healthz.Collector) {
			collector.RegisterChecker(healthz.ReadinessCheck, status)
		}),
		fx.Extract(&ext),

		fx.Provide(
			// gRPC server
			NewService,
			NewGrpcConfig,
			fxgrpc.NewServer,

			fxtracing.NewTracer,

			// Database
			NewDatabaseConfig,
			fxsql.NewConnection,
		),

		User(),

		// Make sure to register this invoke function as the last,
		// so all registered gRPC services are exposed in metrics.
		fx.Invoke(grpc_prometheus.Register, fxprometheus.RegisterHandler),
	)

	err := app.Err()
	if err != nil {
		panic(err)
	}

	// Close resources when the application stops running
	defer ext.Closer.Close()

	// Register error handler to recover from panics
	defer emperror.HandleRecover(ext.ErrorHandler)

	level.Info(ext.Logger).Log(
		"msg", fmt.Sprintf("starting %s", FriendlyServiceName),
		"version", Version,
		"commit_hash", CommitHash,
		"build_date", BuildDate,
	)

	err = app.Start(context.Background())
	if err != nil {
		panic(err)
	}

	select {
	case sig := <-app.Done():
		level.Info(ext.Logger).Log("msg", fmt.Sprintf("captured %v signal", sig))

	case err := <-ext.DebugErr:
		if err != nil {
			err = errors.Wrap(err, "debug server crashed")
			ext.ErrorHandler.Handle(err)
		}

	case err := <-ext.GrpcErr:
		if err != nil {
			err = errors.Wrap(err, "grpc server crashed")
			ext.ErrorHandler.Handle(err)
		}
	}

	status.SetStatus(healthz.Unhealthy)

	ctx, cancel := context.WithTimeout(context.Background(), ext.Config.ShutdownTimeout)
	defer cancel()

	err = app.Stop(ctx)
	emperror.HandleIfErr(ext.ErrorHandler, err)
}