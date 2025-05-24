package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestChainSlices(t *testing.T) {
	values1 := []string{"a", "b", "c"}
	values2 := []string{"x", "y", "z"}
	values3 := []string{"1", "2", "3"}
	chainedValues := []string{"a", "b", "c", "x", "y", "z", "1", "2", "3"}
	i := 0

	for value := range gloop.Chain(
		gloop.Slice(values1),
		gloop.Slice(values2),
		gloop.Slice(values3),
	) {
		require.Equal(t, chainedValues[i], value)

		i++
	}

	require.Equal(t, len(chainedValues), i)
}

func TestChainStrings(t *testing.T) {
	s1 := "Fizz"
	s2 := "Buzz"
	chainedRunes := []rune{'F', 'i', 'z', 'z', 'B', 'u', 'z', 'z'}
	i := 0

	for value := range gloop.Chain(gloop.String(s1), gloop.String(s2)) {
		require.Equal(t, chainedRunes[i], value)

		i++
	}

	require.Equal(t, len(chainedRunes), i)
}

func TestChainSliceAndString(t *testing.T) {
	values := []rune{'F', 'i', 'z', 'z'}
	s := "Buzz"
	chainedRunes := []rune{'F', 'i', 'z', 'z', 'B', 'u', 'z', 'z'}
	i := 0

	for value := range gloop.Chain(gloop.Slice(values), gloop.String(s)) {
		require.Equal(t, chainedRunes[i], value)

		i++
	}

	require.Equal(t, len(chainedRunes), i)
}

func TestChainBreak(t *testing.T) {
	values := []string{"a", "b", "c"}
	chainedValues := []string{"a", "b"}
	i := 0

	for value := range gloop.Chain(gloop.Slice(values)) {
		if i == 2 {
			break
		}

		require.Equal(t, chainedValues[i], value)

		i++
	}
}

func TestChain2Slices(t *testing.T) {
	values1 := []string{"a", "b", "c"}
	values2 := []string{"x", "y", "z"}
	values3 := []string{"1", "2", "3"}
	chainedIdx := []int{0, 1, 2, 0, 1, 2, 0, 1, 2}
	chainedValues := []string{"a", "b", "c", "x", "y", "z", "1", "2", "3"}
	i := 0

	for idx, value := range gloop.Chain2(
		gloop.Enumerate(gloop.Slice(values1)),
		gloop.Enumerate(gloop.Slice(values2)),
		gloop.Enumerate(gloop.Slice(values3)),
	) {
		require.Equal(t, chainedIdx[i], idx)
		require.Equal(t, chainedValues[i], value)

		i++
	}

	require.Equal(t, len(chainedValues), i)
}

func TestChain2Strings(t *testing.T) {
	s1 := "Fizz"
	s2 := "Buzz"
	chainedIdx := []int{0, 1, 2, 3, 0, 1, 2, 3}
	chainedRunes := []rune{'F', 'i', 'z', 'z', 'B', 'u', 'z', 'z'}
	i := 0

	for idx, value := range gloop.Chain2(
		gloop.Enumerate(gloop.String(s1)),
		gloop.Enumerate(gloop.String(s2)),
	) {
		require.Equal(t, chainedIdx[i], idx)
		require.Equal(t, chainedRunes[i], value)

		i++
	}

	require.Equal(t, len(chainedRunes), i)
}

func TestChain2SliceAndString(t *testing.T) {
	values := []rune{'F', 'i', 'z', 'z'}
	s := "Buzz"
	chainedIdx := []int{0, 1, 2, 3, 0, 1, 2, 3}
	chainedRunes := []rune{'F', 'i', 'z', 'z', 'B', 'u', 'z', 'z'}
	i := 0

	for idx, value := range gloop.Chain2(
		gloop.Enumerate(gloop.Slice(values)),
		gloop.Enumerate(gloop.String(s)),
	) {
		require.Equal(t, chainedIdx[i], idx)
		require.Equal(t, chainedRunes[i], value)

		i++
	}

	require.Equal(t, len(chainedRunes), i)
}

func TestChain2Break(t *testing.T) {
	values := []string{"a", "b", "c"}
	chainedIdx := []int{0, 1, 2}
	chainedValues := []string{"a", "b"}
	i := 0

	for idx, value := range gloop.Chain2(gloop.Enumerate(gloop.Slice(values))) {
		if i == 2 {
			break
		}

		require.Equal(t, chainedIdx[i], idx)
		require.Equal(t, chainedValues[i], value)

		i++
	}
}
