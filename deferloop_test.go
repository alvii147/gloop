package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestDeferLoop(t *testing.T) {
	values := []int{3, 1, 4}
	operations := make([]string, 0)
	wantOperations := []string{"Fizz", "Buzz", "Fizz", "Buzz", "Fizz", "Buzz"}

	i := 0

	for value, deferLoop := range gloop.DeferLoop(gloop.Slice(values)) {
		deferLoop(func() {
			operations = append(operations, "Buzz")
		})

		operations = append(operations, "Fizz")

		require.Equal(t, values[i], value)

		i++
	}

	require.Equal(t, wantOperations, operations)
}

func TestDeferLoopMultiple(t *testing.T) {
	values := []int{3, 1, 4}
	operations := make([]string, 0)
	wantOperations := []string{"Fizz", "Buzz", "Bazz", "Fizz", "Buzz", "Bazz", "Fizz", "Buzz", "Bazz"}

	i := 0

	for value, deferLoop := range gloop.DeferLoop(gloop.Slice(values)) {
		deferLoop(func() {
			operations = append(operations, "Bazz")
		})

		deferLoop(func() {
			operations = append(operations, "Buzz")
		})

		operations = append(operations, "Fizz")

		require.Equal(t, values[i], value)

		i++
	}

	require.Equal(t, wantOperations, operations)
}

func TestDeferLoopNoDefer(t *testing.T) {
	values := []int{3, 1, 4}
	operations := make([]string, 0)
	wantOperations := []string{"Fizz", "Fizz", "Fizz"}

	i := 0

	for value := range gloop.DeferLoop(gloop.Slice(values)) {
		operations = append(operations, "Fizz")

		require.Equal(t, values[i], value)

		i++
	}

	require.Equal(t, wantOperations, operations)
}

func TestDeferLoopDeferOnce(t *testing.T) {
	values := []int{3, 1, 4}
	operations := make([]string, 0)
	wantOperations := []string{"Fizz", "Fizz", "Buzz", "Fizz"}

	i := 0
	for value, deferLoop := range gloop.DeferLoop(gloop.Slice(values)) {
		if i == 1 {
			deferLoop(func() {
				operations = append(operations, "Buzz")
			})
		}

		operations = append(operations, "Fizz")

		require.Equal(t, values[i], value)

		i++
	}

	require.Equal(t, wantOperations, operations)
}

func TestDeferLoopBreak(t *testing.T) {
	values := []int{3, 1, 4}
	operations := make([]string, 0)
	wantOperations := []string{"Fizz", "Buzz"}

	for value, deferLoop := range gloop.DeferLoop(gloop.Slice(values)) {
		deferLoop(func() {
			operations = append(operations, "Buzz")
		})

		operations = append(operations, "Fizz")

		require.Equal(t, 3, value)

		break
	}

	require.Equal(t, wantOperations, operations)
}

func TestDeferLoopNilSafety(t *testing.T) {
	values := []int{3, 1, 4}
	operations := make([]string, 0)
	wantOperations := []string{"Fizz", "Fizz", "Fizz"}

	i := 0

	for value, deferLoop := range gloop.DeferLoop(gloop.Slice(values)) {
		deferLoop(nil)

		operations = append(operations, "Fizz")

		require.Equal(t, values[i], value)

		i++
	}

	require.Equal(t, wantOperations, operations)
}
