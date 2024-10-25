package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestWindowSlice(t *testing.T) {
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	wantWindows := [][]int{
		{3, 1, 4},
		{1, 4, 1},
		{4, 1, 5},
		{1, 5, 9},
		{5, 9, 2},
		{9, 2, 6},
		{2, 6, 5},
	}
	i := 0

	for seq := range gloop.Window(gloop.Slice(values), 3) {
		window := gloop.ToSlice(seq)
		require.Equal(t, wantWindows[i], window)
		i++
	}

	require.Equal(t, len(wantWindows), i)
}

func TestWindowString(t *testing.T) {
	s := "FizzBuzz"
	wantWindows := []string{
		"Fizz",
		"izzB",
		"zzBu",
		"zBuz",
		"Buzz",
	}
	i := 0

	for seq := range gloop.Window(gloop.String(s), 4) {
		window := gloop.ToString(seq)
		require.Equal(t, wantWindows[i], window)
		i++
	}

	require.Equal(t, len(wantWindows), i)
}

func TestWindowBreak(t *testing.T) {
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}

	for seq := range gloop.Window(gloop.Slice(values), 3) {
		for value := range seq {
			require.Equal(t, 3, value)
			break
		}
		break
	}
}

func TestWindowZeroSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.Window(gloop.Slice([]int{3, 1, 4}), 0) {
			t.Fatal("expected no iteration")
		}
	})
}

func TestWindowNegativeSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.Window(gloop.Slice([]int{3, 1, 4}), -1) {
			t.Fatal("expected no iteration")
		}
	})
}
