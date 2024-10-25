package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestWithZipPadded(t *testing.T) {
	options := gloop.ZipOptions[string, int]{
		Padded: false,
	}
	gloop.WithZipPadded[string, int](true)(&options)
	require.True(t, options.Padded)
}

func TestWithZipPadValue1(t *testing.T) {
	value := "0xDEADBEEF"
	options := gloop.ZipOptions[string, int]{}
	gloop.WithZipPadValue1[string, int](value)(&options)

	require.NotNil(t, options.PadValue1)
	require.Equal(t, value, *options.PadValue1)
}

func TestWithZipPadValue2(t *testing.T) {
	value := 42
	options := gloop.ZipOptions[string, int]{}
	gloop.WithZipPadValue2[string](value)(&options)

	require.NotNil(t, options.PadValue2)
	require.Equal(t, value, *options.PadValue2)
}

func TestZipEqualLen(t *testing.T) {
	values1 := []string{"a", "b", "c"}
	values2 := []int{1, 2, 3}
	i := 0

	for value1, value2 := range gloop.Zip(gloop.Slice(values1), gloop.Slice(values2)) {
		require.Equal(t, values1[i], value1)
		require.Equal(t, values2[i], value2)
		i++
	}

	require.Equal(t, 3, i)
}

func TestZipUnequalLen(t *testing.T) {
	values1 := []string{"a", "b", "c"}
	values2 := []int{1, 2, 3, 4}
	i := 0

	for value1, value2 := range gloop.Zip(gloop.Slice(values1), gloop.Slice(values2)) {
		require.Equal(t, values1[i], value1)
		require.Equal(t, values2[i], value2)
		i++
	}

	require.Equal(t, 3, i)
}

func TestZipPadZeroString(t *testing.T) {
	values1 := []string{"a", "b", "c"}
	values2 := []int{1, 2, 3, 4, 5}
	wantValues1 := []string{"a", "b", "c", "", ""}
	wantValues2 := []int{1, 2, 3, 4, 5}
	i := 0

	for value1, value2 := range gloop.Zip(
		gloop.Slice(values1),
		gloop.Slice(values2),
		gloop.WithZipPadded[string, int](true),
	) {
		require.Equal(t, wantValues1[i], value1)
		require.Equal(t, wantValues2[i], value2)
		i++
	}

	require.Equal(t, 5, i)
}

func TestZipPadZeroInt(t *testing.T) {
	values1 := []string{"a", "b", "c", "d", "e"}
	values2 := []int{1, 2, 3}
	wantValues1 := []string{"a", "b", "c", "d", "e"}
	wantValues2 := []int{1, 2, 3, 0, 0}
	i := 0

	for value1, value2 := range gloop.Zip(
		gloop.Slice(values1),
		gloop.Slice(values2),
		gloop.WithZipPadded[string, int](true),
	) {
		require.Equal(t, wantValues1[i], value1)
		require.Equal(t, wantValues2[i], value2)
		i++
	}

	require.Equal(t, 5, i)
}

func TestZipPadValue1(t *testing.T) {
	values1 := []string{"a", "b", "c"}
	values2 := []int{1, 2, 3, 4, 5}
	wantValues1 := []string{"a", "b", "c", "x", "x"}
	wantValues2 := []int{1, 2, 3, 4, 5}
	i := 0

	for value1, value2 := range gloop.Zip(
		gloop.Slice(values1),
		gloop.Slice(values2),
		gloop.WithZipPadded[string, int](true),
		gloop.WithZipPadValue1[string, int]("x"),
	) {
		require.Equal(t, wantValues1[i], value1)
		require.Equal(t, wantValues2[i], value2)
		i++
	}

	require.Equal(t, 5, i)
}

func TestZipPadValue2(t *testing.T) {
	values1 := []string{"a", "b", "c", "d", "e"}
	values2 := []int{1, 2, 3}
	wantValues1 := []string{"a", "b", "c", "d", "e"}
	wantValues2 := []int{1, 2, 3, 42, 42}
	i := 0

	for value1, value2 := range gloop.Zip(
		gloop.Slice(values1),
		gloop.Slice(values2),
		gloop.WithZipPadded[string, int](true),
		gloop.WithZipPadValue2[string](42),
	) {
		require.Equal(t, wantValues1[i], value1)
		require.Equal(t, wantValues2[i], value2)
		i++
	}

	require.Equal(t, 5, i)
}

func TestZipBreak(t *testing.T) {
	values1 := []string{"a", "b", "c"}
	values2 := []int{1, 2, 3}
	i := 0

	for value1, value2 := range gloop.Zip(gloop.Slice(values1), gloop.Slice(values2)) {
		if i == 2 {
			break
		}

		require.Equal(t, values1[i], value1)
		require.Equal(t, values2[i], value2)
		i++
	}

	require.Equal(t, 2, i)
}
