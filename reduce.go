package gloop

import "iter"

// FoldOptions defines configurable options for [Fold] and [Fold2].
type FoldOptions[A any] struct {
	// InitialValue is the starting value before folding.
	InitialValue *A
}

// FoldOptionFunc is the function signature of configuration helpers
// for [Fold] and [Fold2].
type FoldOptionFunc[A any] func(*FoldOptions[A])

// WithFoldInitialValue is a helper for configuring initial value for
// [Fold] and [Fold2].
func WithFoldInitialValue[A any](initialValue A) FoldOptionFunc[A] {
	return func(o *FoldOptions[A]) {
		o.InitialValue = &initialValue
	}
}

// FoldFunc is the function signature of the folding function used in
// [Fold].
type FoldFunc[A, V any] func(A, V) A

// Fold runs a given function on each value from an iter.Seq sequence
// and accumulates the result into a single value.
func Fold[A, V any](
	seq iter.Seq[V],
	f FoldFunc[A, V],
	opts ...FoldOptionFunc[A],
) A {
	return Fold2(Enumerate(seq), func(acc A, _ int, value V) A {
		return f(acc, value)
	}, opts...)
}

// Fold2Func is the function signature of the reduction function used
// in [Fold2].
type Fold2Func[A, K, V any] func(A, K, V) A

// Fold2 runs a given function on each value from an iter.Seq2 sequence
// and accumulates the result into a single value.
func Fold2[A, K, V any](
	seq iter.Seq2[K, V],
	f Fold2Func[A, K, V],
	opts ...FoldOptionFunc[A],
) A {
	options := FoldOptions[A]{
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
