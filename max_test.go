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
	require.Equal(t, 2.31, minValue)
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

func TestMax2Int(t *testing.T) {
	values := []int{3, 4, -5}
	minValue := gloop.Max2(gloop.Slice2(values))
	require.Equal(t, 4, minValue)
}

func TestMax2Float(t *testing.T) {
	values := []float64{2.31, -0.03, 0.22}
	minValue := gloop.Max2(gloop.Slice2(values))
	require.Equal(t, 2.31, minValue)
}

func TestMax2String(t *testing.T) {
	values := []string{"Fizz", "Buzz", "Bazz"}
	minValue := gloop.Max2(gloop.Slice2(values))
	require.Equal(t, "Fizz", minValue)
}
func TestMax2Duration(t *testing.T) {
	values := []time.Duration{time.Hour, time.Minute, time.Second}
	duration := gloop.Max2(gloop.Slice2(values))
	require.Equal(t, time.Hour, duration)
}
