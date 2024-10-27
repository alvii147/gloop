package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestValues(t *testing.T) {
	values := []string{"a", "b", "c"}
	i := 0

	for value := range gloop.Values(gloop.Slice2(values)) {
		require.Equal(t, values[i], value)
		i++
	}

	require.Equal(t, len(values), i)
}

func TestValuesBreak(t *testing.T) {
	values := []string{"a", "b", "c"}
	i := 0

	for value := range gloop.Values(gloop.Slice2(values)) {
		if i == 2 {
			break
		}

		require.Equal(t, values[i], value)
		i++
	}
}
