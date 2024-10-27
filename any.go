package gloop

import (
	"iter"
)

// Any computes whether or not any value in an iter.Seq sequence is
// true.
func Any(seq iter.Seq[bool]) bool {
	return Any2(Enumerate(seq))
}

// Any2 computes whether or not any value in an iter.Seq2 sequence is
// true.
func Any2[K any](seq iter.Seq2[K, bool]) bool {
	for _, value := range seq {
		if value {
			return true
		}
	}

	return false
}
