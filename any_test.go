package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestAnyTrue(t *testing.T) {
	values := []bool{true, true, true}
	require.True(t, gloop.Any(gloop.Slice(values)))
}

func TestAnyMixed(t *testing.T) {
	values := []bool{false, false, true, false}
	require.True(t, gloop.Any(gloop.Slice(values)))
}

func TestAnyFalse(t *testing.T) {
	values := []bool{false, false, false, false}
	require.False(t, gloop.Any(gloop.Slice(values)))
}

func TestAny2True(t *testing.T) {
	m := map[string]bool{
		"Fizz": true,
		"Buzz": true,
		"Bazz": true,
	}
	require.True(t, gloop.Any2(gloop.Map2(m)))
}

func TestAny2Mixed(t *testing.T) {
	m := map[string]bool{
		"Fizz": false,
		"Buzz": true,
		"Bazz": false,
	}
	require.True(t, gloop.Any2(gloop.Map2(m)))
}

func TestAny2False(t *testing.T) {
	m := map[string]bool{
		"Fizz": false,
		"Buzz": false,
		"Bazz": false,
	}
	require.False(t, gloop.Any2(gloop.Map2(m)))
}
