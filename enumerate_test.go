package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestEnumerate(t *testing.T) {
	values := []string{"a", "b", "c"}
	i := 0

	for idx, value := range gloop.Enumerate(gloop.Slice(values)) {
		require.Equal(t, i, idx)
		require.Equal(t, values[i], value)

		i++
	}

	require.Equal(t, len(values), i)
}

func TestEnumerateBreak(t *testing.T) {
	values := []string{"a", "b", "c"}
	i := 0

	for idx, value := range gloop.Enumerate(gloop.Slice(values)) {
		if i == 2 {
			break
		}

		require.Equal(t, i, idx)
		require.Equal(t, values[i], value)

		i++
	}
}
