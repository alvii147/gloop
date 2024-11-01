package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestKeyValue(t *testing.T) {
	pairs := []gloop.KeyValuePair[string, int]{
		{
			Key:   "Fizz",
			Value: 3,
		},
		{
			Key:   "Buzz",
			Value: 1,
		},
		{
			Key:   "Bazz",
			Value: 4,
		},
	}

	i := 0
	for key, value := range gloop.KeyValue(gloop.Slice(pairs)) {
		require.Equal(t, pairs[i].Key, key)
		require.Equal(t, pairs[i].Value, value)
		i++
	}

	require.Equal(t, len(pairs), i)
}

func TestKeyValueBreak(t *testing.T) {
	pairs := []gloop.KeyValuePair[string, int]{
		{
			Key:   "Fizz",
			Value: 3,
		},
		{
			Key:   "Buzz",
			Value: 1,
		},
		{
			Key:   "Bazz",
			Value: 4,
		},
	}

	wantPairs := []gloop.KeyValuePair[string, int]{
		{
			Key:   "Fizz",
			Value: 3,
		},
		{
			Key:   "Buzz",
			Value: 1,
		},
	}

	i := 0
	for key, value := range gloop.KeyValue(gloop.Slice(pairs)) {
		if i == 2 {
			break
		}

		require.Equal(t, pairs[i].Key, key)
		require.Equal(t, pairs[i].Value, value)
		i++
	}

	require.Equal(t, len(wantPairs), i)
}

func TestKeyValue2(t *testing.T) {
	seq := func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Buzz", 1) {
			return
		}

		if !yield("Bazz", 4) {
			return
		}
	}

	wantPairs := []gloop.KeyValuePair[string, int]{
		{
			Key:   "Fizz",
			Value: 3,
		},
		{
			Key:   "Buzz",
			Value: 1,
		},
		{
			Key:   "Bazz",
			Value: 4,
		},
	}

	i := 0
	for pair := range gloop.KeyValue2(seq) {
		require.Equal(t, wantPairs[i], pair)
		i++
	}

	require.Equal(t, len(wantPairs), i)
}
