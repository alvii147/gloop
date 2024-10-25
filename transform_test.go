package gloop_test

import (
	"strconv"
	"strings"
	"testing"
	"unicode"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestTransformSliceDoubleValue(t *testing.T) {
	values := []int{3, 4, 5}
	i := 0

	for value := range gloop.Transform(gloop.Slice(values), func(value int) int {
		return value * 2
	}) {
		require.Equal(t, values[i]*2, value)
		i++
	}

	require.Equal(t, len(values), i)
}

func TestTransformSliceItoa(t *testing.T) {
	values := []int{3, 4, 5}
	wantValues := []string{"3", "4", "5"}
	i := 0

	for value := range gloop.Transform(gloop.Slice(values), strconv.Itoa) {
		require.Equal(t, wantValues[i], value)
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestTransformSliceTrimSpace(t *testing.T) {
	values := []string{"   Fizz", " Buzz  "}
	wantValues := []string{"Fizz", "Buzz"}
	i := 0

	for value := range gloop.Transform(gloop.Slice(values), strings.TrimSpace) {
		require.Equal(t, wantValues[i], value)
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestTransformStringToLower(t *testing.T) {
	s := "FiZz"
	wantRunes := []rune{'f', 'i', 'z', 'z'}
	i := 0

	for r := range gloop.Transform(gloop.String(s), unicode.ToLower) {
		require.Equal(t, wantRunes[i], r)
		i++
	}

	require.Equal(t, len(wantRunes), i)
}

func TestTransformBreak(t *testing.T) {
	values := []int{3, 4, 5}
	wantValues := []int{3, 4}
	i := 0

	for value := range gloop.Transform(gloop.Slice(values), func(value int) int {
		return value
	}) {
		if i == 2 {
			break
		}

		require.Equal(t, wantValues[i], value)
		i++
	}

	require.Equal(t, len(wantValues), i)
}
