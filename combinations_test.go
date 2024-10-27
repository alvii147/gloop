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

func TestCombinations2Slice(t *testing.T) {
	values := []int{3, 1, 4, 2}
	wantKeys := [][]int{
		{0, 1, 2},
		{0, 1, 3},
		{0, 2, 3},
		{1, 2, 3},
	}
	wantCombs := [][]int{
		{3, 1, 4},
		{3, 1, 2},
		{3, 4, 2},
		{1, 4, 2},
	}
	i := 0

	for seq := range gloop.Combinations2(gloop.Slice2(values), 3) {
		keys, comb := gloop.ToSlice2(seq)
		require.Equal(t, wantKeys[i], keys)
		require.Equal(t, wantCombs[i], comb)
		i++
	}

	require.Equal(t, len(wantCombs), i)
}

func TestCombinations2String(t *testing.T) {
	s := "ABCD"
	wantKeys := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 2},
		{1, 3},
		{2, 3},
	}
	wantCombs := []string{"AB", "AC", "AD", "BC", "BD", "CD"}
	i := 0

	for seq := range gloop.Combinations2(gloop.String2(s), 2) {
		keys, combRunes := gloop.ToSlice2(seq)
		comb := string(combRunes)
		require.Equal(t, wantKeys[i], keys)
		require.Equal(t, wantCombs[i], comb)
		i++
	}

	require.Equal(t, len(wantCombs), i)
}

func TestCombinations2Break(t *testing.T) {
	values := []int{3, 1, 4, 2}
	i := 0

	for seq := range gloop.Combinations2(gloop.Slice2(values), 2) {
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

func TestCombinations2ZeroSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.Combinations2(gloop.Slice2([]int{3, 1, 4}), 0) {
			t.Fatal("expected no iteration")
		}
	})
}

func TestCombinations2NegativeSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.Combinations2(gloop.Slice2([]int{3, 1, 4}), -3) {
			t.Fatal("expected no iteration")
		}
	})
}
