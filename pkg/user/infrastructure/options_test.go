package infrastructure_test

import (
	"testing"

	"github.com/checkinhq/checkin/pkg/user/infrastructure"
	"github.com/go-kit/kit/log"
	"github.com/goph/clock"
	"github.com/goph/emperror"
	"github.com/stretchr/testify/assert"
)

func TestOptions_Clock(t *testing.T) {
	c := clock.StoppedAt(clock.SystemClock.Now())

	opts := infrastructure.NewOptions(infrastructure.Clock(c))

	assert.Equal(t, c, opts.Clock())
}

func TestOptions_Logger(t *testing.T) {
	logger := log.NewNopLogger()

	opts := infrastructure.NewOptions(infrastructure.Logger(logger))

	assert.Equal(t, logger, opts.Logger())
}

func TestOptions_ErrorHandler(t *testing.T) {
	errorHandler := emperror.NewNopHandler()

	opts := infrastructure.NewOptions(infrastructure.ErrorHandler(errorHandler))

	assert.Equal(t, errorHandler, opts.ErrorHandler())
}
