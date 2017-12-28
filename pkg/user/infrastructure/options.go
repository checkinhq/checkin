package infrastructure

import (
	"github.com/go-kit/kit/log"
	"github.com/goph/clock"
	"github.com/goph/emperror"
)

// Option sets a value in an Options instance.
type Option func(o *Options)

// Clock returns a Option that sets the clock.
func Clock(c clock.Clock) Option {
	return func(o *Options) {
		o.clock = c
	}
}

// Logger returns a Option that sets the logger.
func Logger(l log.Logger) Option {
	return func(o *Options) {
		o.logger = l
	}
}

// ErrorHandler returns a Option that sets the error handler.
func ErrorHandler(h emperror.Handler) Option {
	return func(o *Options) {
		o.errorHandler = h
	}
}

// Options holds a list of options frequently required by different components of the system.
type Options struct {
	clock        clock.Clock
	logger       log.Logger
	errorHandler emperror.Handler
}

// NewOptions returns a new Options instance.
func NewOptions(opts ...Option) *Options {
	o := new(Options)

	for _, opt := range opts {
		opt(o)
	}

	return o
}

// Clock returns a clock.Clock instance.
func (o *Options) Clock() clock.Clock {
	// Default clock
	if o.clock == nil {
		o.clock = clock.SystemClock
	}

	return o.clock
}

// Logger returns a log.Logger instance.
func (o *Options) Logger() log.Logger {
	// Default logger
	if o.logger == nil {
		o.logger = log.NewNopLogger()
	}

	return o.logger
}

// ErrorHandler returns a emperror.Handler instance.
func (o *Options) ErrorHandler() emperror.Handler {
	// Default error handler
	if o.errorHandler == nil {
		o.errorHandler = emperror.NewNopHandler()
	}

	return o.errorHandler
}
