package gloop_test

import (
	"container/list"
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	l := list.New()
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	wantValues := []int{3, 4, 5}
	i := 0

	for elem := range gloop.List(l) {
		require.Equal(t, wantValues[i], elem.Value)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestListBreak(t *testing.T) {
	l := list.New()
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	wantValues := []int{3, 4}
	i := 0

	for elem := range gloop.List(l) {
		if i == 2 {
			break
		}

		require.Equal(t, wantValues[i], elem.Value)

		i++
	}

	require.Equal(t, len(wantValues), i)
}

func TestToList(t *testing.T) {
	seq := func(yield func(int) bool) {
		if !yield(3) {
			return
		}

		if !yield(4) {
			return
		}

		if !yield(5) {
			return
		}
	}

	values := gloop.ToList(seq)
	elem := values.Front()
	require.Equal(t, 3, elem.Value)
	elem = elem.Next()
	require.Equal(t, 4, elem.Value)
	elem = elem.Next()
	require.Equal(t, 5, elem.Value)
}

func TestToList2(t *testing.T) {
	seq := func(yield func(int, int) bool) {
		if !yield(0, 3) {
			return
		}

		if !yield(1, 4) {
			return
		}

		if !yield(2, 5) {
			return
		}
	}

	keys, values := gloop.ToList2(seq)

	elem := keys.Front()
	require.Equal(t, 0, elem.Value)
	elem = elem.Next()
	require.Equal(t, 1, elem.Value)
	elem = elem.Next()
	require.Equal(t, 2, elem.Value)

	elem = values.Front()
	require.Equal(t, 3, elem.Value)
	elem = elem.Next()
	require.Equal(t, 4, elem.Value)
	elem = elem.Next()
	require.Equal(t, 5, elem.Value)
}
