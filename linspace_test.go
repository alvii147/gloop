package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestLinspace(t *testing.T) {
	wantValues := []float64{2, 2.25, 2.5, 2.75}
	i := 0

	for value := range gloop.Linspace(2, 3, 5) {
		require.InDelta(t, wantValues[i], value, 0.01)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestLinspaceClosed(t *testing.T) {
	wantValues := []float64{2, 2.25, 2.5, 2.75, 3}
	i := 0

	for value := range gloop.Linspace(2, 3, 5, gloop.WithLinspaceClosed(true)) {
		require.InDelta(t, wantValues[i], value, 0.01)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestLinspaceBackwards(t *testing.T) {
	wantValues := []float64{10, 8.5, 7, 5.5}
	i := 0

	for value := range gloop.Linspace(10, 4, 5) {
		require.InDelta(t, wantValues[i], value, 0.01)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestLinspaceBackwardsClosed(t *testing.T) {
	wantValues := []float64{10, 8.5, 7, 5.5, 4}
	i := 0

	for value := range gloop.Linspace(10, 4, 5, gloop.WithLinspaceClosed(true)) {
		require.InDelta(t, wantValues[i], value, 0.01)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestLinspaceBreak(t *testing.T) {
	wantValues := []float64{2, 2.25}
	i := 0

	for value := range gloop.Linspace(2, 3, 5) {
		if i == 2 {
			break
		}

		require.InDelta(t, wantValues[i], value, 0.001)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestLinspaceZeroDivisionsPanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.Linspace(2, 3, 0) {
			t.Fatal("expected no iteration")
		}
	})
}

func TestLinspaceOneDivisionPanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.Linspace(2, 3, 1) {
			t.Fatal("expected no iteration")
		}
	})
}

func TestLinspaceNegativeDivisionsPanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.Linspace(2, 3, -5) {
			t.Fatal("expected no iteration")
		}
	})
}
