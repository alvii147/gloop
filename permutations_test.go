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

func TestPermutations2Slice(t *testing.T) {
	values := []int{3, 1, 4, 2}
	wantKeys := [][]int{
		{0, 1, 2},
		{0, 1, 3},
		{0, 2, 1},
		{0, 2, 3},
		{0, 3, 1},
		{0, 3, 2},
		{1, 0, 2},
		{1, 0, 3},
		{1, 2, 0},
		{1, 2, 3},
		{1, 3, 0},
		{1, 3, 2},
		{2, 0, 1},
		{2, 0, 3},
		{2, 1, 0},
		{2, 1, 3},
		{2, 3, 0},
		{2, 3, 1},
		{3, 0, 1},
		{3, 0, 2},
		{3, 1, 0},
		{3, 1, 2},
		{3, 2, 0},
		{3, 2, 1},
	}
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

	for seq := range gloop.Permutations2(gloop.Enumerate(gloop.Slice(values)), 3) {
		keys, perm := gloop.ToSlice2(seq)
		require.Equal(t, wantKeys[i], keys)
		require.Equal(t, wantPerms[i], perm)

		i++
	}

	require.Equal(t, len(wantPerms), i)
}

func TestPermutations2SliceFullLength(t *testing.T) {
	values := []int{3, 1, 4}
	wantKeys := [][]int{
		{0, 1, 2},
		{0, 2, 1},
		{1, 0, 2},
		{1, 2, 0},
		{2, 0, 1},
		{2, 1, 0},
	}
	wantPerms := [][]int{
		{3, 1, 4},
		{3, 4, 1},
		{1, 3, 4},
		{1, 4, 3},
		{4, 3, 1},
		{4, 1, 3},
	}
	i := 0

	for seq := range gloop.Permutations2(gloop.Enumerate(gloop.Slice(values)), 3) {
		keys, perm := gloop.ToSlice2(seq)
		require.Equal(t, wantKeys[i], keys)
		require.Equal(t, wantPerms[i], perm)

		i++
	}

	require.Equal(t, len(wantPerms), i)
}

func TestPermutations2String(t *testing.T) {
	s := "ABCD"
	wantKeys := [][]int{
		{0, 1}, {0, 2}, {0, 3}, {1, 0}, {1, 2}, {1, 3}, {2, 0}, {2, 1}, {2, 3}, {3, 0}, {3, 1}, {3, 2},
	}
	wantPerms := []string{"AB", "AC", "AD", "BA", "BC", "BD", "CA", "CB", "CD", "DA", "DB", "DC"}
	i := 0

	for seq := range gloop.Permutations2(gloop.Enumerate(gloop.String(s)), 2) {
		keys, permRunes := gloop.ToSlice2(seq)
		perm := string(permRunes)

		require.Equal(t, wantKeys[i], keys)
		require.Equal(t, wantPerms[i], perm)

		i++
	}

	require.Equal(t, len(wantPerms), i)
}

func TestPermutationsString2FullLength(t *testing.T) {
	s := "ABC"
	wantKeys := [][]int{
		{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0},
	}
	wantPerms := []string{"ABC", "ACB", "BAC", "BCA", "CAB", "CBA"}
	i := 0

	for seq := range gloop.Permutations2(gloop.Enumerate(gloop.String(s)), 3) {
		keys, permRunes := gloop.ToSlice2(seq)
		perm := string(permRunes)

		require.Equal(t, wantKeys[i], keys)
		require.Equal(t, wantPerms[i], perm)

		i++
	}

	require.Equal(t, len(wantPerms), i)
}

func TestPermutations2Break(t *testing.T) {
	values := []int{3, 1, 4, 2}
	i := 0

	for seq := range gloop.Permutations2(gloop.Enumerate(gloop.Slice(values)), 2) {
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

func TestPermutations2ZeroSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.Permutations2(gloop.Enumerate(gloop.Slice([]int{3, 1, 4})), 0) {
			t.Fatal("expected no iteration")
		}
	})
}

func TestPermutations2NegativeSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.Permutations2(gloop.Enumerate(gloop.Slice([]int{3, 1, 4})), -3) {
			t.Fatal("expected no iteration")
		}
	})
}
