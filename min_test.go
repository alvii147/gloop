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
	require.Equal(t, -0.03, minValue)
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
