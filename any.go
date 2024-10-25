package gloop

import (
	"iter"
)

// Any computes whether or not any value in a sequence is true.
func Any(seq iter.Seq[bool]) bool {
	for value := range seq {
		if value {
			return true
		}
	}

	return false
}
