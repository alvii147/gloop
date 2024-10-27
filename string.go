package gloop

import (
	"iter"
	"strings"
)

// String allows looping over the runes in a given string.
func String(s string) iter.Seq[rune] {
	return Values(String2(s))
}

// String allows looping over the runes in a given string with an
// index.
func String2(s string) iter.Seq2[int, rune] {
	return func(yield func(int, rune) bool) {
		for i, r := range s {
			if !yield(i, r) {
				return
			}
		}
	}
}

// ToString converts an iter.Seq sequence of runes to a string.
func ToString(seq iter.Seq[rune]) string {
	var sb strings.Builder
	for value := range seq {
		sb.WriteRune(value)
	}

	s := sb.String()

	return s
}
