package gloop

import "iter"

// ZipOptions defines configurable options for Zip.
type ZipOptions[V1, V2 any] struct {
	// Padded indicates whether the shorter sequence will be padded. If
	// true, the shorter sequence is padded to match the length of the
	// longer one. If false, the number of iterations is equal to the
	// length of the shorter sequence.
	Padded bool
	// PadValue1 is the value the first sequence is padded with. This
	// is not used if Padded is false or if the first sequence is
	// shorter than the second.
	PadValue1 *V1
	// PadValue2 is the value the second sequence is padded with. This
	// is not used if Padded is false or if the second sequence is
	// shorter than the first.
	PadValue2 *V2
}

// ZipOptionFunc is the function signature of configuration helpers for
// Zip.
type ZipOptionFunc[V1, V2 any] func(*ZipOptions[V1, V2])

// WithZipPadded is a helper for configuring Zip to pad the shorter
// sequence.
func WithZipPadded[V1, V2 any](padded bool) ZipOptionFunc[V1, V2] {
	return func(o *ZipOptions[V1, V2]) {
		o.Padded = true
	}
}

// WithZipPadValue1 is a helper for configuring padded values for the
// first sequence in Zip.
func WithZipPadValue1[V1, V2 any](value V1) ZipOptionFunc[V1, V2] {
	return func(o *ZipOptions[V1, V2]) {
		o.PadValue1 = &value
	}
}

// WithZipPadValue2 is a helper for configuring padded values for the
// second slice in Zip.
func WithZipPadValue2[V1, V2 any](value V2) ZipOptionFunc[V1, V2] {
	return func(o *ZipOptions[V1, V2]) {
		o.PadValue2 = &value
	}
}

// Zip allows a for loop to iterate over two given sequences in pairs.
func Zip[V1, V2 any](
	seq1 iter.Seq[V1],
	seq2 iter.Seq[V2],
	opts ...ZipOptionFunc[V1, V2],
) iter.Seq2[V1, V2] {
	options := ZipOptions[V1, V2]{
		Padded:    false,
		PadValue1: nil,
	}

	for _, opt := range opts {
		opt(&options)
	}

	return func(yield func(V1, V2) bool) {
		next1, stop1 := iter.Pull(seq1)
		defer stop1()

		next2, stop2 := iter.Pull(seq2)
		defer stop2()

		for {
			var value1 V1
			nextValue1, ok1 := next1()
			if ok1 {
				value1 = nextValue1
			} else if options.PadValue1 != nil {
				value1 = *options.PadValue1
			}

			var value2 V2
			nextValue2, ok2 := next2()
			if ok2 {
				value2 = nextValue2
			} else if options.PadValue2 != nil {
				value2 = *options.PadValue2
			}

			if !ok1 && !ok2 {
				return
			}

			if !options.Padded && (!ok1 || !ok2) {
				return
			}

			if !yield(value1, value2) {
				return
			}
		}
	}
}
