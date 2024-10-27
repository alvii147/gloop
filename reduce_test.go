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

func TestReduceSliceProductWithInitialValue(t *testing.T) {
	values := []int{3, 4, 5}

	product := gloop.Reduce(gloop.Slice(values), func(acc int, value int) int {
		return acc * value
	}, gloop.WithReduceInitialValue(1))
	require.Equal(t, 60, product)
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

func TestReduce2MapSumOfProducts(t *testing.T) {
	m := map[int]int{
		3:  4,
		8:  -1,
		-2: -5,
	}
	sum := gloop.Reduce2(gloop.Map2(m), func(acc int, key int, value int) int {
		return acc + (key * value)
	})
	require.Equal(t, 14, sum)
}

func TestReduce2MapProductOfValues(t *testing.T) {
	m := map[string]int{
		"Fizz": 4,
		"Buzz": -1,
		"Bazz": -5,
	}
	product := gloop.Reduce2(gloop.Map2(m), func(acc int, key string, value int) int {
		return acc * value
	}, gloop.WithReduceInitialValue(1))
	require.Equal(t, 20, product)
}
