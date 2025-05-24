package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestCartesianProductSlice(t *testing.T) {
	values := []int{3, 1, 4}
	wantCartesianProducts := [][]int{
		{3, 3},
		{3, 1},
		{3, 4},
		{1, 3},
		{1, 1},
		{1, 4},
		{4, 3},
		{4, 1},
		{4, 4},
	}
	i := 0

	for seq := range gloop.CartesianProduct(gloop.Slice(values), 2) {
		product := gloop.ToSlice(seq)
		require.Equal(t, wantCartesianProducts[i], product)

		i++
	}

	require.Equal(t, len(wantCartesianProducts), i)
}

func TestCartesianProductString(t *testing.T) {
	s := "ABCD"
	wantCartesianProducts := []string{
		"AA", "AB", "AC", "AD",
		"BA", "BB", "BC", "BD",
		"CA", "CB", "CC", "CD",
		"DA", "DB", "DC", "DD",
	}
	i := 0

	for seq := range gloop.CartesianProduct(gloop.String(s), 2) {
		product := gloop.ToString(seq)
		require.Equal(t, wantCartesianProducts[i], product)

		i++
	}

	require.Equal(t, len(wantCartesianProducts), i)
}

func TestCartesianProductBreak(t *testing.T) {
	values := []int{3, 1, 4}
	i := 0

	for seq := range gloop.CartesianProduct(gloop.Slice(values), 2) {
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

func TestCartesianProductZeroSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.CartesianProduct(gloop.Slice([]int{3, 1, 4}), 0) {
			t.Fatal("expected no iteration")
		}
	})
}

func TestCartesianProductNegativeSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.CartesianProduct(gloop.Slice([]int{3, 1, 4}), -3) {
			t.Fatal("expected no iteration")
		}
	})
}

func TestCartesianProduct2Slice(t *testing.T) {
	values := []int{3, 1, 4}
	wantKeys := [][]int{
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 0},
		{2, 1},
		{2, 2},
	}
	wantCartesianProducts := [][]int{
		{3, 3},
		{3, 1},
		{3, 4},
		{1, 3},
		{1, 1},
		{1, 4},
		{4, 3},
		{4, 1},
		{4, 4},
	}
	i := 0

	for seq := range gloop.CartesianProduct2(gloop.Enumerate(gloop.Slice(values)), 2) {
		keys, product := gloop.ToSlice2(seq)
		require.Equal(t, wantKeys[i], keys)
		require.Equal(t, wantCartesianProducts[i], product)

		i++
	}

	require.Equal(t, len(wantCartesianProducts), i)
}

func TestCartesianProduct2String(t *testing.T) {
	s := "ABCD"
	wantKeys := [][]int{
		{0, 0}, {0, 1}, {0, 2}, {0, 3},
		{1, 0}, {1, 1}, {1, 2}, {1, 3},
		{2, 0}, {2, 1}, {2, 2}, {2, 3},
		{3, 0}, {3, 1}, {3, 2}, {3, 3},
	}
	wantCartesianProducts := []string{
		"AA", "AB", "AC", "AD",
		"BA", "BB", "BC", "BD",
		"CA", "CB", "CC", "CD",
		"DA", "DB", "DC", "DD",
	}
	i := 0

	for seq := range gloop.CartesianProduct2(gloop.Enumerate(gloop.String(s)), 2) {
		keys, productRunes := gloop.ToSlice2(seq)
		product := string(productRunes)

		require.Equal(t, wantKeys[i], keys)
		require.Equal(t, wantCartesianProducts[i], product)

		i++
	}

	require.Equal(t, len(wantCartesianProducts), i)
}

func TestCartesianProduct2Break(t *testing.T) {
	values := []int{3, 1, 4}
	i := 0

	for seq := range gloop.CartesianProduct2(gloop.Enumerate(gloop.Slice(values)), 2) {
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

func TestCartesianProduct2ZeroSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.CartesianProduct2(gloop.Enumerate(gloop.Slice([]int{3, 1, 4})), 0) {
			t.Fatal("expected no iteration")
		}
	})
}

func TestCartesianProduct2NegativeSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.CartesianProduct2(gloop.Enumerate(gloop.Slice([]int{3, 1, 4})), -3) {
			t.Fatal("expected no iteration")
		}
	})
}
