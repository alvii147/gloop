package gloop_test

import (
	"testing"
	"unicode"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestWithFoldInitialValue(t *testing.T) {
	initialValue := 42
	options := gloop.FoldOptions[int]{}
	gloop.WithFoldInitialValue(initialValue)(&options)

	require.NotNil(t, options.InitialValue)
	require.Equal(t, initialValue, *options.InitialValue)
}

func TestFoldSliceSum(t *testing.T) {
	values := []int{1, 2, 3}
	sum := gloop.Fold(gloop.Slice(values), func(acc int, value int) int {
		return acc + value
	})
	require.Equal(t, 6, sum)
}

func TestFoldSliceProductWithInitialValue(t *testing.T) {
	values := []int{3, 4, 5}

	product := gloop.Fold(gloop.Slice(values), func(acc int, value int) int {
		return acc * value
	}, gloop.WithFoldInitialValue(1))
	require.Equal(t, 60, product)
}

func TestFoldSliceConcatenate(t *testing.T) {
	values := []string{"Fizz", "Buzz"}

	concatString := gloop.Fold(gloop.Slice(values), func(acc string, value string) string {
		return acc + value
	})
	require.Equal(t, "FizzBuzz", concatString)
}

func TestFoldSliceLen(t *testing.T) {
	values := []string{"Fizz", "Buzz"}

	lenStrings := gloop.Fold(gloop.Slice(values), func(acc int, value string) int {
		return acc + len(value)
	})
	require.Equal(t, 8, lenStrings)
}

func TestFoldStringNumericCount(t *testing.T) {
	s := "F1Z2bU2z"

	sum := gloop.Fold(gloop.String(s), func(acc int, r rune) int {
		if unicode.IsNumber(r) {
			return acc + 1
		}

		return acc
	})
	require.Equal(t, 3, sum)
}

func TestFold2MapSumOfProducts(t *testing.T) {
	m := map[int]int{
		3:  4,
		8:  -1,
		-2: -5,
	}
	sum := gloop.Fold2(gloop.Map(m), func(acc int, key int, value int) int {
		return acc + (key * value)
	})
	require.Equal(t, 14, sum)
}

func TestFold2MapProductOfValues(t *testing.T) {
	m := map[string]int{
		"Fizz": 4,
		"Buzz": -1,
		"Bazz": -5,
	}
	product := gloop.Fold2(gloop.Map(m), func(acc int, key string, value int) int {
		return acc * value
	}, gloop.WithFoldInitialValue(1))
	require.Equal(t, 20, product)
}
