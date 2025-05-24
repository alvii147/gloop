package gloop_test

import (
	"testing"
	"time"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestMaxInt(t *testing.T) {
	values := []int{3, 4, -5}
	minValue := gloop.Max(gloop.Slice(values))
	require.Equal(t, 4, minValue)
}

func TestMaxFloat(t *testing.T) {
	values := []float64{2.31, -0.03, 0.22}
	minValue := gloop.Max(gloop.Slice(values))
	require.InDelta(t, 2.31, minValue, 0.001)
}

func TestMaxString(t *testing.T) {
	values := []string{"Fizz", "Buzz", "Bazz"}
	minValue := gloop.Max(gloop.Slice(values))
	require.Equal(t, "Fizz", minValue)
}
func TestMaxDuration(t *testing.T) {
	values := []time.Duration{time.Hour, time.Minute, time.Second}
	duration := gloop.Max(gloop.Slice(values))
	require.Equal(t, time.Hour, duration)
}

func TestMaxByComparisonStringLen(t *testing.T) {
	values := []string{"Fizzz", "Buzz", "Bazzzzzz"}
	maxValue := gloop.MaxByComparison(gloop.Slice(values), func(s1 string, s2 string) bool {
		return len(s1) < len(s2)
	})
	require.Equal(t, "Bazzzzzz", maxValue)
}

func TestMaxByComparison2KeyValueProduct(t *testing.T) {
	seq := func(yield func(int, int) bool) {
		if !yield(1, 5) {
			return
		}

		if !yield(3, 1) {
			return
		}

		if !yield(4, 9) {
			return
		}
	}

	maxKey, maxValue := gloop.MaxByComparison2(seq, func(k1, v1, k2, v2 int) bool {
		return k1*v1 < k2*v2
	})
	require.Equal(t, 4, maxKey)
	require.Equal(t, 9, maxValue)
}

func TestMaxByRankStringLen(t *testing.T) {
	values := []string{"Fizzz", "Buzz", "Bazzzzzz"}
	maxValue := gloop.MaxByRank(gloop.Slice(values), func(s string) int {
		return len(s)
	})
	require.Equal(t, "Bazzzzzz", maxValue)
}

func TestMaxByRank2KeyValueProduct(t *testing.T) {
	seq := func(yield func(int, int) bool) {
		if !yield(1, 5) {
			return
		}

		if !yield(3, 1) {
			return
		}

		if !yield(4, 9) {
			return
		}
	}

	maxKey, maxValue := gloop.MaxByRank2(seq, func(k, v int) int {
		return k * v
	})
	require.Equal(t, 4, maxKey)
	require.Equal(t, 9, maxValue)
}
