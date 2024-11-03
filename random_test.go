package gloop_test

import (
	"math/rand"
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestWithRandomGenerator(t *testing.T) {
	options := gloop.RandomOptions{}
	gloop.WithRandomGenerator(rand.New(rand.NewSource(314)))(&options)

	generator := rand.New(rand.NewSource(314))
	require.Equal(t, generator.Int63(), options.Generator.Int63())
}

func TestRandomUniform(t *testing.T) {
	i := 0
	for value := range gloop.RandomUniform(-2, 3, 10) {
		require.GreaterOrEqual(t, value, -2.0)
		require.Less(t, value, 3.0)
		i++
	}

	require.Equal(t, 10, i)
}

func TestRandomUniformWithRandomGenerator(t *testing.T) {
	i := 0
	generator := rand.New(rand.NewSource(314))
	for value := range gloop.RandomUniform(-2, 3, 10, gloop.WithRandomGenerator(generator)) {
		require.GreaterOrEqual(t, value, -2.0)
		require.Less(t, value, 3.0)
		i++
	}

	require.Equal(t, 10, i)
}

func TestRandomUniformBreak(t *testing.T) {
	i := 0
	for value := range gloop.RandomUniform(-2, 3, 10) {
		if i == 2 {
			break
		}

		require.GreaterOrEqual(t, value, -2.0)
		require.Less(t, value, 3.0)
		i++
	}

	require.Equal(t, 2, i)
}

func TestRandomUniformZeroSizeNoIteration(t *testing.T) {
	for range gloop.RandomUniform(-2, 3, 0) {
		t.Fatal("expected no iteration")
	}
}

func TestRandomUniformNegativeSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.RandomUniform(2, 3, -1) {
			t.Fatal("expected no iteration")
		}
	})
}

func TestRandomNormalWithRandomGenerator(t *testing.T) {
	i := 0
	generator := rand.New(rand.NewSource(314))
	for value := range gloop.RandomNormal(5, 2, 10, gloop.WithRandomGenerator(generator)) {
		_ = value
		i++
	}

	require.Equal(t, 10, i)
}

func TestRandomNormal(t *testing.T) {
	i := 0
	for value := range gloop.RandomNormal(5, 2, 10) {
		_ = value
		i++
	}

	require.Equal(t, 10, i)
}

func TestRandomNormalBreak(t *testing.T) {
	i := 0
	for value := range gloop.RandomNormal(5, 2, 10) {
		_ = value
		if i == 2 {
			break
		}

		i++
	}

	require.Equal(t, 2, i)
}

func TestRandomNormalZeroSizeNoIteration(t *testing.T) {
	for range gloop.RandomNormal(5, 2, 0) {
		t.Fatal("expected no iteration")
	}
}

func TestRandomNormalNegativeSizePanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.RandomNormal(5, 2, -1) {
			t.Fatal("expected no iteration")
		}
	})
}

func TestRandomNormalZeroStandardDeviationPanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.RandomNormal(5, 0, 10) {
			t.Fatal("expected no iteration")
		}
	})
}

func TestRandomNormalNegativeStandardDeviationPanics(t *testing.T) {
	require.Panics(t, func() {
		for range gloop.RandomNormal(5, -2, 10) {
			t.Fatal("expected no iteration")
		}
	})
}
