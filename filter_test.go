package gloop_test

import (
	"testing"
	"unicode"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestFilterSliceOdd(t *testing.T) {
	values := []int{3, 4, 5}
	wantValues := []int{3, 5}
	i := 0

	for value := range gloop.Filter(gloop.Slice(values), func(value int) bool {
		return value%2 == 1
	}) {
		require.Equal(t, wantValues[i], value)
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestFilterSliceLen(t *testing.T) {
	values := []string{"Fizz", "Buzz", "FizzBuzz"}
	wantValues := []string{"Fizz", "Buzz"}
	i := 0

	for value := range gloop.Filter(gloop.Slice(values), func(value string) bool {
		return len(value) == 4
	}) {
		require.Equal(t, wantValues[i], value)
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestFilterStringLower(t *testing.T) {
	s := "FiZz"
	wantRunes := []rune{'i', 'z'}
	i := 0

	for r := range gloop.Filter(gloop.String(s), unicode.IsLower) {
		require.Equal(t, wantRunes[i], r)
		i++
	}

	require.Equal(t, len(wantRunes), i)
}

func TestFilterBreak(t *testing.T) {
	values := []int{3, 4, 5}
	wantValues := []int{3, 4}
	i := 0

	for value := range gloop.Filter(gloop.Slice(values), func(value int) bool {
		return true
	}) {
		if i == 2 {
			break
		}

		require.Equal(t, wantValues[i], value)
		i++
	}

	require.Equal(t, len(wantValues), i)
}
