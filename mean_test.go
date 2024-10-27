package gloop_test

import (
	"testing"
	"time"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestMeanInt(t *testing.T) {
	values := []int{7, -2, 4}
	mean := gloop.Mean(gloop.Slice(values))
	require.Equal(t, 3.0, mean)
}

func TestMeanFloat(t *testing.T) {
	values := []float64{0.4, 8.9, -4.8}
	mean := gloop.Mean(gloop.Slice(values))
	require.InDelta(t, 1.5, mean, 0.01)
}

func TestMeanDuration(t *testing.T) {
	values := []time.Duration{time.Hour, 2 * time.Minute}
	mean := gloop.Mean(gloop.Slice(values))
	require.Equal(t, float64(31*time.Minute), mean)
}
