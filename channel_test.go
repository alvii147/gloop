package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestChannel(t *testing.T) {
	values := []string{"Fizz", "Buzz", "Bazz"}
	ch := make(chan string, 3)

	go func() {
		for _, value := range values {
			ch <- value
		}
		close(ch)
	}()

	i := 0
	for value := range gloop.Channel(ch) {
		require.Equal(t, values[i], value)
		i++
	}

	require.Equal(t, len(values), i)
}

func TestChannelBreak(t *testing.T) {
	values := []string{"Fizz", "Buzz", "Bazz"}
	ch := make(chan string, 3)

	go func() {
		for _, value := range values {
			ch <- value
		}
		close(ch)
	}()

	i := 0
	for value := range gloop.Channel(ch) {
		require.Equal(t, "Fizz", value)
		i++
		break
	}

	require.Equal(t, 1, i)
	require.Len(t, ch, 2)
}
