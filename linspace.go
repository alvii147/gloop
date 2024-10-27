package gloop

import (
	"iter"
)

// LinspaceOptions defines configurable options for [Linspace].
type LinspaceOptions struct {
	// Closed represents whether or not the interval is closed at the
	// stop point.
	Closed bool
}

// LinspaceOptionFunc is the function signature of configuration
// helpers for [Linspace].
type LinspaceOptionFunc func(*LinspaceOptions)

// WithLinspaceClosed is a helper for configuring the interval to be
// closed for [Linspace].
func WithLinspaceClosed(closed bool) LinspaceOptionFunc {
	return func(o *LinspaceOptions) {
		o.Closed = closed
	}
}

// Linspace allows looping over evenly spaced values within a given
// interval. n must be greater than 1.
func Linspace[N Number](
	start N,
	stop N,
	n int,
	opts ...LinspaceOptionFunc,
) iter.Seq[float64] {
	if n <= 1 {
		panic("n must be greater than 1")
	}

	options := LinspaceOptions{
		Closed: false,
	}

	for _, opt := range opts {
		opt(&options)
	}

	fstart := float64(start)
	fstop := float64(stop)
	fstep := (fstop - fstart) / float64(n-1)

	return Interval(
		fstart,
		fstop,
		fstep,
		WithIntervalClosed(options.Closed),
	)
}
