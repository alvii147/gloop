package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestKeys(t *testing.T) {
	values := []string{"a", "b", "c"}
	i := 0

	for value := range gloop.Keys(gloop.Enumerate(gloop.Slice(values))) {
		require.Equal(t, i, value)

		i++
	}

	require.Equal(t, len(values), i)
}

func TestKeysBreak(t *testing.T) {
	values := []string{"a", "b", "c"}
	i := 0

	for value := range gloop.Keys(gloop.Enumerate(gloop.Slice(values))) {
		if i == 2 {
			break
		}

		require.Equal(t, i, value)

		i++
	}
}
