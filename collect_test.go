package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestCollect(t *testing.T) {
	wantValues := []int{3, 1, 4}
	i := 0
	for value := range gloop.Collect(3, 1, 4) {
		require.Equal(t, wantValues[i], value)
		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestCollectBreak(t *testing.T) {
	wantValues := []int{3, 1}
	i := 0
	for value := range gloop.Collect(3, 1, 4) {
		if i == 2 {
			break
		}

		require.Equal(t, wantValues[i], value)
		i++
	}

	require.Equal(t, len(wantValues), i)
}
