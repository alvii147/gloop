package gloop_test

import (
	"testing"
	"time"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestSumInt(t *testing.T) {
	values := []int{3, 4, -5}
	sum := gloop.Sum(gloop.Slice(values))
	require.Equal(t, 2, sum)
}

func TestSumFloat(t *testing.T) {
	values := []float64{2.31, -0.03, 0.22}
	sum := gloop.Sum(gloop.Slice(values))
	require.InDelta(t, 2.5, sum, 0.01)
}

func TestSumComplex(t *testing.T) {
	values := []complex128{complex(1, 2), complex(3, -4)}
	sum := gloop.Sum(gloop.Slice(values))
	require.Equal(t, complex(4, -2), sum)
}

func TestSumString(t *testing.T) {
	values := []string{"Fizz", "Buzz"}
	concat := gloop.Sum(gloop.Slice(values))
	require.Equal(t, "FizzBuzz", concat)
}

func TestSumDuration(t *testing.T) {
	values := []time.Duration{time.Hour, time.Minute, time.Second}
	duration := gloop.Sum(gloop.Slice(values))
	require.Equal(t, 3661*time.Second, duration)
}
