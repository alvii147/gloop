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

func TestWindow2Slice(t *testing.T) {
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	wantKeys := [][]int{
		{0, 1, 2},
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
		{4, 5, 6},
		{5, 6, 7},
		{6, 7, 8},
	}
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

	for seq := range gloop.Window2(gloop.Enumerate(gloop.Slice(values)), 3) {
		keys, window := gloop.ToSlice2(seq)
		require.Equal(t, wantKeys[i], keys)
		require.Equal(t, wantWindows[i], window)
		i++
	}

	require.Equal(t, len(wantWindows), i)
}

func TestWindow2String(t *testing.T) {
	s := "FizzBuzz"
	wantKeys := [][]int{
		{0, 1, 2, 3},
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{3, 4, 5, 6},
		{4, 5, 6, 7},
	}
	wantWindows := []string{
		"Fizz",
		"izzB",
		"zzBu",
		"zBuz",
		"Buzz",
	}
	i := 0

	for seq := range gloop.Window2(gloop.Enumerate(gloop.String(s)), 4) {
		keys, windowRunes := gloop.ToSlice2(seq)
		window := string(windowRunes)
		require.Equal(t, wantKeys[i], keys)
		require.Equal(t, wantWindows[i], window)
		i++
	}

	require.Equal(t, len(wantWindows), i)
}

func TestWindow2Break(t *testing.T) {
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	i := 0

	for seq := range gloop.Window2(gloop.Enumerate(gloop.Slice(values)), 3) {
		if i == 1 {
			break
		}

		for key, value := range seq {
			require.Equal(t, 0, key)
			require.Equal(t, 3, value)
			break
		}
		i++
	}
}

func TestWindow2ZeroSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.Window2(gloop.Enumerate(gloop.Slice([]int{3, 1, 4})), 0) {
			t.Fatal("expected no iteration")
		}
	})
}

func TestWindow2NegativeSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.Window2(gloop.Enumerate(gloop.Slice([]int{3, 1, 4})), -1) {
			t.Fatal("expected no iteration")
		}
	})
}
