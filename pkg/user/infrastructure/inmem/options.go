package inmem

import (
	"github.com/goph/clock"
	"github.com/hashicorp/go-memdb"
)

// Option sets a value in an Options instance.
type Option func(o *options)

// Clock returns a Option that sets the clock.
func Clock(c clock.Clock) Option {
	return func(o *options) {
		o.clock = c
	}
}

func Database(db *memdb.MemDB) Option {
	return func(o *options) {
		o.db = db
	}
}

// options holds a list of options required by repository implementations.
type options struct {
	clock clock.Clock
	db    *memdb.MemDB
}

// newOptions returns a new options instance.
func newOptions(opts ...Option) *options {
	o := new(options)

	for _, opt := range opts {
		opt(o)
	}

	// Default clock
	if o.clock == nil {
		o.clock = clock.SystemClock
	}

	return o
}
