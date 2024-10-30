package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestReduceSliceMin(t *testing.T) {
	values := []int{3, -1, 4}
	minValue := gloop.Reduce(gloop.Slice(values), func(value1 int, value2 int) int {
		return min(value1, value2)
	})
	require.Equal(t, -1, minValue)
}

func TestReduceSliceMax(t *testing.T) {
	values := []int{3, -1, 4}
	maxValue := gloop.Reduce(gloop.Slice(values), func(value1 int, value2 int) int {
		return max(value1, value2)
	})
	require.Equal(t, 4, maxValue)
}

func TestReduceSliceProduct(t *testing.T) {
	values := []int{3, -1, 4}
	product := gloop.Reduce(gloop.Slice(values), func(value1 int, value2 int) int {
		return value1 * value2
	})
	require.Equal(t, -12, product)
}

func TestReduceStringMinRune(t *testing.T) {
	s := "FizzBuzz"

	minRune := gloop.Reduce(gloop.String(s), func(value1 rune, value2 rune) rune {
		return min(value1, value2)
	})
	require.Equal(t, 'B', minRune)
}

func TestReduceStringMaxRune(t *testing.T) {
	s := "FizzBuzz"

	maxRune := gloop.Reduce(gloop.String(s), func(value1 rune, value2 rune) rune {
		return max(value1, value2)
	})
	require.Equal(t, 'z', maxRune)
}

func TestReduce2MapMinValue(t *testing.T) {
	m := map[int]int{
		0: 3,
		1: -1,
		2: 4,
	}

	minValueKey, minValue := gloop.Reduce2(gloop.Map(m), func(key1 int, value1 int, key2 int, value2 int) (int, int) {
		if value1 < value2 {
			return key1, value1
		}

		return key2, value2
	})

	require.Equal(t, 1, minValueKey)
	require.Equal(t, -1, minValue)
}

func TestReduce2MapMaxValue(t *testing.T) {
	m := map[int]int{
		0: 3,
		1: -1,
		2: 4,
	}

	maxValueKey, maxValue := gloop.Reduce2(gloop.Map(m), func(key1 int, value1 int, key2 int, value2 int) (int, int) {
		if value1 > value2 {
			return key1, value1
		}

		return key2, value2
	})

	require.Equal(t, 2, maxValueKey)
	require.Equal(t, 4, maxValue)
}
