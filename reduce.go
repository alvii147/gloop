package gloop

import "iter"

// ReduceOptions defines configurable options for Reduce.
type ReduceOptions[A any] struct {
	// InitialValue is the reduction's starting value.
	InitialValue *A
}

// ReduceOptionFunc is the function signature of configuration helpers
// for Reduce.
type ReduceOptionFunc[A any] func(*ReduceOptions[A])

// WithReduceInitialValue is a helper for configuring initial value for
// Reduce.
func WithReduceInitialValue[A any](initialValue A) ReduceOptionFunc[A] {
	return func(o *ReduceOptions[A]) {
		o.InitialValue = &initialValue
	}
}

// ReduceFunc is the function signature of the reduction function.
type ReduceFunc[A, V any] func(A, V) A

// Reduce runs a given function on each value from a given sequence and
// accumulates the result into a single value.
func Reduce[A, V any](
	seq iter.Seq[V],
	f ReduceFunc[A, V],
	opts ...ReduceOptionFunc[A],
) A {
	options := ReduceOptions[A]{
		InitialValue: nil,
	}

	for _, opt := range opts {
		opt(&options)
	}

	var acc A
	if options.InitialValue != nil {
		acc = *options.InitialValue
	}

	for value := range seq {
		acc = f(acc, value)
	}

	return acc
}
