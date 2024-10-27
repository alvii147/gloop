package gloop_test

import (
	"testing"

	"github.com/alvii147/gloop"
	"github.com/stretchr/testify/require"
)

func TestProductInt(t *testing.T) {
	values := []int{3, 4, -5}
	sum := gloop.Product(gloop.Slice(values))
	require.Equal(t, -60, sum)
}

func TestProductFloat(t *testing.T) {
	values := []float64{4.2, 0.5}
	sum := gloop.Product(gloop.Slice(values))
	require.InDelta(t, 2.1, sum, 0.1)
}

func TestProductComplex(t *testing.T) {
	values := []complex128{complex(1, 2), complex(3, -4)}
	sum := gloop.Product(gloop.Slice(values))
	require.Equal(t, complex(11, 2), sum)
}
