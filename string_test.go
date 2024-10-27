package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	wantRunes := []rune{'F', 'i', 'z', 'z'}
	i := 0

	for r := range gloop.String("Fizz") {
		require.Equal(t, wantRunes[i], r)
		i++
	}

	require.Equal(t, len(wantRunes), i)
}

func TestStringBreak(t *testing.T) {
	wantRunes := []rune{'F', 'i'}
	i := 0

	for r := range gloop.String("Fizz") {
		if i == 2 {
			break
		}

		require.Equal(t, wantRunes[i], r)
		i++
	}

	require.Equal(t, len(wantRunes), i)
}

func TestString2(t *testing.T) {
	wantRunes := []rune{'F', 'i', 'z', 'z'}
	i := 0

	for idx, r := range gloop.String2("Fizz") {
		require.Equal(t, i, idx)
		require.Equal(t, wantRunes[i], r)
		i++
	}

	require.Equal(t, len(wantRunes), i)
}

func TestString2Break(t *testing.T) {
	wantRunes := []rune{'F', 'i'}
	i := 0

	for idx, r := range gloop.String2("Fizz") {
		if i == 2 {
			break
		}

		require.Equal(t, i, idx)
		require.Equal(t, wantRunes[i], r)
		i++
	}

	require.Equal(t, len(wantRunes), i)
}

func TestToString(t *testing.T) {
	seq := func(yield func(rune) bool) {
		if !yield('F') {
			return
		}

		if !yield('i') {
			return
		}

		if !yield('z') {
			return
		}

		if !yield('z') {
			return
		}
	}

	s := gloop.ToString(seq)
	require.Equal(t, "Fizz", s)
}
