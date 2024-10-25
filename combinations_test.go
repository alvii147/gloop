package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestCombinationsSlice(t *testing.T) {
	values := []int{3, 1, 4, 2}
	wantCombs := [][]int{
		{3, 1, 4},
		{3, 1, 2},
		{3, 4, 2},
		{1, 4, 2},
	}
	i := 0

	for seq := range gloop.Combinations(gloop.Slice(values), 3) {
		comb := gloop.ToSlice(seq)
		require.Equal(t, wantCombs[i], comb)
		i++
	}

	require.Equal(t, len(wantCombs), i)
}

func TestCombinationsString(t *testing.T) {
	s := "ABCD"
	wantCombs := []string{"AB", "AC", "AD", "BC", "BD", "CD"}
	i := 0

	for seq := range gloop.Combinations(gloop.String(s), 2) {
		comb := gloop.ToString(seq)
		require.Equal(t, wantCombs[i], comb)
		i++
	}

	require.Equal(t, len(wantCombs), i)
}

func TestCombinationsBreak(t *testing.T) {
	values := []int{3, 1, 4, 2}
	i := 0

	for seq := range gloop.Combinations(gloop.Slice(values), 2) {
		if i == 1 {
			break
		}

		for value := range seq {
			require.Equal(t, 3, value)
			break
		}
		i++
	}
}

func TestCombinationsZeroSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.Combinations(gloop.Slice([]int{3, 1, 4}), 0) {
			t.Fatal("expected no iteration")
		}
	})
}

func TestCombinationsNegativeSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.Combinations(gloop.Slice([]int{3, 1, 4}), -3) {
			t.Fatal("expected no iteration")
		}
	})
}
