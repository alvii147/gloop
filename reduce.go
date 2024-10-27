package gloop

import "iter"

// ReduceOptions defines configurable options for [Reduce] and
// [Reduce2].
type ReduceOptions[A any] struct {
	// InitialValue is the reduction's starting value.
	InitialValue *A
}

// ReduceOptionFunc is the function signature of configuration helpers
// for [Reduce] and [Reduce2].
type ReduceOptionFunc[A any] func(*ReduceOptions[A])

// WithReduceInitialValue is a helper for configuring initial value for
// [Reduce] and [Reduce2].
func WithReduceInitialValue[A any](initialValue A) ReduceOptionFunc[A] {
	return func(o *ReduceOptions[A]) {
		o.InitialValue = &initialValue
	}
}

// ReduceFunc is the function signature of the reduction function used
// in [Reduce].
type ReduceFunc[A, V any] func(A, V) A

// Reduce runs a given function on each value from an iter.Seq sequence
// and accumulates the result into a single value.
func Reduce[A, V any](
	seq iter.Seq[V],
	f ReduceFunc[A, V],
	opts ...ReduceOptionFunc[A],
) A {
	return Reduce2(Enumerate(seq), func(acc A, _ int, value V) A {
		return f(acc, value)
	}, opts...)
}

// Reduce2Func is the function signature of the reduction function used
// in [Reduce2].
type Reduce2Func[A, K, V any] func(A, K, V) A

// Reduce2 runs a given function on each value from an iter.Seq2
// sequence and accumulates the result into a single value.
func Reduce2[A, K, V any](
	seq iter.Seq2[K, V],
	f Reduce2Func[A, K, V],
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

	for key, value := range seq {
		acc = f(acc, key, value)
	}

	return acc
}
