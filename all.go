package gloop

import (
	"iter"
)

// All computes whether or not all values in an iter.Seq sequence are
// true.
func All(seq iter.Seq[bool]) bool {
	return All2(Enumerate(seq))
}

// All2 computes whether or not all values in an iter.Seq2 sequence are
// true.
func All2[K any](seq iter.Seq2[K, bool]) bool {
	for _, value := range seq {
		if !value {
			return false
		}
	}

	return true
}
