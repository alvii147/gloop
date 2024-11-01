package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestEqualSame(t *testing.T) {
	values1 := []int{3, 1, 4}
	values2 := []int{3, 1, 4}

	eq := gloop.Equal(gloop.Slice(values1), gloop.Slice(values2))
	require.True(t, eq)
}

func TestEqualEmpty(t *testing.T) {
	values1 := []int{}
	values2 := []int{}

	eq := gloop.Equal(gloop.Slice(values1), gloop.Slice(values2))
	require.True(t, eq)
}

func TestEqualDifferentLength(t *testing.T) {
	values1 := []int{3, 1, 4, 2}
	values2 := []int{3, 1, 4}

	eq := gloop.Equal(gloop.Slice(values1), gloop.Slice(values2))
	require.False(t, eq)
}

func TestEqualDifferentOrder(t *testing.T) {
	values1 := []int{1, 3, 4}
	values2 := []int{3, 1, 4}

	eq := gloop.Equal(gloop.Slice(values1), gloop.Slice(values2))
	require.False(t, eq)
}

func TestEqualDifferentValues(t *testing.T) {
	values1 := []int{3, 1, 4}
	values2 := []int{3, 8, 4}

	eq := gloop.Equal(gloop.Slice(values1), gloop.Slice(values2))
	require.False(t, eq)
}

func TestEqual2Same(t *testing.T) {
	seq1 := func(yield func(string, int) bool) {
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

	seq2 := func(yield func(string, int) bool) {
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

	eq := gloop.Equal2(seq1, seq2)
	require.True(t, eq)
}

func TestEqual2Empty(t *testing.T) {
	seq1 := func(yield func(string, int) bool) {}
	seq2 := func(yield func(string, int) bool) {}

	eq := gloop.Equal2(seq1, seq2)
	require.True(t, eq)
}

func TestEqual2DifferentLength(t *testing.T) {
	seq1 := func(yield func(string, int) bool) {
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

	seq2 := func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Buzz", 1) {
			return
		}

		if !yield("Bazz", 4) {
			return
		}

		if !yield("Fuzz", 2) {
			return
		}
	}

	eq := gloop.Equal2(seq1, seq2)
	require.False(t, eq)
}

func TestEqual2DifferentOrder(t *testing.T) {
	seq1 := func(yield func(string, int) bool) {
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

	seq2 := func(yield func(string, int) bool) {
		if !yield("Bazz", 4) {
			return
		}

		if !yield("Fizz", 3) {
			return
		}

		if !yield("Buzz", 1) {
			return
		}
	}

	eq := gloop.Equal2(seq1, seq2)
	require.False(t, eq)
}

func TestEqual2DifferentKeys(t *testing.T) {
	seq1 := func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Cuzz", 1) {
			return
		}

		if !yield("Bazz", 4) {
			return
		}
	}

	seq2 := func(yield func(string, int) bool) {
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

	eq := gloop.Equal2(seq1, seq2)
	require.False(t, eq)
}

func TestEqual2DifferentValues(t *testing.T) {
	seq1 := func(yield func(string, int) bool) {
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

	seq2 := func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Buzz", 1) {
			return
		}

		if !yield("Bazz", 8) {
			return
		}
	}

	eq := gloop.Equal2(seq1, seq2)
	require.False(t, eq)
}
