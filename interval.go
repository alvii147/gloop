package gloop

import "iter"

// IntervalOptions defines configurable options for Interval.
type IntervalOptions struct {
	// Closed represents whether or not the interval is closed at the
	// stop point.
	Closed bool
}

// IntervalOptionFunc is the function signature of configuration
// helpers for Interval.
type IntervalOptionFunc func(*IntervalOptions)

// WithIntervalClosed is a helper for configuring the interval to be
// closed for Interval.
func WithIntervalClosed(closed bool) IntervalOptionFunc {
	return func(o *IntervalOptions) {
		o.Closed = closed
	}
}

// Interval allows looping over values in a given interval with a given
// step size.
func Interval[N Number](
	start N,
	stop N,
	step N,
	opts ...IntervalOptionFunc,
) iter.Seq[N] {
	options := IntervalOptions{
		Closed: false,
	}

	for _, opt := range opts {
		opt(&options)
	}

	if step == 0 {
		return func(yield func(N) bool) {}
	}

	if step > 0 && stop-start < 0 {
		return func(yield func(N) bool) {}
	}

	if step < 0 && stop-start > 0 {
		return func(yield func(N) bool) {}
	}

	var loopCond func(i N, stop N) bool
	switch {
	case step > 0 && options.Closed:
		loopCond = func(i N, stop N) bool {
			return i <= stop
		}
	case step < 0 && options.Closed:
		loopCond = func(i N, stop N) bool {
			return i >= stop
		}
	case step > 0 && !options.Closed:
		loopCond = func(i N, stop N) bool {
			return i < stop
		}
	case step < 0 && !options.Closed:
		loopCond = func(i N, stop N) bool {
			return i > stop
		}
	}

	return func(yield func(N) bool) {
		for i := start; loopCond(i, stop); i += step {
			if !yield(i) {
				return
			}
		}
	}
}
