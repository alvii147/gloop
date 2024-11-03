package gloop

import (
	"iter"
)

// All computes whether or not all values in an [iter.Seq] sequence are
// true.
func All(seq iter.Seq[bool]) bool {
	for value := range seq {
		if !value {
			return false
		}
	}

	return true
}
