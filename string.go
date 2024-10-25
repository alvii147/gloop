package gloop

import (
	"iter"
	"strings"
)

// String allows a for loop to range over the runes in a given string.
func String(s string) iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for _, r := range s {
			if !yield(r) {
				return
			}
		}
	}
}

// ToString converts a sequence of runes to a string.
func ToString(seq iter.Seq[rune]) string {
	var sb strings.Builder
	for value := range seq {
		sb.WriteRune(value)
	}

	s := sb.String()

	return s
}
