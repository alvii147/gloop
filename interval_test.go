package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestIntervalInt(t *testing.T) {
	wantValues := []int{3, 5}
	i := 0

	for value := range gloop.Interval(3, 7, 2) {
		require.Equal(t, wantValues[i], value)
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestIntervalFloat(t *testing.T) {
	wantValues := []float64{2, 2.5, 3, 3.5, 4, 4.5}
	i := 0

	for value := range gloop.Interval(2, 5, 0.5) {
		require.Equal(t, wantValues[i], value)
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestIntervalNegativeStep(t *testing.T) {
	wantValues := []int{10, 7, 4}
	i := 0

	for value := range gloop.Interval(10, 3, -3) {
		require.Equal(t, wantValues[i], value)
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestIntervalZeroStep(t *testing.T) {
	for range gloop.Interval(3, 7, 0) {
		t.Fatal("expected no iteration")
	}
}

func TestIntervalNoIteration(t *testing.T) {
	for range gloop.Interval(10, 3, 2) {
		t.Fatal("expected no iteration")
	}
}

func TestIntervalNoIterationNegativeStep(t *testing.T) {
	for range gloop.Interval(3, 7, -3) {
		t.Fatal("expected no iteration")
	}
}

func TestIntervalClosed(t *testing.T) {
	wantValues := []int{3, 5, 7}
	i := 0

	for value := range gloop.Interval(3, 7, 2, gloop.WithIntervalClosed(true)) {
		require.Equal(t, wantValues[i], value)
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestIntervalClosedNegativeStep(t *testing.T) {
	wantValues := []int{10, 7, 4, 1}
	i := 0

	for value := range gloop.Interval(10, 1, -3, gloop.WithIntervalClosed(true)) {
		require.Equal(t, wantValues[i], value)
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestIntervalBreak(t *testing.T) {
	wantValues := []int{3, 6}
	i := 0

	for value := range gloop.Interval(3, 20, 3) {
		if i == 2 {
			break
		}

		require.Equal(t, wantValues[i], value)
		i++
	}

	require.Equal(t, len(wantValues), i)
}
