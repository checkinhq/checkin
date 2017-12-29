package database

import (
	"testing"

	"github.com/goph/clock"
	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	c := clock.StoppedAt(clock.SystemClock.Now())

	opts := newOptions(Clock(c))

	assert.Equal(t, c, opts.clock)
}
