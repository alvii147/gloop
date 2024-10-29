package gloop

import (
	"iter"
	"math/rand"
	"time"
)

// RandomOptions defines configurable options for [RandomUniform] and
// [RandomNormal].
type RandomOptions struct {
	// Generator is the random number generator.
	Generator *rand.Rand
}

// RandomOptionsFunc is the function signature of configuration helpers
// for [RandomUniform] and [RandomNormal].
type RandomOptionsFunc func(*RandomOptions)

// WithRandomGenerator is a helper for configuring the random number
// generator for [RandomUniform] and [RandomNormal].
func WithRandomGenerator(generator *rand.Rand) RandomOptionsFunc {
	return func(o *RandomOptions) {
		o.Generator = generator
	}
}

// RandomUniform allows looping over a given number of random values
// drawn from a uniform distribution. The size must not be negative.
func RandomUniform[N Number](
	low N,
	high N,
	size int,
	opts ...RandomOptionsFunc,
) iter.Seq[float64] {
	if size < 0 {
		panic("size must not be negative")
	}

	options := RandomOptions{
		Generator: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	for _, opt := range opts {
		opt(&options)
	}

	flow := float64(low)
	fhigh := float64(high)
	delta := fhigh - flow

	return func(yield func(float64) bool) {
		for range size {
			if !yield(((options.Generator.Float64() * delta) + flow)) {
				return
			}
		}
	}
}

// RandomNormal allows looping over a given number of random values
// drawn from a Gaussian distribution. The size must not be negative
// and the standard deviation must be positive.
func RandomNormal[N Number](
	mean N,
	stddev N,
	size int,
	opts ...RandomOptionsFunc,
) iter.Seq[float64] {
	if size < 0 {
		panic("size must not be negative")
	}

	if stddev <= 0 {
		panic("standard deviation must be positive")
	}

	options := RandomOptions{
		Generator: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	for _, opt := range opts {
		opt(&options)
	}

	fmean := float64(mean)
	fstddev := float64(stddev)

	return func(yield func(float64) bool) {
		for range size {
			if !yield((options.Generator.NormFloat64() * fstddev) + fmean) {
				return
			}
		}
	}
}
