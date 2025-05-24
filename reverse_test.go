package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestReverseSlice(t *testing.T) {
	values := []string{"a", "b", "c"}
	reversedValues := []string{"c", "b", "a"}
	i := 0

	for value := range gloop.Reverse(gloop.Slice(values)) {
		require.Equal(t, reversedValues[i], value)

		i++
	}

	require.Equal(t, len(reversedValues), i)
}

func TestReverseString(t *testing.T) {
	s := "FizzBuzz"
	reversedRunes := []rune{'z', 'z', 'u', 'B', 'z', 'z', 'i', 'F'}
	i := 0

	for r := range gloop.Reverse(gloop.String(s)) {
		require.Equal(t, reversedRunes[i], r)

		i++
	}

	require.Equal(t, len(reversedRunes), i)
}

func TestReverseBreak(t *testing.T) {
	values := []string{"a", "b", "c"}
	reversedValues := []string{"c", "b"}
	i := 0

	for value := range gloop.Reverse(gloop.Slice(values)) {
		if i == 2 {
			break
		}

		require.Equal(t, reversedValues[i], value)

		i++
	}

	require.Equal(t, len(reversedValues), i)
}

func TestReverse2(t *testing.T) {
	seq := func(yield func(string, int) bool) {
		if !yield("a", 3) {
			return
		}

		if !yield("b", 4) {
			return
		}

		if !yield("c", 5) {
			return
		}
	}
	reversedKeys := []string{"c", "b", "a"}
	reversedValues := []int{5, 4, 3}
	i := 0

	for key, value := range gloop.Reverse2(seq) {
		require.Equal(t, reversedKeys[i], key)
		require.Equal(t, reversedValues[i], value)

		i++
	}

	require.Equal(t, 3, i)
}

func TestReverse2Break(t *testing.T) {
	seq := func(yield func(string, int) bool) {
		if !yield("a", 3) {
			return
		}

		if !yield("b", 4) {
			return
		}

		if !yield("c", 5) {
			return
		}
	}
	reversedKeys := []string{"c", "b"}
	reversedValues := []int{5, 4}
	i := 0

	for key, value := range gloop.Reverse2(seq) {
		if i == 2 {
			break
		}

		require.Equal(t, reversedKeys[i], key)
		require.Equal(t, reversedValues[i], value)

		i++
	}

	require.Equal(t, 2, i)
}
