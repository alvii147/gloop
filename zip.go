package gloop

import (
	"container/list"
	"iter"
)

// ZipOptions defines configurable options for [Zip].
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
// [Zip].
type ZipOptionFunc[V1, V2 any] func(*ZipOptions[V1, V2])

// WithZipPadded is a helper for configuring [Zip] to pad the shorter
// sequence.
func WithZipPadded[V1, V2 any](padded bool) ZipOptionFunc[V1, V2] {
	return func(o *ZipOptions[V1, V2]) {
		o.Padded = padded
	}
}

// WithZipPadValue1 is a helper for configuring padded values for the
// first sequence in [Zip].
func WithZipPadValue1[V1, V2 any](value V1) ZipOptionFunc[V1, V2] {
	return func(o *ZipOptions[V1, V2]) {
		o.PadValue1 = &value
	}
}

// WithZipPadValue2 is a helper for configuring padded values for the
// second slice in [Zip].
func WithZipPadValue2[V1, V2 any](value V2) ZipOptionFunc[V1, V2] {
	return func(o *ZipOptions[V1, V2]) {
		o.PadValue2 = &value
	}
}

// Zip allows looping over two [iter.Seq] sequences in pairs.
func Zip[V1, V2 any](
	seq1 iter.Seq[V1],
	seq2 iter.Seq[V2],
	opts ...ZipOptionFunc[V1, V2],
) iter.Seq2[V1, V2] {
	options := ZipOptions[V1, V2]{
		Padded:    false,
		PadValue1: nil,
		PadValue2: nil,
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

// Zip2Options defines configurable options for [Zip2].
type Zip2Options[K1, V1, K2, V2 any] struct {
	// Padded indicates whether the shorter sequence will be padded. If
	// true, the shorter sequence is padded to match the length of the
	// longer one. If false, the number of iterations is equal to the
	// length of the shorter sequence.
	Padded bool
	// PadKey1 is the key the first sequence is padded with. This is
	// not used if Padded is false or if the first sequence is shorter
	// than the second.
	PadKey1 *K1
	// PadValue1 is the value the first sequence is padded with. This
	// is not used if Padded is false or if the first sequence is
	// shorter than the second.
	PadValue1 *V1
	// PadKey2 is the key the second sequence is padded with. This is
	// not used if Padded is false or if the first sequence is shorter
	// than the second.
	PadKey2 *K2
	// PadValue2 is the value the second sequence is padded with. This
	// is not used if Padded is false or if the second sequence is
	// shorter than the first.
	PadValue2 *V2
}

// Zip2OptionFunc is the function signature of configuration helpers
// for [Zip2].
type Zip2OptionFunc[K1, V1, K2, V2 any] func(*Zip2Options[K1, V1, K2, V2])

// WithZip2Padded is a helper for configuring [Zip2] to pad the shorter
// sequence.
func WithZip2Padded[K1, V1, K2, V2 any](padded bool) Zip2OptionFunc[K1, V1, K2, V2] {
	return func(o *Zip2Options[K1, V1, K2, V2]) {
		o.Padded = padded
	}
}

// WithZip2PadKey1 is a helper for configuring padded keys for the
// first sequence in [Zip2].
func WithZip2PadKey1[K1, V1, K2, V2 any](key K1) Zip2OptionFunc[K1, V1, K2, V2] {
	return func(o *Zip2Options[K1, V1, K2, V2]) {
		o.PadKey1 = &key
	}
}

// WithZip2PadValue1 is a helper for configuring padded values for the
// first sequence in [Zip2].
func WithZip2PadValue1[K1, V1, K2, V2 any](value V1) Zip2OptionFunc[K1, V1, K2, V2] {
	return func(o *Zip2Options[K1, V1, K2, V2]) {
		o.PadValue1 = &value
	}
}

// WithZip2PadKey2 is a helper for configuring padded keys for the
// second sequence in [Zip2].
func WithZip2PadKey2[K1, V1, K2, V2 any](key K2) Zip2OptionFunc[K1, V1, K2, V2] {
	return func(o *Zip2Options[K1, V1, K2, V2]) {
		o.PadKey2 = &key
	}
}

// WithZip2PadValue2 is a helper for configuring padded values for the
// second sequence in [Zip2].
func WithZip2PadValue2[K1, V1, K2, V2 any](value V2) Zip2OptionFunc[K1, V1, K2, V2] {
	return func(o *Zip2Options[K1, V1, K2, V2]) {
		o.PadValue2 = &value
	}
}

// Zip2 allows looping over two [iter.Seq2] sequences in pairs.
func Zip2[K1, V1, K2, V2 any](
	seq1 iter.Seq2[K1, V1],
	seq2 iter.Seq2[K2, V2],
	opts ...Zip2OptionFunc[K1, V1, K2, V2],
) iter.Seq2[KeyValuePair[K1, V1], KeyValuePair[K2, V2]] {
	options := Zip2Options[K1, V1, K2, V2]{
		Padded:    false,
		PadKey1:   nil,
		PadValue1: nil,
		PadKey2:   nil,
		PadValue2: nil,
	}

	for _, opt := range opts {
		opt(&options)
	}

	padPair1 := KeyValuePair[K1, V1]{}
	padPair2 := KeyValuePair[K2, V2]{}

	if options.PadKey1 != nil {
		padPair1.Key = *options.PadKey1
	}

	if options.PadValue1 != nil {
		padPair1.Value = *options.PadValue1
	}

	if options.PadKey2 != nil {
		padPair2.Key = *options.PadKey2
	}

	if options.PadValue2 != nil {
		padPair2.Value = *options.PadValue2
	}

	return Zip(
		KeyValue2(seq1),
		KeyValue2(seq2),
		WithZipPadded[KeyValuePair[K1, V1], KeyValuePair[K2, V2]](options.Padded),
		WithZipPadValue1[KeyValuePair[K1, V1], KeyValuePair[K2, V2]](padPair1),
		WithZipPadValue2[KeyValuePair[K1, V1]](padPair2),
	)
}

// ZipNOptions defines configurable options for [ZipN].
type ZipNOptions[V any] struct {
	// Padded indicates whether the shorter sequence will be padded. If
	// true, the shorter sequence is padded to match the length of the
	// longer one. If false, the number of iterations is equal to the
	// length of the shorter sequence.
	Padded bool
	// PadValue is the value the shorter sequence is padded with. This
	// is not used if Padded is false or if both sequences are equal
	// in length.
	PadValue *V
}

// ZipNOptionFunc is the function signature of configuration helpers
// for [ZipN].
type ZipNOptionFunc[V any] func(*ZipNOptions[V])

// WithZipNPadded is a helper for configuring [ZipN] to pad the shorter
// sequence.
func WithZipNPadded[V any](padded bool) ZipNOptionFunc[V] {
	return func(o *ZipNOptions[V]) {
		o.Padded = padded
	}
}

// WithZipNPadValue is a helper for configuring padded values in
// [ZipN].
func WithZipNPadValue[V any](value V) ZipNOptionFunc[V] {
	return func(o *ZipNOptions[V]) {
		o.PadValue = &value
	}
}

// ZipN allows looping over multiple [iter.Seq] sequences
// simultaneously.
func ZipN[V any](
	seqs iter.Seq[iter.Seq[V]],
	opts ...ZipNOptionFunc[V],
) iter.Seq[iter.Seq[V]] {
	options := ZipNOptions[V]{
		Padded:   false,
		PadValue: nil,
	}

	for _, opt := range opts {
		opt(&options)
	}

	return func(yield func(iter.Seq[V]) bool) {
		nexts := list.New()

		for seq := range seqs {
			next, stop := iter.Pull(seq)
			defer stop()

			nexts.PushBack(next)
		}

		for {
			i := 0
			exhaused := true
			values := make([]V, nexts.Len())

			for elem := range List(nexts) {
				next := elem.Value.(func() (V, bool))

				var (
					ok    bool
					value V
				)

				value, ok = next()
				if ok {
					exhaused = false
				}

				if !options.Padded && !ok {
					return
				}

				if options.PadValue != nil && !ok {
					value = *options.PadValue
				}

				values[i] = value
				i++
			}

			if exhaused {
				return
			}

			if !yield(Slice(values)) {
				return
			}
		}
	}
}

// ZipN2Options defines configurable options for [ZipN2].
type ZipN2Options[K, V any] struct {
	// Padded indicates whether the shorter sequence will be padded. If
	// true, the shorter sequence is padded to match the length of the
	// longer one. If false, the number of iterations is equal to the
	// length of the shorter sequence.
	Padded bool
	// PadKey is the key the shorter sequence is padded with. This is
	// not used if Padded is false or if both sequences are equal in
	// length.
	PadKey *K
	// PadValue is the value the shorter sequence is padded with. This
	// is not used if Padded is false or if both sequences are equal
	// in length.
	PadValue *V
}

// ZipN2OptionFunc is the function signature of configuration helpers
// for [ZipN2].
type ZipN2OptionFunc[K, V any] func(*ZipN2Options[K, V])

// WithZipN2Padded is a helper for configuring [ZipN2] to pad the shorter
// sequence.
func WithZipN2Padded[K, V any](padded bool) ZipN2OptionFunc[K, V] {
	return func(o *ZipN2Options[K, V]) {
		o.Padded = padded
	}
}

// WithZipN2PadKey is a helper for configuring padded keys in [ZipN2].
// first sequence in [Zip2].
func WithZipN2PadKey[K, V any](key K) ZipN2OptionFunc[K, V] {
	return func(o *ZipN2Options[K, V]) {
		o.PadKey = &key
	}
}

// WithZipN2PadValue is a helper for configuring padded values in
// [ZipN2].
func WithZipN2PadValue[K, V any](value V) ZipN2OptionFunc[K, V] {
	return func(o *ZipN2Options[K, V]) {
		o.PadValue = &value
	}
}

// ZipN2 allows looping over multiple [iter.Seq2] sequences
// simultaneously.
func ZipN2[K, V any](
	seqs iter.Seq[iter.Seq2[K, V]],
	opts ...ZipN2OptionFunc[K, V],
) iter.Seq[iter.Seq2[K, V]] {
	options := ZipN2Options[K, V]{
		Padded:   false,
		PadKey:   nil,
		PadValue: nil,
	}

	for _, opt := range opts {
		opt(&options)
	}

	padPair := KeyValuePair[K, V]{}

	if options.PadKey != nil {
		padPair.Key = *options.PadKey
	}

	if options.PadValue != nil {
		padPair.Value = *options.PadValue
	}

	return Transform(
		ZipN(
			Transform(seqs, KeyValue2),
			WithZipNPadded[KeyValuePair[K, V]](options.Padded),
			WithZipNPadValue(padPair),
		),
		KeyValue,
	)
}
