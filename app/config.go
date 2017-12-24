package app

import (
	"time"
)

// Config holds any kind of configuration that comes from the outside world and is necessary for running.
type Config struct {
	// Meaningful values are recommended (eg. production, development, staging, release/123, etc)
	//
	// "development" is treated special: address types will use the loopback interface as default when none is defined.
	// This is useful when developing locally and listening on all interfaces requires elevated rights.
	Environment string `env:"" default:"production"`

	// Turns on some debug functionality: more verbose logs, exposed pprof, expvar and net trace in the debug server.
	Debug bool `env:""`

	// Defines the log format.
	// Valid values are: json, logfmt
	LogFormat string `env:"" split_words:"true" default:"json"`

	// Debug and health check server address
	DebugAddr string `flag:"" split_words:"true" default:":10000" usage:"Debug and health check server address"`

	// Timeout for graceful shutdown
	ShutdownTimeout time.Duration `flag:"" split_words:"true" default:"15s" usage:"Timeout for graceful shutdown"`

	// gRPC server address
	GrpcAddr string `flag:"" split_words:"true" default:":8000" usage:"gRPC service address"`

	// Enable the gRPC reflection service.
	GrpcEnableReflection bool `split_words:"true"`

	// Database connection details
	DbHost string `env:"" split_words:"true" required:"true"`
	DbPort int    `env:"" split_words:"true" default:"3306"`
	DbUser string `env:"" split_words:"true" required:"true"`
	DbPass string `env:"" split_words:"true"` // Required removed for now because empty value is not supported by Viper
	DbName string `env:"" split_words:"true" required:"true"`
}
