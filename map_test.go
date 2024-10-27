package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	m := map[int]string{
		0: "Fizz",
		1: "Buzz",
		2: "Bazz",
	}
	wantValues := []string{"Fizz", "Buzz", "Bazz"}

	i := 0
	for value := range gloop.Map(m) {
		require.Contains(t, wantValues, value)
		i++
	}

	require.Equal(t, len(m), i)
}

func TestMapBreak(t *testing.T) {
	m := map[int]string{
		0: "Fizz",
		1: "Buzz",
		2: "Bazz",
	}
	wantValues := []string{"Fizz", "Buzz", "Bazz"}

	i := 0
	for value := range gloop.Map(m) {
		if i == 2 {
			break
		}

		require.Contains(t, wantValues, value)
		i++
	}

	require.Equal(t, 2, i)
}

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
