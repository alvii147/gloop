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

func TestFilter2MapPositiveProduct(t *testing.T) {
	m := map[int]int{
		-1: -2,
		8:  -3,
		-4: 9,
	}
	i := 0

	for key, value := range gloop.Filter2(gloop.Map(m), func(key int, value int) bool {
		return (key * value) >= 0
	}) {
		require.Equal(t, -1, key)
		require.Equal(t, -2, value)
		i++
	}

	require.Equal(t, 1, i)
}

func TestFilter2MapCorrectLen(t *testing.T) {
	m := map[string]int{
		"Fizz":     8,
		"Buzz":     4,
		"FizzBuzz": 4,
	}
	i := 0

	for key, value := range gloop.Filter2(gloop.Map(m), func(key string, value int) bool {
		return len(key) == value
	}) {
		require.Equal(t, "Buzz", key)
		require.Equal(t, 4, value)
		i++
	}

	require.Equal(t, 1, i)
}

func TestFilter2Break(t *testing.T) {
	values := []string{"Fizz", "Buzz", "Bazz"}
	wantValues := []string{"Fizz", "Buzz"}
	i := 0

	for idx, value := range gloop.Filter2(gloop.Slice2(values), func(_ int, value string) bool {
		return true
	}) {
		if i == 2 {
			break
		}

		require.Equal(t, i, idx)
		require.Equal(t, wantValues[i], value)
		i++
	}

	require.Equal(t, 2, i)
}
