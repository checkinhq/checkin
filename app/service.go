package app

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goph/fxt"
	fxsql "github.com/goph/fxt/database/sql"
	fxgrpc "github.com/goph/fxt/grpc"
	"github.com/goph/healthz"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"go.uber.org/dig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

// GrpcServerParams provides a set of dependencies for a grpc server constructor.
type GrpcServerParams struct {
	dig.In

	Config          *fxgrpc.Config
	Logger          log.Logger        `optional:"true"`
	HealthCollector healthz.Collector `optional:"true"`
	Lifecycle       fxt.Lifecycle
}

// NewGrpcServer creates a new grpc server.
func NewGrpcServer(params GrpcServerParams) (*grpc.Server, *http.Server, fxgrpc.Err) {
	logger := params.Logger
	if logger == nil {
		logger = log.NewNopLogger()
	}

	logger = log.With(logger, "server", "grpc")

	// TODO: separate log levels
	// TODO: only set logger once
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(
		log.NewStdlibAdapter(level.Info(logger)),
		log.NewStdlibAdapter(level.Warn(logger)),
		log.NewStdlibAdapter(level.Error(logger)),
	))

	if params.HealthCollector != nil {
		params.HealthCollector.RegisterChecker(healthz.ReadinessCheck, healthz.NewTCPChecker(params.Config.Addr))
	}

	server := grpc.NewServer(params.Config.Options...)

	if params.Config.ReflectionEnabled {
		level.Debug(logger).Log("msg", "grpc reflection service enabled")

		reflection.Register(server)
	}

	wrappedServer := grpcweb.WrapServer(server)
	httpServer := &http.Server{
		Handler: wrappedServer,
	}

	errCh := make(chan error, 1)

	params.Lifecycle.Append(fxt.Hook{
		OnStart: func(ctx context.Context) error {
			level.Info(logger).Log(
				"msg", "listening on address",
				"addr", params.Config.Addr,
				"network", params.Config.Network,
			)

			lis, err := net.Listen(params.Config.Network, params.Config.Addr)
			if err != nil {
				return errors.WithStack(err)
			}

			go func() {
				errCh <- httpServer.Serve(lis)
			}()

			return nil
		},
		OnStop:  httpServer.Shutdown,
		OnClose: httpServer.Close,
	})

	return server, httpServer, errCh
}

// NewGrpcConfig creates a grpc config.
func NewGrpcConfig(config Config, tracer opentracing.Tracer) *fxgrpc.Config {
	addr := config.GrpcAddr

	// Listen on loopback interface in development mode
	if config.Environment == "development" && addr[0] == ':' {
		addr = "127.0.0.1" + addr
	}

	c := fxgrpc.NewConfig(addr)
	c.ReflectionEnabled = config.GrpcEnableReflection
	c.Options = []grpc.ServerOption{
		grpc_middleware.WithStreamServerChain(
			grpc_opentracing.StreamServerInterceptor(grpc_opentracing.WithTracer(tracer)),
			grpc_prometheus.StreamServerInterceptor,
			grpc_recovery.StreamServerInterceptor(),
		),
		grpc_middleware.WithUnaryServerChain(
			grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
		),
	}

	return c
}

// NewDatabaseConfig returns a new database connection configuration.
func NewDatabaseConfig(config Config) *fxsql.Config {
	return fxsql.NewConfig(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			config.DbUser,
			config.DbPass,
			config.DbHost,
			config.DbPort,
			config.DbName,
		),
	)
}
