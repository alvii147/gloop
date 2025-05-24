package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	m := map[string]bool{
		"Fizz": true,
		"Buzz": true,
		"Bazz": true,
	}

	i := 0

	for key, value := range gloop.Map(m) {
		require.Contains(t, m, key)
		require.Equal(t, m[key], value)

		i++
	}

	require.Equal(t, len(m), i)
}

func TestMapBreak(t *testing.T) {
	m := map[string]bool{
		"Fizz": true,
		"Buzz": true,
		"Bazz": true,
	}

	i := 0
	for key, value := range gloop.Map(m) {
		if i == 2 {
			break
		}

		require.Contains(t, m, key)
		require.Equal(t, m[key], value)

		i++
	}

	require.Equal(t, 2, i)
}
