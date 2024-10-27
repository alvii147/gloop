package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestMap2(t *testing.T) {
	m := map[string]bool{
		"Fizz": true,
		"Buzz": true,
		"Bazz": true,
	}

	i := 0
	for key, value := range gloop.Map2(m) {
		require.Contains(t, m, key)
		require.Equal(t, m[key], value)
		i++
	}

	require.Equal(t, len(m), i)
}

func TestMap2Break(t *testing.T) {
	m := map[string]bool{
		"Fizz": true,
		"Buzz": true,
		"Bazz": true,
	}

	i := 0
	for key, value := range gloop.Map2(m) {
		if i == 2 {
			break
		}

		require.Contains(t, m, key)
		require.Equal(t, m[key], value)
		i++
	}

	require.Equal(t, 2, i)
}
