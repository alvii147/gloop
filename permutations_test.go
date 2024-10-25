package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestPermutationsSlice(t *testing.T) {
	values := []int{3, 1, 4, 2}
	wantPerms := [][]int{
		{3, 1, 4},
		{3, 1, 2},
		{3, 4, 1},
		{3, 4, 2},
		{3, 2, 1},
		{3, 2, 4},
		{1, 3, 4},
		{1, 3, 2},
		{1, 4, 3},
		{1, 4, 2},
		{1, 2, 3},
		{1, 2, 4},
		{4, 3, 1},
		{4, 3, 2},
		{4, 1, 3},
		{4, 1, 2},
		{4, 2, 3},
		{4, 2, 1},
		{2, 3, 1},
		{2, 3, 4},
		{2, 1, 3},
		{2, 1, 4},
		{2, 4, 3},
		{2, 4, 1},
	}
	i := 0

	for seq := range gloop.Permutations(gloop.Slice(values), 3) {
		perm := gloop.ToSlice(seq)
		require.Equal(t, wantPerms[i], perm)
		i++
	}

	require.Equal(t, len(wantPerms), i)
}

func TestPermutationsSliceFullLength(t *testing.T) {
	values := []int{3, 1, 4}
	wantPerms := [][]int{
		{3, 1, 4},
		{3, 4, 1},
		{1, 3, 4},
		{1, 4, 3},
		{4, 3, 1},
		{4, 1, 3},
	}
	i := 0

	for seq := range gloop.Permutations(gloop.Slice(values), 3) {
		perm := gloop.ToSlice(seq)
		require.Equal(t, wantPerms[i], perm)
		i++
	}

	require.Equal(t, len(wantPerms), i)
}

func TestPermutationsString(t *testing.T) {
	s := "ABCD"
	wantPerms := []string{"AB", "AC", "AD", "BA", "BC", "BD", "CA", "CB", "CD", "DA", "DB", "DC"}
	i := 0

	for seq := range gloop.Permutations(gloop.String(s), 2) {
		perm := gloop.ToString(seq)
		require.Equal(t, wantPerms[i], perm)
		i++
	}

	require.Equal(t, len(wantPerms), i)
}

func TestPermutationsStringFullLength(t *testing.T) {
	s := "ABC"
	wantPerms := []string{"ABC", "ACB", "BAC", "BCA", "CAB", "CBA"}
	i := 0

	for seq := range gloop.Permutations(gloop.String(s), 3) {
		perm := gloop.ToString(seq)
		require.Equal(t, wantPerms[i], perm)
		i++
	}

	require.Equal(t, len(wantPerms), i)
}

func TestPermutationsBreak(t *testing.T) {
	values := []int{3, 1, 4, 2}
	i := 0

	for seq := range gloop.Permutations(gloop.Slice(values), 2) {
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

func TestPermutationsZeroSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.Permutations(gloop.Slice([]int{3, 1, 4}), 0) {
			t.Fatal("expected no iteration")
		}
	})
}

func TestPermutationsNegativeSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.Permutations(gloop.Slice([]int{3, 1, 4}), -3) {
			t.Fatal("expected no iteration")
		}
	})
}
