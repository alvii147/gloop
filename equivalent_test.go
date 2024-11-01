package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestEquivalentEqual(t *testing.T) {
	values1 := []int{3, 1, 4}
	values2 := []int{3, 1, 4}

	eq := gloop.Equivalent(gloop.Slice(values1), gloop.Slice(values2))
	require.True(t, eq)
}

func TestEquivalentEqualWithRepetition(t *testing.T) {
	values1 := []int{3, 1, 4, 1, 5, 9, 2}
	values2 := []int{3, 1, 4, 1, 5, 9, 2}

	eq := gloop.Equivalent(gloop.Slice(values1), gloop.Slice(values2))
	require.True(t, eq)
}

func TestEquivalentDifferentOrder(t *testing.T) {
	values1 := []int{3, 4, 1}
	values2 := []int{3, 1, 4}

	eq := gloop.Equivalent(gloop.Slice(values1), gloop.Slice(values2))
	require.True(t, eq)

	eq = gloop.Equivalent(gloop.Slice(values2), gloop.Slice(values1))
	require.True(t, eq)
}

func TestEquivalentDifferentOrderWithRepetition(t *testing.T) {
	values1 := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 9}
	values2 := []int{4, 6, 5, 1, 5, 5, 2, 9, 3, 1, 3, 9}

	eq := gloop.Equivalent(gloop.Slice(values1), gloop.Slice(values2))
	require.True(t, eq)

	eq = gloop.Equivalent(gloop.Slice(values2), gloop.Slice(values1))
	require.True(t, eq)
}

func TestEquivalentDifferentLength(t *testing.T) {
	values1 := []int{3, 1, 4, 2}
	values2 := []int{3, 1, 4}

	eq := gloop.Equivalent(gloop.Slice(values1), gloop.Slice(values2))
	require.False(t, eq)

	eq = gloop.Equivalent(gloop.Slice(values2), gloop.Slice(values1))
	require.False(t, eq)
}

func TestEquivalentDifferentValues(t *testing.T) {
	values1 := []int{3, 1, 4}
	values2 := []int{3, 8, 4}

	eq := gloop.Equivalent(gloop.Slice(values1), gloop.Slice(values2))
	require.False(t, eq)

	eq = gloop.Equivalent(gloop.Slice(values2), gloop.Slice(values1))
	require.False(t, eq)
}

func TestEquivalent2Equal(t *testing.T) {
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

	eq := gloop.Equivalent2(seq1, seq2)
	require.True(t, eq)
}

func TestEquivalent2EqualWithRepetition(t *testing.T) {
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

		if !yield("Fuzz", 1) {
			return
		}

		if !yield("Fazz", 5) {
			return
		}

		if !yield("Bizz", 9) {
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

		if !yield("Fuzz", 1) {
			return
		}

		if !yield("Fazz", 5) {
			return
		}

		if !yield("Bizz", 9) {
			return
		}
	}

	eq := gloop.Equivalent2(seq1, seq2)
	require.True(t, eq)
}

func TestEquivalent2DifferentOrder(t *testing.T) {
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

	eq := gloop.Equivalent2(seq1, seq2)
	require.True(t, eq)

	eq = gloop.Equivalent2(seq2, seq1)
	require.True(t, eq)
}

func TestEquivalent2DifferentOrderWithRepetition(t *testing.T) {
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

		if !yield("Fuzz", 1) {
			return
		}

		if !yield("Fazz", 5) {
			return
		}

		if !yield("Bizz", 9) {
			return
		}

		if !yield("Dizz", 2) {
			return
		}

		if !yield("Duzz", 6) {
			return
		}

		if !yield("Fazz", 5) {
			return
		}

		if !yield("Buzz", 3) {
			return
		}

		if !yield("Fazz", 5) {
			return
		}

		if !yield("Bizz", 9) {
			return
		}
	}

	seq2 := func(yield func(string, int) bool) {
		if !yield("Bazz", 4) {
			return
		}

		if !yield("Duzz", 6) {
			return
		}

		if !yield("Fazz", 5) {
			return
		}

		if !yield("Buzz", 1) {
			return
		}

		if !yield("Fazz", 5) {
			return
		}

		if !yield("Fazz", 5) {
			return
		}

		if !yield("Dizz", 2) {
			return
		}

		if !yield("Bizz", 9) {
			return
		}

		if !yield("Fizz", 3) {
			return
		}

		if !yield("Fuzz", 1) {
			return
		}

		if !yield("Buzz", 3) {
			return
		}

		if !yield("Bizz", 9) {
			return
		}
	}

	eq := gloop.Equivalent2(seq1, seq2)
	require.True(t, eq)

	eq = gloop.Equivalent2(seq2, seq1)
	require.True(t, eq)
}

func TestEquivalent2DifferentLength(t *testing.T) {
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

		if !yield("Bizz", 2) {
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

	eq := gloop.Equivalent2(seq1, seq2)
	require.False(t, eq)

	eq = gloop.Equivalent2(seq2, seq1)
	require.False(t, eq)
}

func TestEquivalent2DifferentValues(t *testing.T) {
	seq1 := func(yield func(string, int) bool) {
		if !yield("Fizz", 3) {
			return
		}

		if !yield("Huzz", 1) {
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

	eq := gloop.Equivalent2(seq1, seq2)
	require.False(t, eq)

	eq = gloop.Equivalent2(seq2, seq1)
	require.False(t, eq)
}
