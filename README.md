[![Genocide Watch](https://hinds-banner.vercel.app/genocide-watch?variant=crimson)](https://www.pcrf.net/)

<p align="center">
    <img alt="gloop logo" src="img/logo.svg" width=500 />
</p>

<p align="center">
    <strong><i>gloop</i></strong> is a Go utility library for convenient looping using the [range-over-func](https://go.dev/blog/range-functions) feature.
</p>

<div align="center">

[![Go Report Card](https://goreportcard.com/badge/github.com/alvii147/gloop)](https://goreportcard.com/report/github.com/alvii147/gloop)

</div>

## Installation

Install `gloop` using the `go get` command:

```bash
go get github.com/alvii147/gloop
```

> [!NOTE]
> Go version 1.23+ required as older versions don't offer the range-over-func feature.

## Usage

Once installed, `gloop` can be imported and used directly in your project:

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	for seq := range gloop.Permutations(gloop.String("CAT"), 3) {
		perm := gloop.ToString(seq)
		fmt.Println(perm)
	}
}
```

This ranges over and outputs all permutations of `CAT`:

```
CAT
CTA
ACT
ATC
TCA
TAC
```

[See more usage examples.](examples_test.go)

## Features

### Generators

* `Interval` - allows looping over values in a given interval of a given step size
* `Linspace` - allows looping over evenly spaced values within a given interval

### Scalar Iterators

* `Chain` - allows looping over multiple [iter.Seq](https://pkg.go.dev/iter#Seq) sequences
* `Chain2` - allows looping over multiple [iter.Seq2](https://pkg.go.dev/iter#Seq2) sequences
* `Channel` - allows looping over values from a given channel
* `Enumerate` - allows looping over an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence with an index, converting it to an [iter.Seq2](https://pkg.go.dev/iter#Seq2) sequence
* `Filter` - runs a given function on each value from an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence and allows looping over values for which the function returns true
* `Filter2` - runs a given function on each value from an [iter.Seq2](https://pkg.go.dev/iter#Seq2) sequence and allows looping over values for which the function returns true
* `Keys` - allows looping over an [iter.Seq2](https://pkg.go.dev/iter#Seq2), converting it to an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence by discarding the value
* `List` - allows looping over a given [linked list](https://pkg.go.dev/container/list#List)
* `Map` - allows looping over keys and values in a map
* `Reverse` - allows looping over an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence in order of descending index
* `Reverse2` - allows looping over an [iter.Seq2](https://pkg.go.dev/iter#Seq2) sequence in order of descending index
* `Slice` - allows looping over a given slice
* `String` - allows looping over the runes in a given string
* `Transform` - runs a given function on each value over an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence and allows looping over the returned values
* `Transform2` - runs a given function on each key and value over an [iter.Seq2](https://pkg.go.dev/iter#Seq2) sequence and allows looping over the returned values
* `Values` - allows looping over an [iter.Seq2](https://pkg.go.dev/iter#Seq2) and converting it to an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence by discarding the key
* `Zip` - allows looping over two [iter.Seq](https://pkg.go.dev/iter#Seq) sequences in pairs

### Vector Iterators

* `Batch` - allows looping over an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence in batches of a given size
* `Batch2` - allows looping over an [iter.Seq2](https://pkg.go.dev/iter#Seq2) sequence in batches of a given size
* `CartesianProduct` - allows looping over the Cartesian product of a given size for an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence
* `CartesianProduct2` - allows looping over the Cartesian product of a given size for an [iter.Seq2](https://pkg.go.dev/iter#Seq2) sequence
* `Combinations` - allows looping over all combinations of a given size for an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence
* `Combinations2` - allows looping over all combinations of a given size for an [iter.Seq2](https://pkg.go.dev/iter#Seq2) sequence
* `Permutations` - allows looping over all permutations of a given size for an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence
* `Permutations2` - allows looping over all permutations of a given size for an [iter.Seq2](https://pkg.go.dev/iter#Seq2) sequence
* `Window` - allows looping over an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence in sliding windows of a given size
* `Window2` - allows looping over an [iter.Seq2](https://pkg.go.dev/iter#Seq2) sequence in sliding windows of a given size

### Accumulators

* `All` - computes whether or not all values in an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence are true
* `Any` - computes whether or not any value in an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence is true
* `Max` - computes the maximum value over an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence
* `Mean` - computes the mean value over an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence
* `Min` - computes the minimum value over an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence
* `Product` - computes the product of values over an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence
* `Reduce` - runs a given function on each value from an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence and accumulates the result into a single value
* `Reduce2` - runs a given function on each value from an [iter.Seq2](https://pkg.go.dev/iter#Seq2) sequence and accumulates the result into a single value
* `Sum` - computes summation over an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence
* `ToList` - converts an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence to a [linked list](https://pkg.go.dev/container/list#List)
* `ToList2` - ToList2 converts an [iter.Seq2](https://pkg.go.dev/iter#Seq2) sequence to [linked lists](https://pkg.go.dev/container/list#List) of keys and values
* `ToSlice` - converts an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence to a slice
* `ToSlice2` - converts an [iter.Seq2](https://pkg.go.dev/iter#Seq2) sequence to slices of keys and values
* `ToString` - converts an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence of runes to a string

### Miscellaneous

* `Parallelize` - runs a function on each value in an [iter.Seq](https://pkg.go.dev/iter#Seq) sequence on separate goroutines
* `Parallelize2` - runs a function on each value in an [iter.Seq2](https://pkg.go.dev/iter#Seq2) sequence on separate goroutines
