[![Genocide Watch](https://hinds-banner.vercel.app/genocide-watch?variant=crimson)](https://www.pcrf.net/)

<p align="center">
    <img alt="gloop logo" src="img/logo.svg" width=500 />
</p>

<p align="center">
    <strong><i>gloop</i></strong> is a Go utility library for convenient looping using the range-over-func feature.
</p>

## Installation

```bash
go get github.com/alvii147/gloop
```

> [!NOTE]
> Go version 1.23+ required as older versions don't offer the range-over-func feature.

## Features

* [Generators](#generators)
    * [Interval](#interval)
    * [Linspace](#linspace)
* [Scalar Iterators](#scalar-iterators)
    * [Chain](#chain)
    * [Channel](#channel)
    * [Enumerate](#enumerate)
    * [Filter](#filter)
    * [List](#list)
    * [Reverse](#reverse)
    * [Slice](#slice)
    * [String](#string)
    * [Transform](#transform)
    * [Zip](#zip)
* [Vector Iterators](#vector-iterators)
    * [Batch](#batch)
    * [Cartesian Product](#cartesian-product)
    * [Combinations](#combinations)
    * [Permutations](#permutations)
    * [Window](#window)
* [Accumulators](#accumulators)
    * [All](#all)
    * [Any](#any)
    * [Max](#max)
    * [Mean](#mean)
    * [Min](#min)
    * [Product](#product)
    * [Reduce](#reduce)
    * [Sum](#sum)
    * [To List](#to-list)
    * [To Slice](#to-slice)
    * [To String](#to-string)
* [Parallelizers](#parallelizers)
    * [Parallelize](#parallelize)

## Generators

### Interval

`Interval` allows a for loop to range over values in a given interval with a given step size.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	for i := range gloop.Interval(3, 9, 2) {
		fmt.Println(i)
	}
}
```

```
3
5
7
```

`WithIntervalClosed` can be used to configure the interval as closed so the end point value is included.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	for i := range gloop.Interval(3, 9, 2, gloop.WithIntervalClosed(true)) {
		fmt.Println(i)
	}
}
```

```
3
5
7
9
```

### Linspace

`Linspace` allows a for loop to range over evenly spaced values within a given interval.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	for i := range gloop.Linspace(2, 3, 5) {
		fmt.Println(i)
	}
}
```

```
2
2.25
2.5
2.75
```

`WithLinspaceClosed` can be used to configure the interval as closed so the end point value is included.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	for i := range gloop.Linspace(2, 3, 5, gloop.WithLinspaceClosed(true)) {
		fmt.Println(i)
	}
}
```

```
2
2.25
2.5
2.75
3
```

## Scalar Iterators

### Chain

`Chain` allows a for loop to range over multiple sequences.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	values1 := []int{3, 1, 4}
	values2 := []int{1, 6}
	for i := range gloop.Chain(gloop.Slice(values1), gloop.Slice(values2)) {
		fmt.Println(i)
	}
}
```

```
3
1
4
1
6
```

### Channel

`Channel` allows a for loop to receive and range over values from a given channel.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 3
		ch <- 1
		ch <- 4
		close(ch)
	}()

	for i := range gloop.Channel(ch) {
		fmt.Println(i)
	}
}
```

```
3
1
4
```

### Enumerate

`Enumerate` allows a for loop to iterate over a sequence with an index.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
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

```
0 3
1 1
2 4
```

### Filter

`Filter` runs a given function on each value from a given sequence and allows a for loop to range over values for which the function returns true.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func isOdd(i int) bool {
	return i%2 == 1
}

func main() {
	values := []int{3, 1, 4}
	for i := range gloop.Filter(gloop.Slice(values), isOdd) {
		fmt.Println(i)
	}
}
```

```
3
1
```

### List

`List` allows a for loop to range over a given list from [container/list](https://pkg.go.dev/container/list).

```go
package main

import (
	"container/list"
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	l := list.New()
	l.PushBack(3)
	l.PushBack(1)
	l.PushBack(4)

	for elem := range gloop.List(l) {
		fmt.Println(elem.Value)
	}
}
```

```
3
1
4
```

### Reverse

`Reverse` allows a for loop to range over a given sequence in order of descending index.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	values := []int{3, 1, 4}
	for i := range gloop.Reverse(gloop.Slice(values)) {
		fmt.Println(i)
	}
}
```

```
4
1
3
```

### Slice

`Slice` allows a for loop to range over a given slice.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	values := []int{3, 1, 4}
	for i := range gloop.Slice(values) {
		fmt.Println(i)
	}
}
```

```
3
1
4
```

### String

`String` allows a for loop to range over the runes in a given string.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	for r := range gloop.String("CAT") {
		fmt.Println(string(r))
	}
}
```

```
C
A
T
```

### Transform

`Transform` runs a given function on each value from a given sequence and allows a for loop to range over the returned values.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/alvii147/gloop"
)

func main() {
	values := []string{"CaT", "dOg"}
	for s := range gloop.Transform(gloop.Slice(values), strings.ToLower) {
		fmt.Println(s)
	}
}
```

```
cat
dog
```

### Zip

`Zip` allows a for loop to iterate over two given sequences in pairs.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	values1 := []int{3, 1, 4}
	values2 := []int{1, 5, 9}
	for value1, value2 := range gloop.Zip(gloop.Slice(values1), gloop.Slice(values2)) {
		fmt.Println(value1, value2)
	}
}
```

```
3 1
1 5
4 9
```

By default, iteration terminates when the shorter sequence is exhaused.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	values1 := []int{3, 1, 4}
	values2 := []int{1, 5, 9, 2, 6, 5}
	for value1, value2 := range gloop.Zip(gloop.Slice(values1), gloop.Slice(values2)) {
		fmt.Println(value1, value2)
	}
}
```

```
3 1
1 5
4 9
```

`WithZipPadded` can be used to configured iteration to continue till the longer sequence is exhaused with the shorter sequence padded with zero-values.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	values1 := []int{3, 1, 4}
	values2 := []int{1, 5, 9, 2, 6, 5}
	for value1, value2 := range gloop.Zip(
		gloop.Slice(values1),
		gloop.Slice(values2),
		gloop.WithZipPadded[int, int](true),
	) {
		fmt.Println(value1, value2)
	}
}
```

```
3 1
1 5
4 9
0 2
0 6
0 5
```

`WithZipPadValue1` can be used to set the padding value for the shorter sequence.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	values1 := []int{3, 1, 4}
	values2 := []int{1, 5, 9, 2, 6, 5}
	for value1, value2 := range gloop.Zip(
		gloop.Slice(values1),
		gloop.Slice(values2),
		gloop.WithZipPadded[int, int](true),
		gloop.WithZipPadValue1[int, int](-1),
	) {
		fmt.Println(value1, value2)
	}
}
```

```
3 1
1 5
4 9
-1 2
-1 6
-1 5
```

## Vector Iterators

### Batch

`Batch` allows a for loop to range over a given sequence in batches of a given size.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	for seq := range gloop.Batch(gloop.Slice(values), 3) {
		batch := gloop.ToSlice(seq)
		fmt.Println(batch)
	}
}
```

```
[3 1 4]
[1 5 9]
[2 6 5]
```

### Cartesian Product

`CartesianProduct` allows a for loop to range over the Cartesian product of a given size for a given sequence.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	s := "CAT"
	for seq := range gloop.CartesianProduct(gloop.String(s), 2) {
		product := gloop.ToString(seq)
		fmt.Println(product)
	}
}
```

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

### Combinations

`Combinations` allows a for loop to range over all combinations of a given size for a given sequence.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	s := "CAT"
	for seq := range gloop.Combinations(gloop.String(s), 2) {
		comb := gloop.ToString(seq)
		fmt.Println(comb)
	}
}
```

```
CA
CT
AT
```

### Permutations

`Permutations` allows a for loop to range over all permutations of a given size for a given sequence.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	s := "CAT"
	for seq := range gloop.Permutations(gloop.String(s), 2) {
		perm := gloop.ToString(seq)
		fmt.Println(perm)
	}
}
```

```
CA
CT
AC
AT
TC
TA
```

### Window

`Window` allows a for loop to range over a given sequence in sliding windows of a given size.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	values := []int{3, 1, 4, 1, 5, 9}
	for seq := range gloop.Window(gloop.Slice(values), 3) {
		window := gloop.ToSlice(seq)
		fmt.Println(window)
	}
}
```

```
[3 1 4]
[1 4 1]
[4 1 5]
[1 5 9]
```

## Accumulators

### All

`All` computes whether or not all values in a sequence are true.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	allTrue := []bool{true, true, true}
	someTrue := []bool{true, false, true}
	allFalse := []bool{false, false, false}

	fmt.Println(gloop.All(gloop.Slice(allTrue)))
	fmt.Println(gloop.All(gloop.Slice(someTrue)))
	fmt.Println(gloop.All(gloop.Slice(allFalse)))
}
```

```
true
false
false
```

### Any

`Any` computes whether or not any value in a sequence is true.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	allTrue := []bool{true, true, true}
	someTrue := []bool{true, false, true}
	allFalse := []bool{false, false, false}

	fmt.Println(gloop.Any(gloop.Slice(allTrue)))
	fmt.Println(gloop.Any(gloop.Slice(someTrue)))
	fmt.Println(gloop.Any(gloop.Slice(allFalse)))
}
```

```
true
true
false
```

### Max

`Max` computes the maximum value over a given sequence.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	values := []int{3, 1, 4, 2}
	maxValue := gloop.Max(gloop.Slice(values))
	fmt.Println(maxValue)
}
```

```
4
```

### Mean

`Mean` computes the mean value over a given sequence.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	values := []int{3, 1, 4, 2}
	mean := gloop.Mean(gloop.Slice(values))
	fmt.Println(mean)
}
```

```
2.5
```

### Min

`Min` computes the minumum value over a given sequence.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	values := []int{3, 1, 4, 2}
	minValue := gloop.Min(gloop.Slice(values))
	fmt.Println(minValue)
}
```

```
1
```

### Product

`Product` executes the product of values a given sequence.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	values := []int{3, 1, 4}
	prod := gloop.Product(gloop.Slice(values))
	fmt.Println(prod)
}
```

```
12
```

### Reduce

`Reduce` runs a given function on each value from a given sequence and accumulates the result into a single value.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func add(a, b int) int {
	return a + b
}

func main() {
	values := []int{3, 1, 4}
	sum := gloop.Reduce(gloop.Slice(values), add)
	fmt.Println(sum)
}
```

```
8
```

`WithReduceInitialValue` can be used to set the initial value of the accumulator.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func multiply(a, b int) int {
	return a * b
}

func main() {
	values := []int{3, 1, 4}
	product := gloop.Reduce(gloop.Slice(values), multiply, gloop.WithReduceInitialValue(1))
	fmt.Println(product)
}
```

```
12
```

### Sum

`Sum` executes summation over a given sequence.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	values := []int{3, 1, 4}
	sum := gloop.Sum(gloop.Slice(values))
	fmt.Println(sum)
}
```

```
8
```

### To List

`ToList` converts a sequence to a list from [container/list](https://pkg.go.dev/container/list).

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
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

```
3
1
4
```

### To Slice

`ToSlice` converts a sequence to a slice.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	seq := func(yield func(int) bool) {
		yield(3)
		yield(1)
		yield(4)
	}

	fmt.Println(gloop.ToSlice(seq))
}
```

```
[3 1 4]
```

### To String

`ToString` converts a sequence of runes to a string.

```go
package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	seq := func(yield func(rune) bool) {
		yield('C')
		yield('A')
		yield('T')
	}

	fmt.Println(gloop.ToString(seq))
}
```

```
CAT
```

## Parallelizers

### Parallelize

`Parallelize` runs a function on each value from given a slice of values on separate goroutines.

```go
package main

import (
	"fmt"
	"time"

	"github.com/alvii147/gloop"
)

func printlnWithDelay(s string) {
	time.Sleep(time.Second)
	fmt.Println(s)
}

func main() {
	values := []string{"CAT", "DOG", "MOUSE"}
	timeElaped := time.Now()
	gloop.Parallelize(gloop.Slice(values), printlnWithDelay)
	fmt.Println("Time Elapsed", time.Since(timeElaped))
}
```

```
MOUSE
DOG
CAT
Time Elapsed 1.001531166s
```

`WithParallelizeContext` can be used to pass in a context. If the context is cancelled, no more goroutines are started.

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/alvii147/gloop"
)

func printlnWithDelay(s string) {
	time.Sleep(time.Second)
	fmt.Println(s)
}

func main() {
	values := []string{"CAT", "DOG", "MOUSE"}
	timeElaped := time.Now()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	gloop.Parallelize(gloop.Slice(values), printlnWithDelay, gloop.WithParallelizeContext(ctx))
	fmt.Println("Time Elapsed", time.Since(timeElaped))
}
```

```
Time Elapsed 50.5µs
```

`WithParallelizeMaxThreads` can be used to configure the maximum number of concurrent goroutines.

```go
package main

import (
	"fmt"
	"time"

	"github.com/alvii147/gloop"
)

func printlnWithDelay(s string) {
	time.Sleep(time.Second)
	fmt.Println(s)
}

func main() {
	values := []string{"CAT", "DOG", "MOUSE"}
	timeElaped := time.Now()
	gloop.Parallelize(
		gloop.Slice(values),
		func(s string) {
			printlnWithDelay(s)
			fmt.Println(time.Since(timeElaped))
		},
		gloop.WithParallelizeMaxThreads(2),
	)
	fmt.Println("Time Elapsed", time.Since(timeElaped))
}
```

```
CAT
DOG
1.001607s
1.001604792s
MOUSE
2.0021425s
Time Elapsed 2.002274583s
```
