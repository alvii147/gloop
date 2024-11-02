package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestWithZipPaddedTrue(t *testing.T) {
	options := gloop.ZipOptions[string, int]{
		Padded: false,
	}
	gloop.WithZipPadded[string, int](true)(&options)
	require.True(t, options.Padded)
}

func TestWithZipPaddedFalse(t *testing.T) {
	options := gloop.ZipOptions[string, int]{
		Padded: true,
	}
	gloop.WithZipPadded[string, int](false)(&options)
	require.False(t, options.Padded)
}

func TestWithZipPadValue1(t *testing.T) {
	value := "FizzBuzz"
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

func TestZipPaddedZeroString(t *testing.T) {
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

func TestZipPaddedZeroInt(t *testing.T) {
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

func TestWithZip2PaddedTrue(t *testing.T) {
	options := gloop.Zip2Options[string, int, rune, float64]{
		Padded: false,
	}
	gloop.WithZip2Padded[string, int, rune, float64](true)(&options)
	require.True(t, options.Padded)
}

func TestWithZip2PaddedFalse(t *testing.T) {
	options := gloop.Zip2Options[string, int, rune, float64]{
		Padded: true,
	}
	gloop.WithZip2Padded[string, int, rune, float64](false)(&options)
	require.False(t, options.Padded)
}

func TestWithZip2PadKey1(t *testing.T) {
	key := "FizzBuzz"
	options := gloop.Zip2Options[string, int, rune, float64]{}
	gloop.WithZip2PadKey1[string, int, rune, float64](key)(&options)

	require.NotNil(t, options.PadKey1)
	require.Equal(t, key, *options.PadKey1)
}

func TestWithZip2PadValue1(t *testing.T) {
	value := 42
	options := gloop.Zip2Options[string, int, rune, float64]{}
	gloop.WithZip2PadValue1[string, int, rune, float64](value)(&options)

	require.NotNil(t, options.PadValue1)
	require.Equal(t, value, *options.PadValue1)
}

func TestWithZip2PadKey2(t *testing.T) {
	key := 'F'
	options := gloop.Zip2Options[string, int, rune, float64]{}
	gloop.WithZip2PadKey2[string, int, rune, float64](key)(&options)

	require.NotNil(t, options.PadKey2)
	require.Equal(t, key, *options.PadKey2)
}

func TestWithZip2PadValue2(t *testing.T) {
	value := 42.0
	options := gloop.Zip2Options[string, int, rune, float64]{}
	gloop.WithZip2PadValue2[string, int, rune](value)(&options)

	require.NotNil(t, options.PadValue2)
	require.Equal(t, value, *options.PadValue2)
}

func TestZip2EqualLen(t *testing.T) {
	seq1 := func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Buzz", 1) {
			return
		}

		if !yield("Bazz", 4) {
			return
		}
	}

	seq2 := func(yield func(rune, float64) bool) {
		if !yield('a', 3.14) {
			return
		}

		if !yield('b', 1.59) {
			return
		}

		if !yield('c', 2.65) {
			return
		}
	}

	wantPairs1 := []gloop.KeyValuePair[string, int]{
		{
			Key:   "Fizz",
			Value: 3,
		},
		{
			Key:   "Buzz",
			Value: 1,
		},
		{
			Key:   "Bazz",
			Value: 4,
		},
	}

	wantPairs2 := []gloop.KeyValuePair[rune, float64]{
		{
			Key:   'a',
			Value: 3.14,
		},
		{
			Key:   'b',
			Value: 1.59,
		},
		{
			Key:   'c',
			Value: 2.65,
		},
	}

	i := 0
	for pair1, pair2 := range gloop.Zip2(seq1, seq2) {
		require.Equal(t, wantPairs1[i], pair1)
		require.Equal(t, wantPairs2[i], pair2)
		i++
	}

	require.Equal(t, 3, i)
}

func TestZip2UnequalLen(t *testing.T) {
	seq1 := func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Buzz", 1) {
			return
		}

		if !yield("Bazz", 4) {
			return
		}

		if !yield("Bizz", 2) {
			return
		}
	}

	seq2 := func(yield func(rune, float64) bool) {
		if !yield('a', 3.14) {
			return
		}

		if !yield('b', 1.59) {
			return
		}

		if !yield('c', 2.65) {
			return
		}
	}

	wantPairs1 := []gloop.KeyValuePair[string, int]{
		{
			Key:   "Fizz",
			Value: 3,
		},
		{
			Key:   "Buzz",
			Value: 1,
		},
		{
			Key:   "Bazz",
			Value: 4,
		},
	}

	wantPairs2 := []gloop.KeyValuePair[rune, float64]{
		{
			Key:   'a',
			Value: 3.14,
		},
		{
			Key:   'b',
			Value: 1.59,
		},
		{
			Key:   'c',
			Value: 2.65,
		},
	}

	i := 0
	for pair1, pair2 := range gloop.Zip2(seq1, seq2) {
		require.Equal(t, wantPairs1[i], pair1)
		require.Equal(t, wantPairs2[i], pair2)
		i++
	}

	require.Equal(t, 3, i)
}

func TestZip2PaddedZeroKeyValue1(t *testing.T) {
	seq1 := func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}
	}

	seq2 := func(yield func(rune, float64) bool) {
		if !yield('a', 3.14) {
			return
		}

		if !yield('b', 1.59) {
			return
		}
	}

	wantPairs1 := []gloop.KeyValuePair[string, int]{
		{
			Key:   "Fizz",
			Value: 3,
		},
		{
			Key:   "",
			Value: 0,
		},
	}

	wantPairs2 := []gloop.KeyValuePair[rune, float64]{
		{
			Key:   'a',
			Value: 3.14,
		},
		{
			Key:   'b',
			Value: 1.59,
		},
	}

	i := 0
	for pair1, pair2 := range gloop.Zip2(
		seq1,
		seq2,
		gloop.WithZip2Padded[string, int, rune, float64](true),
	) {
		require.Equal(t, wantPairs1[i], pair1)
		require.Equal(t, wantPairs2[i], pair2)
		i++
	}

	require.Equal(t, 2, i)
}

func TestZip2PaddedZeroKeyValue2(t *testing.T) {
	seq1 := func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Buzz", 1) {
			return
		}
	}

	seq2 := func(yield func(rune, float64) bool) {
		if !yield('a', 3.14) {
			return
		}
	}

	wantPairs1 := []gloop.KeyValuePair[string, int]{
		{
			Key:   "Fizz",
			Value: 3,
		},
		{
			Key:   "Buzz",
			Value: 1,
		},
	}

	wantPairs2 := []gloop.KeyValuePair[rune, float64]{
		{
			Key:   'a',
			Value: 3.14,
		},
		{
			Key:   0,
			Value: 0.0,
		},
	}

	i := 0
	for pair1, pair2 := range gloop.Zip2(
		seq1,
		seq2,
		gloop.WithZip2Padded[string, int, rune, float64](true),
	) {
		require.Equal(t, wantPairs1[i], pair1)
		require.Equal(t, wantPairs2[i], pair2)
		i++
	}

	require.Equal(t, 2, i)
}

func TestZip2PadKey1(t *testing.T) {
	seq1 := func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}
	}

	seq2 := func(yield func(rune, float64) bool) {
		if !yield('a', 3.14) {
			return
		}

		if !yield('b', 1.59) {
			return
		}
	}

	wantPairs1 := []gloop.KeyValuePair[string, int]{
		{
			Key:   "Fizz",
			Value: 3,
		},
		{
			Key:   "ZZZZ",
			Value: 0,
		},
	}

	wantPairs2 := []gloop.KeyValuePair[rune, float64]{
		{
			Key:   'a',
			Value: 3.14,
		},
		{
			Key:   'b',
			Value: 1.59,
		},
	}

	i := 0
	for pair1, pair2 := range gloop.Zip2(
		seq1,
		seq2,
		gloop.WithZip2Padded[string, int, rune, float64](true),
		gloop.WithZip2PadKey1[string, int, rune, float64]("ZZZZ"),
	) {
		require.Equal(t, wantPairs1[i], pair1)
		require.Equal(t, wantPairs2[i], pair2)
		i++
	}

	require.Equal(t, 2, i)
}

func TestZip2PadValue1(t *testing.T) {
	seq1 := func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}
	}

	seq2 := func(yield func(rune, float64) bool) {
		if !yield('a', 3.14) {
			return
		}

		if !yield('b', 1.59) {
			return
		}
	}

	wantPairs1 := []gloop.KeyValuePair[string, int]{
		{
			Key:   "Fizz",
			Value: 3,
		},
		{
			Key:   "",
			Value: 42,
		},
	}

	wantPairs2 := []gloop.KeyValuePair[rune, float64]{
		{
			Key:   'a',
			Value: 3.14,
		},
		{
			Key:   'b',
			Value: 1.59,
		},
	}

	i := 0
	for pair1, pair2 := range gloop.Zip2(
		seq1,
		seq2,
		gloop.WithZip2Padded[string, int, rune, float64](true),
		gloop.WithZip2PadValue1[string, int, rune, float64](42),
	) {
		require.Equal(t, wantPairs1[i], pair1)
		require.Equal(t, wantPairs2[i], pair2)
		i++
	}

	require.Equal(t, 2, i)
}

func TestZip2PadKey2(t *testing.T) {
	seq1 := func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Buzz", 1) {
			return
		}
	}

	seq2 := func(yield func(rune, float64) bool) {
		if !yield('a', 3.14) {
			return
		}
	}

	wantPairs1 := []gloop.KeyValuePair[string, int]{
		{
			Key:   "Fizz",
			Value: 3,
		},
		{
			Key:   "Buzz",
			Value: 1,
		},
	}

	wantPairs2 := []gloop.KeyValuePair[rune, float64]{
		{
			Key:   'a',
			Value: 3.14,
		},
		{
			Key:   'Z',
			Value: 0.0,
		},
	}

	i := 0
	for pair1, pair2 := range gloop.Zip2(
		seq1,
		seq2,
		gloop.WithZip2Padded[string, int, rune, float64](true),
		gloop.WithZip2PadKey2[string, int, rune, float64]('Z'),
	) {
		require.Equal(t, wantPairs1[i], pair1)
		require.Equal(t, wantPairs2[i], pair2)
		i++
	}

	require.Equal(t, 2, i)
}

func TestZip2PadValue2(t *testing.T) {
	seq1 := func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Buzz", 1) {
			return
		}
	}

	seq2 := func(yield func(rune, float64) bool) {
		if !yield('a', 3.14) {
			return
		}
	}

	wantPairs1 := []gloop.KeyValuePair[string, int]{
		{
			Key:   "Fizz",
			Value: 3,
		},
		{
			Key:   "Buzz",
			Value: 1,
		},
	}

	wantPairs2 := []gloop.KeyValuePair[rune, float64]{
		{
			Key:   'a',
			Value: 3.14,
		},
		{
			Key:   0,
			Value: 4.2,
		},
	}

	i := 0
	for pair1, pair2 := range gloop.Zip2(
		seq1,
		seq2,
		gloop.WithZip2Padded[string, int, rune, float64](true),
		gloop.WithZip2PadValue2[string, int, rune](4.2),
	) {
		require.Equal(t, wantPairs1[i], pair1)
		require.Equal(t, wantPairs2[i], pair2)
		i++
	}

	require.Equal(t, 2, i)
}
