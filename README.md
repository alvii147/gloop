[![Genocide Watch](https://hinds-banner.vercel.app/genocide-watch?variant=plum)](https://www.pcrf.net/)

<p align="center">
    <img alt="gloop logo" src="img/logo.svg" width=500 />
</p>

<p align="center">
    <strong><i>gloop</i></strong> is a Go utility library for convenient looping using Go's <a href="https://go.dev/blog/range-functions">range-over-func</a> feature.
</p>

<div align="center">

[![Go Reference](https://pkg.go.dev/badge/github.com/alvii147/gloop.svg)](https://pkg.go.dev/github.com/alvii147/gloop) [![Tests](https://img.shields.io/github/actions/workflow/status/alvii147/gloop/github-ci.yml?branch=main&label=tests&logo=github)](https://github.com/alvii147/gloop/actions) ![Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen) [![Go Report Card](https://goreportcard.com/badge/github.com/alvii147/gloop)](https://goreportcard.com/report/github.com/alvii147/gloop) [![License](https://img.shields.io/github/license/alvii147/gloop)](https://github.com/alvii147/gloop/blob/main/LICENSE)

</div>

# Installation

Install `gloop` using the `go get` command:

```bash
go get github.com/alvii147/gloop
```

> [!NOTE]
> Go version 1.23+ required as older versions don't offer the range-over-func feature.

# Usage

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

See more specific documentation and examples in the [features section](#features) below.

# Features

## Generators

* [`Interval`](https://pkg.go.dev/github.com/alvii147/gloop#Interval) allows looping over values in a given interval of a given step size. 
* [`Linspace`](https://pkg.go.dev/github.com/alvii147/gloop#Linspace) allows looping over evenly spaced values within a given interval. n must be greater than 1. 
* [`RandomNormal`](https://pkg.go.dev/github.com/alvii147/gloop#RandomNormal) allows looping over a given number of random values drawn from a Gaussian distribution. The size must not be negative and the standard deviation must be positive. 
* [`RandomUniform`](https://pkg.go.dev/github.com/alvii147/gloop#RandomUniform) allows looping over a given number of random values drawn from a uniform distribution. The size must not be negative. 

## Scalar Iterators

* [`Chain`](https://pkg.go.dev/github.com/alvii147/gloop#Chain) allows looping over multiple [iter.Seq] sequences.
* [`Chain2`](https://pkg.go.dev/github.com/alvii147/gloop#Chain2) allows looping over multiple [iter.Seq2] sequences.
* [`Channel`](https://pkg.go.dev/github.com/alvii147/gloop#Channel) allows looping over values from a given channel. The values are consumed from the channel.
* [`Collect`](https://pkg.go.dev/github.com/alvii147/gloop#Collect) allows looping over a given set of values.
* [`Enumerate`](https://pkg.go.dev/github.com/alvii147/gloop#Enumerate) allows looping over an [iter.Seq] sequence with an index, converting it to an [iter.Seq2] sequence.
* [`Filter`](https://pkg.go.dev/github.com/alvii147/gloop#Filter) runs a given function on each value from an [iter.Seq] sequence and allows looping over values for which the function returns true.
* [`Filter2`](https://pkg.go.dev/github.com/alvii147/gloop#Filter2) runs a given function on each value from an [iter.Seq2] sequence and allows looping over values for which the function returns true.
* [`Keys`](https://pkg.go.dev/github.com/alvii147/gloop#Keys) allows looping over an [iter.Seq2], converting it to an [iter.Seq] sequence by discarding the value.
* [`KeyValue`](https://pkg.go.dev/github.com/alvii147/gloop#KeyValue) converts an [iter.Seq] sequence of [KeyValuePair] values to an [iter.Seq2] sequence.
* [`KeyValue2`](https://pkg.go.dev/github.com/alvii147/gloop#KeyValue2) converts an [iter.Seq2] sequence to an [iter.Seq] sequence of [KeyValuePair] values.
* [`List`](https://pkg.go.dev/github.com/alvii147/gloop#List) allows looping over a given [container/list.List].
* [`Map`](https://pkg.go.dev/github.com/alvii147/gloop#Map) allows looping over keys and values in a map.
* [`Reverse`](https://pkg.go.dev/github.com/alvii147/gloop#Reverse) allows looping over an [iter.Seq] sequence in order of descending index.
* [`Reverse2`](https://pkg.go.dev/github.com/alvii147/gloop#Reverse2) allows looping over an [iter.Seq2] sequence in order of descending index.
* [`Slice`](https://pkg.go.dev/github.com/alvii147/gloop#Slice) allows looping over a given slice.
* [`Sort`](https://pkg.go.dev/github.com/alvii147/gloop#Sort) allows looping over an [iter.Seq] sequence in sorted order.
* [`SortByComparison`](https://pkg.go.dev/github.com/alvii147/gloop#SortByComparison) allows looping over an [iter.Seq] sequence in sorted order using a comparison function.
* [`SortByComparison2`](https://pkg.go.dev/github.com/alvii147/gloop#SortByComparison2) allows looping over an [iter.Seq2] sequence in sorted order using a comparison function.
* [`SortByRank`](https://pkg.go.dev/github.com/alvii147/gloop#SortByRank) allows looping over an [iter.Seq] sequence in sorted order using a ranking function.
* [`SortByRank2`](https://pkg.go.dev/github.com/alvii147/gloop#SortByRank2) allows looping over an [iter.Seq2] sequence in sorted order using a ranking function.
* [`String`](https://pkg.go.dev/github.com/alvii147/gloop#String) allows looping over the runes in a given string.
* [`Transform`](https://pkg.go.dev/github.com/alvii147/gloop#Transform) runs a given function on each value over an [iter.Seq] sequence and allows looping over the returned values.
* [`Transform2`](https://pkg.go.dev/github.com/alvii147/gloop#Transform2) runs a given function on each key and value over an [iter.Seq2] sequence and allows looping over the returned values.
* [`Values`](https://pkg.go.dev/github.com/alvii147/gloop#Values) allows looping over an [iter.Seq2] and converting it to an [iter.Seq] sequence by discarding the key.
* [`Zip`](https://pkg.go.dev/github.com/alvii147/gloop#Zip) allows looping over two [iter.Seq] sequences in pairs.
* [`Zip2`](https://pkg.go.dev/github.com/alvii147/gloop#Zip2) allows looping over two [iter.Seq2] sequences in pairs.

## Vector Iterators

* [`Batch`](https://pkg.go.dev/github.com/alvii147/gloop#Batch) allows looping over an [iter.Seq] sequence in batches of a given size. The batch size must be positive.
* [`Batch2`](https://pkg.go.dev/github.com/alvii147/gloop#Batch2) allows looping over an [iter.Seq2] sequence in batches of a given size. The batch size must be positive.
* [`CartesianProduct`](https://pkg.go.dev/github.com/alvii147/gloop#CartesianProduct) allows looping over the Cartesian product of a given size for an [iter.Seq] sequence. The size must be positive.
* [`CartesianProduct2`](https://pkg.go.dev/github.com/alvii147/gloop#CartesianProduct2) allows looping over the Cartesian product of a given size for an [iter.Seq2] sequence. The size must be positive.
* [`Combinations`](https://pkg.go.dev/github.com/alvii147/gloop#Combinations) allows looping over all combinations of a given size for an [iter.Seq] sequence. The size must be positive.
* [`Combinations2`](https://pkg.go.dev/github.com/alvii147/gloop#Combinations2) allows looping over all combinations of a given size for an [iter.Seq2] sequence. The size must be positive.
* [`Permutations`](https://pkg.go.dev/github.com/alvii147/gloop#Permutations) allows looping over all permutations of a given size for an [iter.Seq] sequence. The size must be positive.
* [`Permutations2`](https://pkg.go.dev/github.com/alvii147/gloop#Permutations2) allows looping over all permutations of a given size for an [iter.Seq2] sequence. The size must be positive.
* [`Window`](https://pkg.go.dev/github.com/alvii147/gloop#Window) allows looping over an [iter.Seq] sequence in sliding windows of a given size.
* [`Window2`](https://pkg.go.dev/github.com/alvii147/gloop#Window2) allows looping over an [iter.Seq2] sequence in sliding windows of a given size.
* [`ZipN`](https://pkg.go.dev/github.com/alvii147/gloop#ZipN) allows looping over multiple [iter.Seq] sequences simultaneously.
* [`ZipN2`](https://pkg.go.dev/github.com/alvii147/gloop#ZipN2) allows looping over multiple [iter.Seq2] sequences simultaneously.

## Aggregators

* [`All`](https://pkg.go.dev/github.com/alvii147/gloop#All) computes whether or not all values in an [iter.Seq] sequence are true.
* [`Any`](https://pkg.go.dev/github.com/alvii147/gloop#Any) computes whether or not any value in an [iter.Seq] sequence is true.
* [`Equal`](https://pkg.go.dev/github.com/alvii147/gloop#Equal) checks if two given [iter.Seq] sequences are exactly equal in contents and order.
* [`Equal2`](https://pkg.go.dev/github.com/alvii147/gloop#Equal2) checks if two given [iter.Seq2] sequences are exactly equal in contents and order.
* [`Equivalent`](https://pkg.go.dev/github.com/alvii147/gloop#Equivalent) checks if two given [iter.Seq] sequences are equal in contents, ignoring order.
* [`Equivalent2`](https://pkg.go.dev/github.com/alvii147/gloop#Equivalent2) checks if two given [iter.Seq2] sequences are equal in contents, ignoring order.
* [`Fold`](https://pkg.go.dev/github.com/alvii147/gloop#Fold) runs a given function on each value from an [iter.Seq] sequence and accumulates the result into a single value.
* [`Fold2`](https://pkg.go.dev/github.com/alvii147/gloop#Fold2) runs a given function on each value from an [iter.Seq2] sequence and accumulates the result into a single value.
* [`Max`](https://pkg.go.dev/github.com/alvii147/gloop#Max) computes the maximum value over an [iter.Seq] sequence.
* [`MaxByComparison`](https://pkg.go.dev/github.com/alvii147/gloop#MaxByComparison) computes the maximum value over an [iter.Seq] sequence using a comparison function.
* [`MaxByComparison2`](https://pkg.go.dev/github.com/alvii147/gloop#MaxByComparison2) computes the maximum key and value over an [iter.Seq2] sequence using a comparison function.
* [`MaxByRank`](https://pkg.go.dev/github.com/alvii147/gloop#MaxByRank) computes the maximum value over an [iter.Seq] sequence using a ranking function.
* [`MaxByRank2`](https://pkg.go.dev/github.com/alvii147/gloop#MaxByRank2) computes the maximum value over an [iter.Seq2] sequence using a ranking function.
* [`Mean`](https://pkg.go.dev/github.com/alvii147/gloop#Mean) computes the mean value over an [iter.Seq] sequence.
* [`Min`](https://pkg.go.dev/github.com/alvii147/gloop#Min) computes the minimum value over an [iter.Seq] sequence.
* [`MinByComparison`](https://pkg.go.dev/github.com/alvii147/gloop#MinByComparison) computes the minimum value over an [iter.Seq] sequence using a comparison function.
* [`MinByComparison2`](https://pkg.go.dev/github.com/alvii147/gloop#MinByComparison2) computes the minimum key and value over an [iter.Seq2] sequence using a comparison function.
* [`MinByRank`](https://pkg.go.dev/github.com/alvii147/gloop#MinByRank) computes the minimum value over an [iter.Seq] sequence using a ranking function.
* [`MinByRank2`](https://pkg.go.dev/github.com/alvii147/gloop#MinByRank2) computes the minimum value over an [iter.Seq2] sequence using a ranking function.
* [`Product`](https://pkg.go.dev/github.com/alvii147/gloop#Product) computes the product of values over an [iter.Seq] sequence.
* [`Reduce`](https://pkg.go.dev/github.com/alvii147/gloop#Reduce) runs a given function on each adjacent pair in an [iter.Seq] sequence and accumulates the result into a single value.
* [`Reduce2`](https://pkg.go.dev/github.com/alvii147/gloop#Reduce2) runs a given function on each adjacent pair of keys and values in an [iter.Seq2] sequence and accumulates the result into a single key and value pair.
* [`Sum`](https://pkg.go.dev/github.com/alvii147/gloop#Sum) computes summation over an [iter.Seq] sequence.
* [`ToList`](https://pkg.go.dev/github.com/alvii147/gloop#ToList) converts an [iter.Seq] sequence to a [container/list.List].
* [`ToList2`](https://pkg.go.dev/github.com/alvii147/gloop#ToList2) converts an [iter.Seq2] sequence to [container/list.List] of keys and values.
* [`ToSlice`](https://pkg.go.dev/github.com/alvii147/gloop#ToSlice) converts an [iter.Seq] sequence to a slice.
* [`ToSlice2`](https://pkg.go.dev/github.com/alvii147/gloop#ToSlice2) converts an [iter.Seq2] sequence to slices of keys and values.
* [`ToString`](https://pkg.go.dev/github.com/alvii147/gloop#ToString) converts an [iter.Seq] sequence of runes to a string.

### Miscellaneous

* [`DeferLoop`](https://pkg.go.dev/github.com/alvii147/gloop#DeferLoop) allows looping over an [iter.Seq] sequence, yielding a defer function that can register another function to be executed at the end of the currently running loop. If multiple functions are registered, they are executed in FIFO order.
* [`Parallelize`](https://pkg.go.dev/github.com/alvii147/gloop#Parallelize) runs a function on each value in an [iter.Seq] sequence on separate goroutines.
* [`Parallelize2`](https://pkg.go.dev/github.com/alvii147/gloop#Parallelize2) runs a function on each value in an [iter.Seq2] sequence on separate goroutines.

[iter.Seq]: https://pkg.go.dev/iter#Seq
[iter.Seq2]: https://pkg.go.dev/iter#Seq2
[container/list.List]: https://pkg.go.dev/container/list#List
