package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestSortAscendingInt(t *testing.T) {
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	wantValues := []int{1, 1, 2, 3, 4, 5, 5, 6, 9}

	i := 0
	for value := range gloop.Sort(gloop.Slice(values), true) {
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestSortDescendingInt(t *testing.T) {
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	wantValues := []int{9, 6, 5, 5, 4, 3, 2, 1, 1}

	i := 0
	for value := range gloop.Sort(gloop.Slice(values), false) {
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestSortAscendingFloat(t *testing.T) {
	values := []float64{3.14, -1.59, 2.65}
	wantValues := []float64{-1.59, 2.65, 3.14}

	i := 0
	for value := range gloop.Sort(gloop.Slice(values), true) {
		require.InDelta(t, wantValues[i], value, 0.001)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestSortDescendingFloat(t *testing.T) {
	values := []float64{3.14, -1.59, 2.65}
	wantValues := []float64{3.14, 2.65, -1.59}

	i := 0
	for value := range gloop.Sort(gloop.Slice(values), false) {
		require.InDelta(t, wantValues[i], value, 0.001)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestSortAscendingString(t *testing.T) {
	values := []string{"Fizz", "Buzz", "Bazz"}
	wantValues := []string{"Bazz", "Buzz", "Fizz"}

	i := 0
	for value := range gloop.Sort(gloop.Slice(values), true) {
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestSortDescendingString(t *testing.T) {
	values := []string{"Fizz", "Buzz", "Bazz"}
	wantValues := []string{"Fizz", "Buzz", "Bazz"}

	i := 0
	for value := range gloop.Sort(gloop.Slice(values), false) {
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestSortBreak(t *testing.T) {
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	wantValues := []int{1, 1, 2}

	i := 0
	for value := range gloop.Sort(gloop.Slice(values), true) {
		if i == 3 {
			break
		}

		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestSortByComparisonAscending(t *testing.T) {
	values := []string{"Fizzzzzz", "Buz", "Bazzz"}
	wantValues := []string{"Buz", "Bazzz", "Fizzzzzz"}

	i := 0
	for value := range gloop.SortByComparison(
		gloop.Slice(values),
		func(s1 string, s2 string) bool {
			return len(s1) < len(s2)
		},
		true,
	) {
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestSortByComparisonDescending(t *testing.T) {
	values := []string{"Fizzzzzz", "Buz", "Bazzz"}
	wantValues := []string{"Fizzzzzz", "Bazzz", "Buz"}

	i := 0
	for value := range gloop.SortByComparison(
		gloop.Slice(values),
		func(s1 string, s2 string) bool {
			return len(s1) < len(s2)
		},
		false,
	) {
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestSortByComparisonBreak(t *testing.T) {
	values := []string{"Fizzzzzz", "Buz", "Bazzz"}
	wantValues := []string{"Buz", "Bazzz"}

	i := 0
	for value := range gloop.SortByComparison(
		gloop.Slice(values),
		func(s1 string, s2 string) bool {
			return len(s1) < len(s2)
		},
		true,
	) {
		if i == 2 {
			break
		}

		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestSortByComparison2Ascending(t *testing.T) {
	m := map[int]int{
		3:  -1,
		1:  5,
		-4: -9,
	}
	wantKeys := []int{3, 1, -4}
	wantValues := []int{-1, 5, -9}

	i := 0
	for key, value := range gloop.SortByComparison2(
		gloop.Map(m),
		func(key1 int, value1 int, key2 int, value2 int) bool {
			return key1*value1 < key2*value2
		},
		true,
	) {
		require.Equal(t, wantKeys[i], key)
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantKeys), i)
}

func TestSortByComparison2Descending(t *testing.T) {
	m := map[int]int{
		3:  -1,
		1:  5,
		-4: -9,
	}
	wantKeys := []int{-4, 1, 3}
	wantValues := []int{-9, 5, -1}

	i := 0
	for key, value := range gloop.SortByComparison2(
		gloop.Map(m),
		func(key1 int, value1 int, key2 int, value2 int) bool {
			return key1*value1 < key2*value2
		},
		false,
	) {
		require.Equal(t, wantKeys[i], key)
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantKeys), i)
}

func TestSortByComparison2Break(t *testing.T) {
	m := map[int]int{
		3:  -1,
		1:  5,
		-4: -9,
	}
	wantKeys := []int{3, 1}
	wantValues := []int{-1, 5}

	i := 0
	for key, value := range gloop.SortByComparison2(
		gloop.Map(m),
		func(key1 int, value1 int, key2 int, value2 int) bool {
			return key1*value1 < key2*value2
		},
		true,
	) {
		if i == 2 {
			break
		}

		require.Equal(t, wantKeys[i], key)
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, 2, i)
}

func TestSortByRankAscending(t *testing.T) {
	values := []string{"Fizzzzzz", "Buz", "Bazzz"}
	wantValues := []string{"Buz", "Bazzz", "Fizzzzzz"}

	i := 0
	for value := range gloop.SortByRank(
		gloop.Slice(values),
		func(s string) int {
			return len(s)
		},
		true,
	) {
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestSortByRankDescending(t *testing.T) {
	values := []string{"Fizzzzzz", "Buz", "Bazzz"}
	wantValues := []string{"Fizzzzzz", "Bazzz", "Buz"}

	i := 0
	for value := range gloop.SortByRank(
		gloop.Slice(values),
		func(s string) int {
			return len(s)
		},
		false,
	) {
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestSortByRankBreak(t *testing.T) {
	values := []string{"Fizzzzzz", "Buz", "Bazzz"}
	wantValues := []string{"Buz", "Bazzz"}

	i := 0
	for value := range gloop.SortByRank(
		gloop.Slice(values),
		func(s string) int {
			return len(s)
		},
		true,
	) {
		if i == 2 {
			break
		}

		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestSortByRank2Ascending(t *testing.T) {
	m := map[int]int{
		3:  -1,
		1:  5,
		-4: -9,
	}
	wantKeys := []int{3, 1, -4}
	wantValues := []int{-1, 5, -9}

	i := 0
	for key, value := range gloop.SortByRank2(
		gloop.Map(m),
		func(key int, value int) int {
			return key * value
		},
		true,
	) {
		require.Equal(t, wantKeys[i], key)
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantKeys), i)
}

func TestSortByRank2Descending(t *testing.T) {
	m := map[int]int{
		3:  -1,
		1:  5,
		-4: -9,
	}
	wantKeys := []int{-4, 1, 3}
	wantValues := []int{-9, 5, -1}

	i := 0
	for key, value := range gloop.SortByRank2(
		gloop.Map(m),
		func(key int, value int) int {
			return key * value
		},
		false,
	) {
		require.Equal(t, wantKeys[i], key)
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, len(wantKeys), i)
}

func TestSortByRank2Break(t *testing.T) {
	m := map[int]int{
		3:  -1,
		1:  5,
		-4: -9,
	}
	wantKeys := []int{3, 1}
	wantValues := []int{-1, 5}

	i := 0
	for key, value := range gloop.SortByRank2(
		gloop.Map(m),
		func(key int, value int) int {
			return key * value
		},
		true,
	) {
		if i == 2 {
			break
		}

		require.Equal(t, wantKeys[i], key)
		require.Equal(t, wantValues[i], value)

		i++
	}

	require.Equal(t, 2, i)
}
