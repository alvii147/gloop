package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestUniqueDuplicates(t *testing.T) {
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	wantValues := []int{3, 1, 4, 5, 9, 2, 6}
	i := 0

	for value := range gloop.Unique(gloop.Slice(values)) {
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestUniqueNoDuplicates(t *testing.T) {
	values := []int{3, 1, 4}
	wantValues := []int{3, 1, 4}
	i := 0

	for value := range gloop.Unique(gloop.Slice(values)) {
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestUniqueBreak(t *testing.T) {
	values := []int{3, 1, 4}
	wantValues := []int{3, 1}
	i := 0

	for value := range gloop.Unique(gloop.Slice(values)) {
		if i == 2 {
			break
		}

		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestUnique2Duplicates(t *testing.T) {
	seq := func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Buzz", 1) {
			return
		}

		if !yield("Fizz", 3) {
			return
		}
	}

	wantKeys := []string{"Fizz", "Buzz"}
	wantValues := []int{3, 1}
	i := 0

	for key, value := range gloop.Unique2(seq) {
		require.Equal(t, wantKeys[i], key)
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantKeys), i)
	require.Equal(t, len(wantValues), i)
}

func TestUnique2NoDuplicates(t *testing.T) {
	seq := func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Buzz", 1) {
			return
		}

		if !yield("Fizz", 4) {
			return
		}
	}

	wantKeys := []string{"Fizz", "Buzz", "Fizz"}
	wantValues := []int{3, 1, 4}
	i := 0

	for key, value := range gloop.Unique2(seq) {
		require.Equal(t, wantKeys[i], key)
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantKeys), i)
	require.Equal(t, len(wantValues), i)
}

func TestUnique2Break(t *testing.T) {
	seq := func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Buzz", 1) {
			return
		}

		if !yield("Fizz", 4) {
			return
		}
	}

	wantKeys := []string{"Fizz", "Buzz"}
	wantValues := []int{3, 1}
	i := 0

	for key, value := range gloop.Unique2(seq) {
		if i == 2 {
			break
		}

		require.Equal(t, wantKeys[i], key)
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantKeys), i)
	require.Equal(t, len(wantValues), i)
}
