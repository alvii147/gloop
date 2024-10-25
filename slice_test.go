package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestSlice(t *testing.T) {
	values := []int{3, 4, 5}
	i := 0
	for value := range gloop.Slice(values) {
		require.Equal(t, values[i], value)
		i++
	}

	require.Equal(t, len(values), i)
}

func TestSliceBreak(t *testing.T) {
	values := []int{3, 4, 5}
	wantValues := []int{3, 4}
	i := 0

	for value := range gloop.Slice(values) {
		if i == 2 {
			break
		}

		require.Equal(t, wantValues[i], value)
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestToSlice(t *testing.T) {
	seq := func(yield func(int) bool) {
		if !yield(3) {
			return
		}

		if !yield(4) {
			return
		}

		if !yield(5) {
			return
		}
	}

	values := gloop.ToSlice(seq)
	require.Equal(t, []int{3, 4, 5}, values)
}
