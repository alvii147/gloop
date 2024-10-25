package gloop_test

import (
	"testing"
	"unicode"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestWithReduceInitialValue(t *testing.T) {
	initialValue := 42
	options := gloop.ReduceOptions[int]{}
	gloop.WithReduceInitialValue(initialValue)(&options)

	require.NotNil(t, options.InitialValue)
	require.Equal(t, initialValue, *options.InitialValue)
}

func TestReduceSliceSum(t *testing.T) {
	values := []int{1, 2, 3}
	sum := gloop.Reduce(gloop.Slice(values), func(acc int, value int) int {
		return acc + value
	})
	require.Equal(t, 6, sum)
}

func TestReduceSliceMultiplicationWithInitialValue(t *testing.T) {
	values := []int{3, 4, 5}

	multiplication := gloop.Reduce(gloop.Slice(values), func(acc int, value int) int {
		return acc * value
	}, gloop.WithReduceInitialValue(1))
	require.Equal(t, 60, multiplication)
}

func TestReduceSliceConcatenate(t *testing.T) {
	values := []string{"Fizz", "Buzz"}

	concatString := gloop.Reduce(gloop.Slice(values), func(acc string, value string) string {
		return acc + value
	})
	require.Equal(t, "FizzBuzz", concatString)
}

func TestReduceSliceLen(t *testing.T) {
	values := []string{"Fizz", "Buzz"}

	lenStrings := gloop.Reduce(gloop.Slice(values), func(acc int, value string) int {
		return acc + len(value)
	})
	require.Equal(t, 8, lenStrings)
}

func TestReduceStringNumericCount(t *testing.T) {
	s := "F1Z2bU2z"

	sum := gloop.Reduce(gloop.String(s), func(acc int, r rune) int {
		if unicode.IsNumber(r) {
			return acc + 1
		}

		return acc
	})
	require.Equal(t, 3, sum)
}
