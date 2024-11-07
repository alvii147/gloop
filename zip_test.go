package gloop_test

import (
	"iter"
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

func TestWithZipNPaddedTrue(t *testing.T) {
	options := gloop.ZipNOptions[int]{
		Padded: false,
	}
	gloop.WithZipNPadded[int](true)(&options)
	require.True(t, options.Padded)
}

func TestWithZipNPaddedFalse(t *testing.T) {
	options := gloop.ZipNOptions[int]{
		Padded: true,
	}
	gloop.WithZipNPadded[int](false)(&options)
	require.False(t, options.Padded)
}

func TestWithZipNPadValue(t *testing.T) {
	value := 42
	options := gloop.ZipNOptions[int]{}
	gloop.WithZipNPadValue(value)(&options)

	require.NotNil(t, options.PadValue)
	require.Equal(t, value, *options.PadValue)
}

func TestWithZipN2PaddedTrue(t *testing.T) {
	options := gloop.ZipN2Options[string, int]{
		Padded: false,
	}
	gloop.WithZipN2Padded[string, int](true)(&options)
	require.True(t, options.Padded)
}

func TestWithZipN2PaddedFalse(t *testing.T) {
	options := gloop.ZipN2Options[string, int]{
		Padded: true,
	}
	gloop.WithZipN2Padded[string, int](false)(&options)
	require.False(t, options.Padded)
}

func TestWithZipN2PadKey(t *testing.T) {
	value := "Fizz"
	options := gloop.ZipN2Options[string, int]{}
	gloop.WithZipN2PadKey[string, int](value)(&options)

	require.NotNil(t, options.PadKey)
	require.Equal(t, value, *options.PadKey)
}

func TestWithZipN2PadValue(t *testing.T) {
	value := 42
	options := gloop.ZipN2Options[string, int]{}
	gloop.WithZipN2PadValue[string](value)(&options)

	require.NotNil(t, options.PadValue)
	require.Equal(t, value, *options.PadValue)
}

func TestZipNEqualLen(t *testing.T) {
	seq1 := gloop.Slice([]int{3, 1, 4})
	seq2 := gloop.Slice([]int{1, 5, 9})
	seq3 := gloop.Slice([]int{2, 6, 5})
	wantValues := [][]int{
		{3, 1, 2},
		{1, 5, 6},
		{4, 9, 5},
	}

	i := 0
	for seq := range gloop.ZipN(gloop.Collect(seq1, seq2, seq3)) {
		require.Equal(t, wantValues[i], gloop.ToSlice(seq))
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestZipNUnequalLen(t *testing.T) {
	seq1 := gloop.Slice([]int{3, 1, 4})
	seq2 := gloop.Slice([]int{1, 5})
	seq3 := gloop.Slice([]int{2, 6, 5})
	wantValues := [][]int{
		{3, 1, 2},
		{1, 5, 6},
	}

	i := 0
	for seq := range gloop.ZipN(gloop.Collect(seq1, seq2, seq3)) {
		require.Equal(t, wantValues[i], gloop.ToSlice(seq))
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestZipNPaddedZeroValue(t *testing.T) {
	seq1 := gloop.Slice([]int{3, 1, 4})
	seq2 := gloop.Slice([]int{1, 5})
	seq3 := gloop.Slice([]int{2, 6, 5})
	wantValues := [][]int{
		{3, 1, 2},
		{1, 5, 6},
		{4, 0, 5},
	}

	i := 0
	for seq := range gloop.ZipN(
		gloop.Collect(seq1, seq2, seq3),
		gloop.WithZipNPadded[int](true),
	) {
		require.Equal(t, wantValues[i], gloop.ToSlice(seq))
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestZipNPadValue(t *testing.T) {
	seq1 := gloop.Slice([]int{3, 1, 4})
	seq2 := gloop.Slice([]int{1, 5})
	seq3 := gloop.Slice([]int{2, 6, 5})
	wantValues := [][]int{
		{3, 1, 2},
		{1, 5, 6},
		{4, -1, 5},
	}

	i := 0
	for seq := range gloop.ZipN(
		gloop.Collect(seq1, seq2, seq3),
		gloop.WithZipNPadded[int](true),
		gloop.WithZipNPadValue(-1),
	) {
		require.Equal(t, wantValues[i], gloop.ToSlice(seq))
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestZipNBreak(t *testing.T) {
	seq1 := gloop.Slice([]int{3, 1, 4})
	seq2 := gloop.Slice([]int{1, 5, 9})
	seq3 := gloop.Slice([]int{2, 6, 5})
	wantValues := [][]int{
		{3, 1, 2},
		{1, 5, 6},
	}

	i := 0
	for seq := range gloop.ZipN(gloop.Collect(seq1, seq2, seq3)) {
		if i == 2 {
			break
		}

		require.Equal(t, wantValues[i], gloop.ToSlice(seq))
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestZipN2EqualLen(t *testing.T) {
	var seq1 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Fuzz", 1) {
			return
		}

		if !yield("Fazz", 4) {
			return
		}
	}

	var seq2 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		if !yield("Bizz", 1) {
			return
		}

		if !yield("Buzz", 5) {
			return
		}

		if !yield("Bazz", 9) {
			return
		}
	}

	var seq3 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		if !yield("Dizz", 2) {
			return
		}

		if !yield("Duzz", 6) {
			return
		}

		if !yield("Dazz", 5) {
			return
		}
	}

	wantKeys := [][]string{
		{"Fizz", "Bizz", "Dizz"},
		{"Fuzz", "Buzz", "Duzz"},
		{"Fazz", "Bazz", "Dazz"},
	}
	wantValues := [][]int{
		{3, 1, 2},
		{1, 5, 6},
		{4, 9, 5},
	}

	i := 0
	for seq := range gloop.ZipN2(gloop.Collect(seq1, seq2, seq3)) {
		keys, values := gloop.ToSlice2(seq)
		require.Equal(t, wantKeys[i], keys)
		require.Equal(t, wantValues[i], values)
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestZipN2UnequalLen(t *testing.T) {
	var seq1 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Fuzz", 1) {
			return
		}

		if !yield("Fazz", 4) {
			return
		}
	}

	var seq2 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		if !yield("Bizz", 1) {
			return
		}

		if !yield("Buzz", 5) {
			return
		}
	}

	var seq3 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		if !yield("Dizz", 2) {
			return
		}

		if !yield("Duzz", 6) {
			return
		}

		if !yield("Dazz", 5) {
			return
		}
	}

	wantKeys := [][]string{
		{"Fizz", "Bizz", "Dizz"},
		{"Fuzz", "Buzz", "Duzz"},
	}
	wantValues := [][]int{
		{3, 1, 2},
		{1, 5, 6},
	}

	i := 0
	for seq := range gloop.ZipN2(gloop.Collect(seq1, seq2, seq3)) {
		keys, values := gloop.ToSlice2(seq)
		require.Equal(t, wantKeys[i], keys)
		require.Equal(t, wantValues[i], values)
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestZipN2Padded(t *testing.T) {
	var seq1 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Fuzz", 1) {
			return
		}

		if !yield("Fazz", 4) {
			return
		}
	}

	var seq2 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		if !yield("Bizz", 1) {
			return
		}

		if !yield("Buzz", 5) {
			return
		}
	}

	var seq3 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		if !yield("Dizz", 2) {
			return
		}

		if !yield("Duzz", 6) {
			return
		}

		if !yield("Dazz", 5) {
			return
		}
	}

	wantKeys := [][]string{
		{"Fizz", "Bizz", "Dizz"},
		{"Fuzz", "Buzz", "Duzz"},
		{"Fazz", "", "Dazz"},
	}
	wantValues := [][]int{
		{3, 1, 2},
		{1, 5, 6},
		{4, 0, 5},
	}

	i := 0
	for seq := range gloop.ZipN2(
		gloop.Collect(seq1, seq2, seq3),
		gloop.WithZipN2Padded[string, int](true),
	) {
		keys, values := gloop.ToSlice2(seq)
		require.Equal(t, wantKeys[i], keys)
		require.Equal(t, wantValues[i], values)
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestZipN2PadKey(t *testing.T) {
	var seq1 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Fuzz", 1) {
			return
		}

		if !yield("Fazz", 4) {
			return
		}
	}

	var seq2 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		if !yield("Bizz", 1) {
			return
		}

		if !yield("Buzz", 5) {
			return
		}
	}

	var seq3 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		if !yield("Dizz", 2) {
			return
		}

		if !yield("Duzz", 6) {
			return
		}

		if !yield("Dazz", 5) {
			return
		}
	}

	wantKeys := [][]string{
		{"Fizz", "Bizz", "Dizz"},
		{"Fuzz", "Buzz", "Duzz"},
		{"Fazz", "Zapp", "Dazz"},
	}
	wantValues := [][]int{
		{3, 1, 2},
		{1, 5, 6},
		{4, 0, 5},
	}

	i := 0
	for seq := range gloop.ZipN2(
		gloop.Collect(seq1, seq2, seq3),
		gloop.WithZipN2Padded[string, int](true),
		gloop.WithZipN2PadKey[string, int]("Zapp"),
	) {
		keys, values := gloop.ToSlice2(seq)
		require.Equal(t, wantKeys[i], keys)
		require.Equal(t, wantValues[i], values)
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestZipN2PadValue(t *testing.T) {
	var seq1 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Fuzz", 1) {
			return
		}

		if !yield("Fazz", 4) {
			return
		}
	}

	var seq2 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		if !yield("Bizz", 1) {
			return
		}

		if !yield("Buzz", 5) {
			return
		}
	}

	var seq3 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		if !yield("Dizz", 2) {
			return
		}

		if !yield("Duzz", 6) {
			return
		}

		if !yield("Dazz", 5) {
			return
		}
	}

	wantKeys := [][]string{
		{"Fizz", "Bizz", "Dizz"},
		{"Fuzz", "Buzz", "Duzz"},
		{"Fazz", "", "Dazz"},
	}
	wantValues := [][]int{
		{3, 1, 2},
		{1, 5, 6},
		{4, -1, 5},
	}

	i := 0
	for seq := range gloop.ZipN2(
		gloop.Collect(seq1, seq2, seq3),
		gloop.WithZipN2Padded[string, int](true),
		gloop.WithZipN2PadValue[string](-1),
	) {
		keys, values := gloop.ToSlice2(seq)
		require.Equal(t, wantKeys[i], keys)
		require.Equal(t, wantValues[i], values)
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestZipN2PadKeyPadValue(t *testing.T) {
	var seq1 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Fuzz", 1) {
			return
		}

		if !yield("Fazz", 4) {
			return
		}
	}

	var seq2 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		if !yield("Bizz", 1) {
			return
		}

		if !yield("Buzz", 5) {
			return
		}
	}

	var seq3 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		if !yield("Dizz", 2) {
			return
		}

		if !yield("Duzz", 6) {
			return
		}

		if !yield("Dazz", 5) {
			return
		}
	}

	wantKeys := [][]string{
		{"Fizz", "Bizz", "Dizz"},
		{"Fuzz", "Buzz", "Duzz"},
		{"Fazz", "Zapp", "Dazz"},
	}
	wantValues := [][]int{
		{3, 1, 2},
		{1, 5, 6},
		{4, -1, 5},
	}

	i := 0
	for seq := range gloop.ZipN2(
		gloop.Collect(seq1, seq2, seq3),
		gloop.WithZipN2Padded[string, int](true),
		gloop.WithZipN2PadKey[string, int]("Zapp"),
		gloop.WithZipN2PadValue[string](-1),
	) {
		keys, values := gloop.ToSlice2(seq)
		require.Equal(t, wantKeys[i], keys)
		require.Equal(t, wantValues[i], values)
		i++
	}

	require.Equal(t, len(wantValues), i)
}
