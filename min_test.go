package gloop_test

import (
	"testing"
	"time"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestMinInt(t *testing.T) {
	values := []int{3, 4, -5}
	minValue := gloop.Min(gloop.Slice(values))
	require.Equal(t, -5, minValue)
}

func TestMinFloat(t *testing.T) {
	values := []float64{2.31, -0.03, 0.22}
	minValue := gloop.Min(gloop.Slice(values))
	require.InDelta(t, -0.03, minValue, 0.001)
}

func TestMinString(t *testing.T) {
	values := []string{"Fizz", "Buzz", "Bazz"}
	minValue := gloop.Min(gloop.Slice(values))
	require.Equal(t, "Bazz", minValue)
}

func TestMinDuration(t *testing.T) {
	values := []time.Duration{time.Hour, time.Minute, time.Second}
	duration := gloop.Min(gloop.Slice(values))
	require.Equal(t, time.Second, duration)
}

func TestMinByComparisonStringLen(t *testing.T) {
	values := []string{"Fizzz", "Buzz", "Bazzzzzz"}
	minValue := gloop.MinByComparison(gloop.Slice(values), func(s1 string, s2 string) bool {
		return len(s1) < len(s2)
	})
	require.Equal(t, "Buzz", minValue)
}

func TestMinByComparison2KeyValueProduct(t *testing.T) {
	seq := func(yield func(int, int) bool) {
		if !yield(1, 5) {
			return
		}

		if !yield(4, 9) {
			return
		}

		if !yield(3, 1) {
			return
		}
	}

	minKey, minValue := gloop.MinByComparison2(seq, func(k1, v1, k2, v2 int) bool {
		return k1*v1 < k2*v2
	})
	require.Equal(t, 3, minKey)
	require.Equal(t, 1, minValue)
}

func TestMinByRankStringLen(t *testing.T) {
	values := []string{"Fizzz", "Buzz", "Bazzzzzz"}
	minValue := gloop.MinByRank(gloop.Slice(values), func(s string) int {
		return len(s)
	})
	require.Equal(t, "Buzz", minValue)
}

func TestMinByRank2KeyValueProduct(t *testing.T) {
	seq := func(yield func(int, int) bool) {
		if !yield(1, 5) {
			return
		}

		if !yield(4, 9) {
			return
		}

		if !yield(3, 1) {
			return
		}
	}

	minKey, minValue := gloop.MinByRank2(seq, func(k, v int) int {
		return k * v
	})
	require.Equal(t, 3, minKey)
	require.Equal(t, 1, minValue)
}
