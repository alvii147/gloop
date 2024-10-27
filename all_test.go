package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestAllTrue(t *testing.T) {
	values := []bool{true, true, true}
	require.True(t, gloop.All(gloop.Slice(values)))
}

func TestAllMixed(t *testing.T) {
	values := []bool{false, false, true, false}
	require.False(t, gloop.All(gloop.Slice(values)))
}

func TestAllFalse(t *testing.T) {
	values := []bool{false, false, false, false}
	require.False(t, gloop.All(gloop.Slice(values)))
}

func TestAll2True(t *testing.T) {
	m := map[string]bool{
		"Fizz": true,
		"Buzz": true,
		"Bazz": true,
	}
	require.True(t, gloop.All2(gloop.Map2(m)))
}

func TestAll2Mixed(t *testing.T) {
	m := map[string]bool{
		"Fizz": true,
		"Buzz": false,
		"Bazz": true,
	}
	require.False(t, gloop.All2(gloop.Map2(m)))
}

func TestAll2False(t *testing.T) {
	m := map[string]bool{
		"Fizz": false,
		"Buzz": false,
		"Bazz": false,
	}
	require.False(t, gloop.All2(gloop.Map2(m)))
}
