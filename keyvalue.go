package gloop

import "iter"

// KeyValuePair represents a generic key value pair.
type KeyValuePair[K, V any] struct {
	Key   K
	Value V
}

// KeyValue converts an iter.Seq [KeyValuePair] sequence to an
// iter.Seq2 sequence.
func KeyValue[K, V any](seq iter.Seq[KeyValuePair[K, V]]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for pair := range seq {
			if !yield(pair.Key, pair.Value) {
				return
			}
		}
	}
}

// KeyValue2 converts an iter.Seq2 sequence to an iter.Seq
// [KeyValuePair] sequence.
func KeyValue2[K, V any](seq iter.Seq2[K, V]) iter.Seq[KeyValuePair[K, V]] {
	return Transform2(seq, func(k K, v V) KeyValuePair[K, V] {
		return KeyValuePair[K, V]{
			Key:   k,
			Value: v,
		}
	})
}
