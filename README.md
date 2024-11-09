[![Genocide Watch](https://hinds-banner.vercel.app/genocide-watch?variant=plum)](https://www.pcrf.net/)

<p align="center">
    <img alt="gloop logo" src="img/logo.svg" width=500 />
</p>

<p align="center">
    <strong><i>gloop</i></strong> is a Go utility library for convenient looping using Go's <a href="https://go.dev/blog/range-functions">range-over-func</a> feature.
</p>

<div align="center">

[![Go Reference](https://pkg.go.dev/badge/github.com/alvii147/gloop.svg)](https://pkg.go.dev/github.com/alvii147/gloop) [![GitHub CI](https://img.shields.io/github/actions/workflow/status/alvii147/gloop/github-ci.yml?branch=main&label=GitHub%20CI&logo=github)](https://github.com/alvii147/gloop/actions) [![Go Report Card](https://goreportcard.com/badge/github.com/alvii147/gloop)](https://goreportcard.com/report/github.com/alvii147/gloop) ![License](https://img.shields.io/github/license/alvii147/gloop)

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

See more examples below.

# Features

## Generators

<details>
<summary><h3>Interval</h3></summary>

`Interval` allows looping over values in a given interval of a given step size. 

```go
{
	for i := range gloop.Interval(3, 9, 2) {
		fmt.Println(i)
	}

}
```

Output:

```
3
5
7
```

</details>

<details>
<summary><h3>Linspace</h3></summary>

`Linspace` allows looping over evenly spaced values within a given interval. n must be greater than 1. 

```go
{
	for i := range gloop.Linspace(2, 3, 5) {
		fmt.Println(i)
	}

}
```

Output:

```
2
2.25
2.5
2.75
```

</details>

<details>
<summary><h3>RandomNormal</h3></summary>

`RandomNormal` allows looping over a given number of random values drawn from a Gaussian distribution. The size must not be negative and the standard deviation must be positive. 

```go
{
	for i := range gloop.RandomNormal(2, 2, 5) {
		fmt.Println(i)
	}

}
```

Output:

```
3.6053409359773543
4.8493077906535165
1.321369004660313
1.3549030774712296
-0.6521572615302738
```

</details>

<details>
<summary><h3>RandomUniform</h3></summary>

`RandomUniform` allows looping over a given number of random values drawn from a uniform distribution. The size must not be negative. 

```go
{
	for i := range gloop.RandomUniform(0, 2, 5) {
		fmt.Println(i)
	}

}
```

Output:

```
1.7336396942444041
0.9684446802268123
1.5762348358917075
0.5000463191262544
1.1113562403363295
```

</details>

## Scalar Iterators

<details>
<summary><h3>Chain</h3></summary>

`Chain` allows looping over multiple [iter.Seq] sequences. 

```go
{
	values1 := []int{3, 1, 4}
	values2 := []int{1, 6}
	for i := range gloop.Chain(gloop.Slice(values1), gloop.Slice(values2)) {
		fmt.Println(i)
	}

}
```

Output:

```
3
1
4
1
6
```

</details>

<details>
<summary><h3>Chain2</h3></summary>

`Chain2` allows looping over multiple [iter.Seq2] sequences. 

```go
{
	m1 := map[string]int{
		"CAT":	3,
		"DOG":	1,
	}

	m2 := map[string]int{
		"MOUSE": 4,
	}

	for key, value := range gloop.Chain2(gloop.Map(m1), gloop.Map(m2)) {
		fmt.Println(key, value)
	}

}
```

Output:

```
CAT 3
DOG 1
MOUSE 4
```

</details>

<details>
<summary><h3>Channel</h3></summary>

`Channel` allows looping over values from a given channel. The values are consumed from the channel. 

```go
{
	ch := make(chan string)
	go func() {
		ch <- "CAT"
		ch <- "DOG"
		ch <- "MOUSE"
		close(ch)
	}()

	for i := range gloop.Channel(ch) {
		fmt.Println(i)
	}

}
```

Output:

```
CAT
DOG
MOUSE
```

</details>

<details>
<summary><h3>Collect</h3></summary>

`Collect` allows looping over a given set of values. 

```go
{
	for i := range gloop.Collect(3, 1, 4) {
		fmt.Println(i)
	}

}
```

Output:

```
3
1
4
```

</details>

<details>
<summary><h3>Enumerate</h3></summary>

`Enumerate` allows looping over an [iter.Seq] sequence with an index, converting it to an [iter.Seq2] sequence. 

```go
{
	ch := make(chan int)
	go func() {
		ch <- 3
		ch <- 1
		ch <- 4
		close(ch)
	}()

	for i, value := range gloop.Enumerate(gloop.Channel(ch)) {
		fmt.Println(i, value)
	}

}
```

Output:

```
0 3
1 1
2 4
```

</details>

<details>
<summary><h3>Filter</h3></summary>

`Filter` runs a given function on each value from an [iter.Seq] sequence and allows looping over values for which the function returns true. 

```go
{
	isOdd := func(i int) bool {
		return i%2 == 1
	}

	values := []int{3, 1, 4}
	for i := range gloop.Filter(gloop.Slice(values), isOdd) {
		fmt.Println(i)
	}

}
```

Output:

```
3
1
```

</details>

<details>
<summary><h3>Filter2</h3></summary>

`Filter2` runs a given function on each value from an [iter.Seq2] sequence and allows looping over values for which the function returns true. 

```go
{
	isProductPositive := func(i, j int) bool {
		return i*j >= 0
	}

	m := map[int]int{
		-3:	3,
		-1:	-1,
		4:	4,
	}

	for i, j := range gloop.Filter2(gloop.Map(m), isProductPositive) {
		fmt.Println(i, j)
	}

}
```

Output:

```
-1 -1
4 4
```

</details>

<details>
<summary><h3>Keys</h3></summary>

`Keys` allows looping over an [iter.Seq2], converting it to an [iter.Seq] sequence by discarding the value. 

```go
{
	m := map[string]int{
		"CAT":		3,
		"DOG":		1,
		"MOUSE":	4,
	}

	for key := range gloop.Keys(gloop.Map(m)) {
		fmt.Println(key)
	}

}
```

Output:

```
CAT
DOG
MOUSE
```

</details>

<details>
<summary><h3>KeyValue</h3></summary>

`KeyValue` converts an [iter.Seq] sequence of [KeyValuePair] values to an [iter.Seq2] sequence. 

```go
{
	pairs := []gloop.KeyValuePair[string, int]{
		{
			Key:	"CAT",
			Value:	3,
		},
		{
			Key:	"DOG",
			Value:	1,
		},
		{
			Key:	"MOUSE",
			Value:	4,
		},
	}

	for key, value := range gloop.KeyValue(gloop.Slice(pairs)) {
		fmt.Println(key, value)
	}

}
```

Output:

```
CAT 3
DOG 1
MOUSE 4
```

</details>

<details>
<summary><h3>KeyValue2</h3></summary>

`KeyValue2` converts an [iter.Seq2] sequence to an [iter.Seq] sequence of [KeyValuePair] values. 

```go
{
	m := map[string]int{
		"CAT":		3,
		"DOG":		1,
		"MOUSE":	4,
	}

	for pair := range gloop.KeyValue2(gloop.Map(m)) {
		fmt.Println(pair.Key, pair.Value)
	}

}
```

Output:

```
CAT 3
DOG 1
MOUSE 4
```

</details>

<details>
<summary><h3>List</h3></summary>

`List` allows looping over a given [container/list.List]. 

```go
{
	l := list.New()
	l.PushBack(3)
	l.PushBack(1)
	l.PushBack(4)

	for elem := range gloop.List(l) {
		fmt.Println(elem.Value)
	}

}
```

Output:

```
3
1
4
```

</details>

<details>
<summary><h3>Map</h3></summary>

`Map` allows looping over keys and values in a map. 

```go
{
	m := map[string]int{
		"CAT":		3,
		"DOG":		1,
		"MOUSE":	4,
	}

	for key, value := range gloop.Map(m) {
		fmt.Println(key, value)
	}

}
```

Output:

```
CAT 3
DOG 1
MOUSE 4
```

</details>

<details>
<summary><h3>Reverse</h3></summary>

`Reverse` allows looping over an [iter.Seq] sequence in order of descending index. 

```go
{
	values := []int{3, 1, 4}
	for i := range gloop.Reverse(gloop.Slice(values)) {
		fmt.Println(i)
	}

}
```

Output:

```
4
1
3
```

</details>

<details>
<summary><h3>Reverse2</h3></summary>

`Reverse2` allows looping over an [iter.Seq2] sequence in order of descending index. 

```go
{
	values := []int{3, 1, 4}
	for i, value := range gloop.Reverse2(gloop.Enumerate(gloop.Slice(values))) {
		fmt.Println(i, value)
	}

}
```

Output:

```
2 4
1 1
0 3
```

</details>

<details>
<summary><h3>Slice</h3></summary>

`Slice` allows looping over a given slice. 

```go
{
	values := []int{3, 1, 4}
	for i := range gloop.Slice(values) {
		fmt.Println(i)
	}

}
```

Output:

```
3
1
4
```

</details>

<details>
<summary><h3>Sort</h3></summary>

`Sort` allows looping over an [iter.Seq] sequence in sorted order. 

```go
{
	values := []int{3, 1, 4, 1, 5, 9}
	for i := range gloop.Sort(gloop.Slice(values), true) {
		fmt.Println(i)
	}

}
```

Output:

```
1
1
3
4
5
9
```

</details>

<details>
<summary><h3>SortByComparison</h3></summary>

`SortByComparison` allows looping over an [iter.Seq] sequence in sorted order using a comparison function. 

```go
{
	compareStringLens := func(s1, s2 string) bool {
		return len(s1) < len(s2)
	}

	values := []string{"CAT", "MOUSE", "DOG"}
	for s := range gloop.SortByComparison(gloop.Slice(values), compareStringLens, true) {
		fmt.Println(s)
	}

}
```

Output:

```
CAT
DOG
MOUSE
```

</details>

<details>
<summary><h3>SortByComparison2</h3></summary>

`SortByComparison2` allows looping over an [iter.Seq2] sequence in sorted order using a comparison function. 

```go
{
	compareKeyValueConcatLen := func(k1, v1, k2, v2 string) bool {
		return len(k1+v1) < len(k2+v2)
	}

	values := map[string]string{
		"CAT":		"DOG",
		"MOUSE":	"CHICKEN",
		"BUNNY":	"BEAR",
	}
	for key, value := range gloop.SortByComparison2(gloop.Map(values), compareKeyValueConcatLen, true) {
		fmt.Println(key, value)
	}

}
```

Output:

```
CAT DOG
BUNNY BEAR
MOUSE CHICKEN
```

</details>

<details>
<summary><h3>SortByRank</h3></summary>

`SortByRank` allows looping over an [iter.Seq] sequence in sorted order using a ranking function. 

```go
{
	stringLen := func(s string) int {
		return len(s)
	}

	values := []string{"CAT", "MOUSE", "DOG"}
	for s := range gloop.SortByRank(gloop.Slice(values), stringLen, true) {
		fmt.Println(s)
	}

}
```

Output:

```
CAT
DOG
MOUSE
```

</details>

<details>
<summary><h3>SortByRank2</h3></summary>

`SortByRank2` allows looping over an [iter.Seq2] sequence in sorted order using a ranking function. 

```go
{
	stringConcatLen := func(k1, v1 string) int {
		return len(k1 + v1)
	}

	values := map[string]string{
		"CAT":		"DOG",
		"MOUSE":	"CHICKEN",
		"BUNNY":	"BEAR",
	}
	for key, value := range gloop.SortByRank2(gloop.Map(values), stringConcatLen, true) {
		fmt.Println(key, value)
	}

}
```

Output:

```
CAT DOG
BUNNY BEAR
MOUSE CHICKEN
```

</details>

<details>
<summary><h3>String</h3></summary>

`String` allows looping over the runes in a given string. 

```go
{
	for r := range gloop.String("CAT") {
		fmt.Println(string(r))
	}

}
```

Output:

```
C
A
T
```

</details>

<details>
<summary><h3>Transform</h3></summary>

`Transform` runs a given function on each value over an [iter.Seq] sequence and allows looping over the returned values. 

```go
{
	values := []string{"CaT", "dOg"}
	for s := range gloop.Transform(gloop.Slice(values), strings.ToUpper) {
		fmt.Println(s)
	}

}
```

Output:

```
CAT
DOG
```

</details>

<details>
<summary><h3>Transform2</h3></summary>

`Transform2` runs a given function on each key and value over an [iter.Seq2] sequence and allows looping over the returned values. 

```go
{
	concat := func(s1, s2 string) string {
		return s1 + s2
	}

	m := map[string]string{
		"CAT":		"DOG",
		"MOUSE":	"CHICKEN",
	}

	for s := range gloop.Transform2(gloop.Map(m), concat) {
		fmt.Println(s)
	}

}
```

Output:

```
CATDOG
MOUSECHICKEN
```

</details>

<details>
<summary><h3>Values</h3></summary>

`Values` allows looping over an [iter.Seq2] and converting it to an [iter.Seq] sequence by discarding the key. 

```go
{
	m := map[string]int{
		"CAT":		3,
		"DOG":		1,
		"MOUSE":	4,
	}

	for value := range gloop.Values(gloop.Map(m)) {
		fmt.Println(value)
	}

}
```

Output:

```
3
1
4
```

</details>

<details>
<summary><h3>Zip</h3></summary>

`Zip` allows looping over two [iter.Seq] sequences in pairs. 

```go
{
	values1 := []string{"CAT", "DOG", "MOUSE"}
	values2 := []int{3, 1, 4}
	for value1, value2 := range gloop.Zip(gloop.Slice(values1), gloop.Slice(values2)) {
		fmt.Println(value1, value2)
	}

}
```

Output:

```
CAT 3
DOG 1
MOUSE 4
```

</details>

<details>
<summary><h3>Zip2</h3></summary>

`Zip2` allows looping over two [iter.Seq2] sequences in pairs. 

```go
{
	seq1 := func(yield func(string, int) bool) {
		yield("CAT", 3)
		yield("DOG", 1)
		yield("MOUSE", 4)
	}

	seq2 := func(yield func(int, float64) bool) {
		yield(3, 1.2)
		yield(1, 3.4)
		yield(4, 5.6)
	}

	for pair1, pair2 := range gloop.Zip2(seq1, seq2) {
		fmt.Println(pair1.Key, pair1.Value, pair2.Key, pair2.Value)
	}

}
```

Output:

```
CAT 3 3 1.2
DOG 1 1 3.4
MOUSE 4 4 5.6
```

</details>

## Vector Iterators

<details>
<summary><h3>Batch</h3></summary>

`Batch` allows looping over an [iter.Seq] sequence in batches of a given size. The batch size must be positive. 

```go
{
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	for seq := range gloop.Batch(gloop.Slice(values), 3) {
		batch := gloop.ToSlice(seq)
		fmt.Println(batch)
	}

}
```

Output:

```
[3 1 4]
[1 5 9]
[2 6 5]
```

</details>

<details>
<summary><h3>Batch2</h3></summary>

`Batch2` allows looping over an [iter.Seq2] sequence in batches of a given size. The batch size must be positive. 

```go
{
	values := []string{"CAT", "DOG", "MOUSE", "CHICKEN", "BUNNY", "BEAR"}
	for seq := range gloop.Batch2(gloop.Enumerate(gloop.Slice(values)), 3) {
		batchKeys, batchValues := gloop.ToSlice2(seq)
		fmt.Println(batchKeys, batchValues)
	}

}
```

Output:

```
[0 1 2] [CAT DOG MOUSE]
[3 4 5] [CHICKEN BUNNY BEAR]
```

</details>

<details>
<summary><h3>CartesianProduct</h3></summary>

`CartesianProduct` allows looping over the Cartesian product of a given size for an [iter.Seq] sequence. The size must be positive. 

```go
{
	s := "CAT"
	for seq := range gloop.CartesianProduct(gloop.String(s), 2) {
		product := gloop.ToString(seq)
		fmt.Println(product)
	}

}
```

Output:

```
CC
CA
CT
AC
AA
AT
TC
TA
TT
```

</details>

<details>
<summary><h3>CartesianProduct2</h3></summary>

`CartesianProduct2` allows looping over the Cartesian product of a given size for an [iter.Seq2] sequence. The size must be positive. 

```go
{
	m := map[string]int{
		"CAT":	3,
		"DOG":	1,
	}
	for seq := range gloop.CartesianProduct2(gloop.Map(m), 2) {
		productKeys, productValues := gloop.ToSlice2(seq)
		fmt.Println(productKeys, productValues)
	}

}
```

Output:

```
[CAT CAT] [3 3]
[CAT DOG] [3 1]
[DOG CAT] [1 3]
[DOG DOG] [1 1]
```

</details>

<details>
<summary><h3>Combinations</h3></summary>

`Combinations` allows looping over all combinations of a given size for an [iter.Seq] sequence. The size must be positive. 

```go
{
	s := "CAT"
	for seq := range gloop.Combinations(gloop.String(s), 2) {
		comb := gloop.ToString(seq)
		fmt.Println(comb)
	}

}
```

Output:

```
CA
CT
AT
```

</details>

<details>
<summary><h3>Combinations2</h3></summary>

`Combinations2` allows looping over all combinations of a given size for an [iter.Seq2] sequence. The size must be positive. 

```go
{
	m := map[string]int{
		"CAT":		3,
		"DOG":		1,
		"MOUSE":	4,
	}
	for seq := range gloop.Combinations2(gloop.Map(m), 2) {
		combKeys, combValues := gloop.ToSlice2(seq)
		fmt.Println(combKeys, combValues)
	}

}
```

Output:

```
[CAT DOG] [3 1]
[CAT MOUSE] [3 4]
[DOG MOUSE] [1 4]
```

</details>

<details>
<summary><h3>Permutations</h3></summary>

`Permutations` allows looping over all permutations of a given size for an [iter.Seq] sequence. The size must be positive. 

```go
{
	s := "CAT"
	for seq := range gloop.Permutations(gloop.String(s), 2) {
		perm := gloop.ToString(seq)
		fmt.Println(perm)
	}

}
```

Output:

```
CA
CT
AC
AT
TC
TA
```

</details>

<details>
<summary><h3>Permutations2</h3></summary>

`Permutations2` allows looping over all permutations of a given size for an [iter.Seq2] sequence. The size must be positive. 

```go
{
	m := map[string]int{
		"CAT":		3,
		"DOG":		1,
		"MOUSE":	4,
	}
	for seq := range gloop.Permutations2(gloop.Map(m), 2) {
		permKeys, permValues := gloop.ToSlice2(seq)
		fmt.Println(permKeys, permValues)
	}

}
```

Output:

```
[CAT DOG] [3 1]
[CAT MOUSE] [3 4]
[DOG CAT] [1 3]
[DOG MOUSE] [1 4]
[MOUSE CAT] [4 3]
[MOUSE DOG] [4 1]
```

</details>

<details>
<summary><h3>Window</h3></summary>

`Window` allows looping over an [iter.Seq] sequence in sliding windows of a given size. 

```go
{
	values := []int{3, 1, 4, 1, 5, 9}
	for seq := range gloop.Window(gloop.Slice(values), 3) {
		window := gloop.ToSlice(seq)
		fmt.Println(window)
	}

}
```

Output:

```
[3 1 4]
[1 4 1]
[4 1 5]
[1 5 9]
```

</details>

<details>
<summary><h3>Window2</h3></summary>

`Window2` allows looping over an [iter.Seq2] sequence in sliding windows of a given size. 

```go
{
	values := []string{"CAT", "DOG", "MOUSE", "CHICKEN", "BUNNY", "BEAR"}
	for seq := range gloop.Window2(gloop.Enumerate(gloop.Slice(values)), 3) {
		windowKeys, windowValues := gloop.ToSlice2(seq)
		fmt.Println(windowKeys, windowValues)
	}

}
```

Output:

```
[0 1 2] [CAT DOG MOUSE]
[1 2 3] [DOG MOUSE CHICKEN]
[2 3 4] [MOUSE CHICKEN BUNNY]
[3 4 5] [CHICKEN BUNNY BEAR]
```

</details>

<details>
<summary><h3>ZipN</h3></summary>

`ZipN` allows looping over multiple [iter.Seq] sequences simultaneously. 

```go
{
	seq1 := gloop.Slice([]string{"CAT", "DOG"})
	seq2 := gloop.Slice([]string{"MOUSE", "CHICKEN"})
	seq3 := gloop.Slice([]string{"BUNNY", "BEAR"})

	for seq := range gloop.ZipN(gloop.Collect(seq1, seq2, seq3)) {
		fmt.Println(gloop.ToSlice(seq))
	}

}
```

Output:

```
[CAT MOUSE BUNNY]
[DOG CHICKEN BEAR]
```

</details>

<details>
<summary><h3>ZipN2</h3></summary>

`ZipN2` allows looping over multiple [iter.Seq2] sequences simultaneously. 

```go
{
	var seq1 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		yield("CAT", 3)
		yield("DOG", 1)
		yield("MOUSE", 4)
	}

	var seq2 iter.Seq2[string, int] = func(yield func(string, int) bool) {
		yield("MOUSE", 1)
		yield("BUNNY", 5)
		yield("BEAR", 9)
	}

	for seq := range gloop.ZipN2(gloop.Collect(seq1, seq2)) {
		keys, values := gloop.ToSlice2(seq)
		fmt.Println(keys, values)
	}

}
```

Output:

```
[CAT MOUSE] [3 1]
[DOG BUNNY] [1 5]
[MOUSE BEAR] [4 9]
```

</details>

## Aggregators

<details>
<summary><h3>All</h3></summary>

`All` computes whether or not all values in an [iter.Seq] sequence are true. 

```go
{
	allTrue := []bool{true, true, true}
	someTrue := []bool{true, false, true}
	allFalse := []bool{false, false, false}

	fmt.Println(gloop.All(gloop.Slice(allTrue)))
	fmt.Println(gloop.All(gloop.Slice(someTrue)))
	fmt.Println(gloop.All(gloop.Slice(allFalse)))

}
```

Output:

```
true
false
false
```

</details>

<details>
<summary><h3>Any</h3></summary>

`Any` computes whether or not any value in an [iter.Seq] sequence is true. 

```go
{
	allTrue := []bool{true, true, true}
	someTrue := []bool{true, false, true}
	allFalse := []bool{false, false, false}

	fmt.Println(gloop.Any(gloop.Slice(allTrue)))
	fmt.Println(gloop.Any(gloop.Slice(someTrue)))
	fmt.Println(gloop.Any(gloop.Slice(allFalse)))

}
```

Output:

```
true
true
false
```

</details>

<details>
<summary><h3>Equal</h3></summary>

`Equal` checks if two given [iter.Seq] sequences are exactly equal in contents and order. 

```go
{
	values1 := []int{3, 1, 4}
	values2 := []int{3, 1, -4}
	values3 := []int{3, 1, 4}

	fmt.Println(gloop.Equal(gloop.Slice(values1), gloop.Slice(values2)))
	fmt.Println(gloop.Equal(gloop.Slice(values1), gloop.Slice(values3)))

}
```

Output:

```
false
true
```

</details>

<details>
<summary><h3>Equal2</h3></summary>

`Equal2` checks if two given [iter.Seq2] sequences are exactly equal in contents and order. 

```go
{
	seq1 := func(yield func(string, int) bool) {
		yield("CAT", 3)
		yield("DOG", 1)
		yield("MOUSE", 4)
	}

	seq2 := func(yield func(string, int) bool) {
		yield("CAT", 3)
		yield("DOG", 1)
		yield("CHICKEN", 4)
	}

	seq3 := func(yield func(string, int) bool) {
		yield("CAT", 3)
		yield("DOG", 1)
		yield("MOUSE", 4)
	}

	fmt.Println(gloop.Equal2(seq1, seq2))
	fmt.Println(gloop.Equal2(seq1, seq3))

}
```

Output:

```
false
true
```

</details>

<details>
<summary><h3>Equivalent</h3></summary>

`Equivalent` checks if two given [iter.Seq] sequences are equal in contents, ignoring order. 

```go
{
	values1 := []int{3, 1, 4}
	values2 := []int{3, 1, -4}
	values3 := []int{3, 4, 1}

	fmt.Println(gloop.Equivalent(gloop.Slice(values1), gloop.Slice(values2)))
	fmt.Println(gloop.Equivalent(gloop.Slice(values1), gloop.Slice(values3)))

}
```

Output:

```
false
true
```

</details>

<details>
<summary><h3>Equivalent2</h3></summary>

`Equivalent2` checks if two given [iter.Seq2] sequences are equal in contents, ignoring order. 

```go
{
	seq1 := func(yield func(string, int) bool) {
		yield("CAT", 3)
		yield("DOG", 1)
		yield("MOUSE", 4)
	}

	seq2 := func(yield func(string, int) bool) {
		yield("CAT", 3)
		yield("DOG", 1)
		yield("CHICKEN", 4)
	}

	seq3 := func(yield func(string, int) bool) {
		yield("CAT", 3)
		yield("MOUSE", 4)
		yield("DOG", 1)
	}

	fmt.Println(gloop.Equivalent2(seq1, seq2))
	fmt.Println(gloop.Equivalent2(seq1, seq3))

}
```

Output:

```
false
true
```

</details>

<details>
<summary><h3>Fold</h3></summary>

`Fold` runs a given function on each value from an [iter.Seq] sequence and accumulates the result into a single value. 

```go
{
	add := func(a, b int) int {
		return a + b
	}

	values := []int{3, 1, 4}
	sum := gloop.Fold(gloop.Slice(values), add)
	fmt.Println(sum)

}
```

Output:

```
8
```

</details>

<details>
<summary><h3>Fold2</h3></summary>

`Fold2` runs a given function on each value from an [iter.Seq2] sequence and accumulates the result into a single value. 

```go
{
	addKeyValueProduct := func(acc, key, value int) int {
		return acc + (key * value)
	}

	m := map[int]int{
		3:	1,
		1:	5,
		4:	9,
	}

	sumOfProducts := gloop.Fold2(gloop.Map(m), addKeyValueProduct)
	fmt.Println(sumOfProducts)

}
```

Output:

```
44
```

</details>

<details>
<summary><h3>Max</h3></summary>

`Max` computes the maximum value over an [iter.Seq] sequence. 

```go
{
	values := []int{3, 1, 4, 2}
	maxValue := gloop.Max(gloop.Slice(values))
	fmt.Println(maxValue)

}
```

Output:

```
4
```

</details>

<details>
<summary><h3>MaxByComparison</h3></summary>

`MaxByComparison` computes the maximum value over an [iter.Seq] sequence using a comparison function.

```go
{
	compareStringLens := func(s1, s2 string) bool {
		return len(s1) < len(s2)
	}

	values := []string{"CAT", "MOUSE", "CHICKEN"}
	maxValue := gloop.MaxByComparison(gloop.Slice(values), compareStringLens)
	fmt.Println(maxValue)

}
```

Output:

```
CHICKEN
```

</details>

<details>
<summary><h3>MaxByComparison2</h3></summary>

`MaxByComparison2` computes the maximum key and value over an [iter.Seq2] sequence using a comparison function.

```go
{
	compareKeyValueProducts := func(k1, v1, k2, v2 int) bool {
		return k1*v1 < k2*v2
	}

	m := map[int]int{
		3: 1,
		1: 5,
		4: 9,
	}

	maxKey, maxValue := gloop.MaxByComparison2(gloop.Map(m), compareKeyValueProducts)
	fmt.Println(maxKey, maxValue)
	// Output:
	// 4 9

}
```

Output:

```
4 9
```

</details>

<details>
<summary><h3>Mean</h3></summary>

`Mean` computes the mean value over an [iter.Seq] sequence. 

```go
{
	values := []int{3, 1, 4, 2}
	mean := gloop.Mean(gloop.Slice(values))
	fmt.Println(mean)

}
```

Output:

```
2.5
```

</details>

<details>
<summary><h3>Min</h3></summary>

`Min` computes the minimum value over an [iter.Seq] sequence. 

```go
{
	values := []int{3, 1, 4, 2}
	minValue := gloop.Min(gloop.Slice(values))
	fmt.Println(minValue)

}
```

Output:

```
1
```

</details>

<details>
<summary><h3>MinByComparison</h3></summary>

`MinByComparison` computes the minimum value over an [iter.Seq] sequence using a comparison function.

```go
{
	compareStringLens := func(s1, s2 string) bool {
		return len(s1) < len(s2)
	}

	values := []string{"CAT", "MOUSE", "CHICKEN"}
	minValue := gloop.MinByComparison(gloop.Slice(values), compareStringLens)
	fmt.Println(minValue)

}
```

Output:

```
CAT
```

</details>

<details>
<summary><h3>MinByComparison2</h3></summary>

`MinByComparison2` computes the minimum key and value over an [iter.Seq2] sequence using a comparison function.

```go
{
	compareKeyValueProducts := func(k1, v1, k2, v2 int) bool {
		return k1*v1 < k2*v2
	}

	m := map[int]int{
		3: 1,
		1: 5,
		4: 9,
	}

	minKey, minValue := gloop.MinByComparison2(gloop.Map(m), compareKeyValueProducts)
	fmt.Println(minKey, minValue)

}
```

Output:

```
3 1
```

</details>

<details>
<summary><h3>Product</h3></summary>

`Product` computes the product of values over an [iter.Seq] sequence. 

```go
{
	values := []int{3, 1, 4}
	prod := gloop.Product(gloop.Slice(values))
	fmt.Println(prod)

}
```

Output:

```
12
```

</details>

<details>
<summary><h3>Reduce</h3></summary>

`Reduce` runs a given function on each adjacent pair in an [iter.Seq] sequence and accumulates the result into a single value. 

```go
{
	values := []int{3, 1, 4}
	minValue := gloop.Reduce(gloop.Slice(values), func(value1 int, value2 int) int {
		return min(value1, value2)
	})
	fmt.Println(minValue)

}
```

Output:

```
1
```

</details>

<details>
<summary><h3>Reduce2</h3></summary>

`Reduce2` runs a given function on each adjacent pair of keys and values in an [iter.Seq2] sequence and accumulates the result into a single key and value pair. 

```go
{
	minKeyValueFunc := func(k1 int, v1 int, k2 int, v2 int) (int, int) {
		if v1 < v2 {
			return k1, v1
		}

		return k2, v2
	}

	m := map[int]int{
		0:	3,
		1:	1,
		2:	4,
	}

	minValueKey, minValue := gloop.Reduce2(gloop.Map(m), minKeyValueFunc)
	fmt.Println(minValueKey, minValue)

}
```

Output:

```
1 1
```

</details>

<details>
<summary><h3>Sum</h3></summary>

`Sum` computes summation over an [iter.Seq] sequence. 

```go
{
	values := []int{3, 1, 4}
	sum := gloop.Sum(gloop.Slice(values))
	fmt.Println(sum)

}
```

Output:

```
8
```

</details>

<details>
<summary><h3>ToList</h3></summary>

`ToList` converts an [iter.Seq] sequence to a [container/list.List]. 

```go
{
	seq := func(yield func(int) bool) {
		yield(3)
		yield(1)
		yield(4)
	}

	l := gloop.ToList(seq)
	fmt.Println(l.Remove(l.Front()))
	fmt.Println(l.Remove(l.Front()))
	fmt.Println(l.Remove(l.Front()))

}
```

Output:

```
3
1
4
```

</details>

<details>
<summary><h3>ToList2</h3></summary>

`ToList2` converts an [iter.Seq2] sequence to [container/list.List] of keys and values. 

```go
{
	seq := func(yield func(string, int) bool) {
		yield("CAT", 3)
		yield("DOG", 1)
		yield("MOUSE", 4)
	}

	keys, values := gloop.ToList2(seq)

	fmt.Println(keys.Remove(keys.Front()))
	fmt.Println(keys.Remove(keys.Front()))
	fmt.Println(keys.Remove(keys.Front()))
	fmt.Println(values.Remove(values.Front()))
	fmt.Println(values.Remove(values.Front()))
	fmt.Println(values.Remove(values.Front()))

}
```

Output:

```
CAT
DOG
MOUSE
3
1
4
```

</details>

<details>
<summary><h3>ToSlice</h3></summary>

`ToSlice` converts an [iter.Seq] sequence to a slice. 

```go
{
	seq := func(yield func(int) bool) {
		yield(3)
		yield(1)
		yield(4)
	}

	fmt.Println(gloop.ToSlice(seq))

}
```

Output:

```
[3 1 4]
```

</details>

<details>
<summary><h3>ToSlice2</h3></summary>

`ToSlice2` converts an [iter.Seq2] sequence to slices of keys and values. 

```go
{
	seq := func(yield func(string, int) bool) {
		yield("CAT", 3)
		yield("DOG", 1)
		yield("MOUSE", 4)
	}

	keys, values := gloop.ToSlice2(seq)
	fmt.Println(keys, values)

}
```

Output:

```
[CAT DOG MOUSE] [3 1 4]
```

</details>

<details>
<summary><h3>ToString</h3></summary>

`ToString` converts an [iter.Seq] sequence of runes to a string. 

```go
{
	seq := func(yield func(rune) bool) {
		yield('C')
		yield('A')
		yield('T')
	}

	fmt.Println(gloop.ToString(seq))

}
```

Output:

```
CAT
```

</details>

### Miscellaneous

<details>
<summary><h3>DeferLoop</h3></summary>

`DeferLoop` allows looping over an [iter.Seq] sequence, yielding a defer function that can register another function to be executed at the end of the currently running loop. If multiple functions are registered, they are executed in FIFO order. 

```go
{
	values := []int{3, 1, 4}
	for i, deferLoop := range gloop.DeferLoop(gloop.Slice(values)) {
		deferLoop(func() {
			fmt.Println("defer loop", i)
		})

		fmt.Println("regular loop", i)
	}

}
```

Output:

```
regular loop 3
defer loop 3
regular loop 1
defer loop 1
regular loop 4
defer loop 4
```

</details>

<details>
<summary><h3>Parallelize</h3></summary>

`Parallelize` runs a function on each value in an [iter.Seq] sequence on separate goroutines. 

```go
{
	printlnWithDelay := func(s string) {
		time.Sleep(time.Second)
		fmt.Println(s)
	}

	values := []string{"CAT", "DOG", "MOUSE"}
	timeElaped := time.Now()

	gloop.Parallelize(gloop.Slice(values), printlnWithDelay)
	fmt.Println("Time Elapsed", time.Since(timeElaped))

}
```

Output:

```
DOG
MOUSE
CAT
Time Elapsed 1.00134375s
```

</details>

<details>
<summary><h3>Parallelize2</h3></summary>

`Parallelize2` runs a function on each value in an [iter.Seq2] sequence on separate goroutines. 

```go
{
	printlnWithDelay := func(k string, v int) {
		time.Sleep(time.Second)
		fmt.Println(k, v)
	}

	m := map[string]int{
		"CAT":		3,
		"DOG":		1,
		"MOUSE":	4,
	}
	timeElaped := time.Now()

	gloop.Parallelize2(gloop.Map(m), printlnWithDelay)
	fmt.Println("Time Elapsed", time.Since(timeElaped))

}
```

Output:

```
MOUSE 4
CAT 3
DOG 1
Time Elapsed 1.00058975s
```

</details>


[iter.Seq]: https://pkg.go.dev/iter#Seq
[iter.Seq2]: https://pkg.go.dev/iter#Seq2
[container/list.List]: https://pkg.go.dev/container/list#List
