package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestBatchSliceDivisibleLength(t *testing.T) {
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	wantBatches := [][]int{
		{3, 1, 4},
		{1, 5, 9},
		{2, 6, 5},
	}
	i := 0

	for seq := range gloop.Batch(gloop.Slice(values), 3) {
		batch := gloop.ToSlice(seq)
		require.Equal(t, wantBatches[i], batch)
		i++
	}

	require.Equal(t, len(wantBatches), i)
}

func TestBatchSliceIndivisibleLength(t *testing.T) {
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	wantBatches := [][]int{
		{3, 1, 4, 1},
		{5, 9, 2, 6},
		{5},
	}
	i := 0

	for seq := range gloop.Batch(gloop.Slice(values), 4) {
		batch := gloop.ToSlice(seq)
		require.Equal(t, wantBatches[i], batch)
		i++
	}

	require.Equal(t, len(wantBatches), i)
}

func TestBatchStringDivisibleLength(t *testing.T) {
	s := "FizzBuzz"
	wantBatches := []string{
		"Fizz",
		"Buzz",
	}
	i := 0

	for seq := range gloop.Batch(gloop.String(s), 4) {
		batch := gloop.ToString(seq)
		require.Equal(t, wantBatches[i], batch)
		i++
	}

	require.Equal(t, len(wantBatches), i)
}

func TestBatchStringIndivisibleLength(t *testing.T) {
	s := "FizzBuzz"
	wantBatches := []string{
		"Fiz",
		"zBu",
		"zz",
	}
	i := 0

	for seq := range gloop.Batch(gloop.String(s), 3) {
		batch := gloop.ToString(seq)
		require.Equal(t, wantBatches[i], batch)
		i++
	}

	require.Equal(t, len(wantBatches), i)
}

func TestBatchBreak(t *testing.T) {
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	i := 0

	for seq := range gloop.Batch(gloop.Slice(values), 3) {
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

func TestBatchZeroSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.Batch(gloop.Slice([]int{3, 1, 4}), 0) {
			t.Fatal("expected no iteration")
		}
	})
}

func TestBatchNegativeSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.Batch(gloop.Slice([]int{3, 1, 4}), -1) {
			t.Fatal("expected no iteration")
		}
	})
}
