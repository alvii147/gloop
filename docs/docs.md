### Generators

<details>
<summary><h4>Interval</h2></summary>

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
<summary><h4>Linspace</h2></summary>

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
<summary><h4>RandomUniform</h2></summary>

`RandomUniform` allows looping over a given number of random values drawn from a uniform distribution. The size must not be negative. 

```go
{
	for i := range gloop.RandomUniform(0, 2, 5) {
		fmt.Println(i)
	}

}
```

</details>

<details>
<summary><h4>RandomNormal</h2></summary>

`RandomNormal` allows looping over a given number of random values drawn from a Gaussian distribution. The size must not be negative and the standard deviation must be positive. 

```go
{
	for i := range gloop.RandomNormal(2, 2, 5) {
		fmt.Println(i)
	}

}
```

</details>

### Scalar Iterators

<details>
<summary><h4>Chain</h2></summary>

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

</details>

<details>
<summary><h4>Chain2</h2></summary>

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

</details>

<details>
<summary><h4>Channel</h2></summary>

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

</details>

<details>
<summary><h4>Collect</h2></summary>

`Collect` allows looping over a given set of values. 

```go
{
	for i := range gloop.Collect(3, 1, 4) {
		fmt.Println(i)
	}

}
```

</details>

<details>
<summary><h4>Enumerate</h2></summary>

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

</details>

<details>
<summary><h4>Filter</h2></summary>

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

</details>

<details>
<summary><h4>Filter2</h2></summary>

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

</details>

<details>
<summary><h4>Keys</h2></summary>

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

</details>

<details>
<summary><h4>KeyValue</h2></summary>

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

</details>

<details>
<summary><h4>KeyValue2</h2></summary>

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

</details>

<details>
<summary><h4>List</h2></summary>

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

</details>

<details>
<summary><h4>Map</h2></summary>

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

</details>

<details>
<summary><h4>Reverse</h2></summary>

`Reverse` allows looping over an [iter.Seq] sequence in order of descending index. 

```go
{
	values := []int{3, 1, 4}
	for i := range gloop.Reverse(gloop.Slice(values)) {
		fmt.Println(i)
	}

}
```

</details>

<details>
<summary><h4>Reverse2</h2></summary>

`Reverse2` allows looping over an [iter.Seq2] sequence in order of descending index. 

```go
{
	values := []int{3, 1, 4}
	for i, value := range gloop.Reverse2(gloop.Enumerate(gloop.Slice(values))) {
		fmt.Println(i, value)
	}

}
```

</details>

<details>
<summary><h4>Slice</h2></summary>

`Slice` allows looping over a given slice. 

```go
{
	values := []int{3, 1, 4}
	for i := range gloop.Slice(values) {
		fmt.Println(i)
	}

}
```

</details>

<details>
<summary><h4>Sort</h2></summary>

`Sort` allows looping over an [iter.Seq] sequence in sorted order. 

```go
{
	values := []int{3, 1, 4, 1, 5, 9}
	for i := range gloop.Sort(gloop.Slice(values), true) {
		fmt.Println(i)
	}

}
```

</details>

<details>
<summary><h4>SortByComparison</h2></summary>

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

</details>

<details>
<summary><h4>SortByComparison2</h2></summary>

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

</details>

<details>
<summary><h4>SortByRank</h2></summary>

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

</details>

<details>
<summary><h4>SortByRank2</h2></summary>

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

</details>

<details>
<summary><h4>String</h2></summary>

`String` allows looping over the runes in a given string. 

```go
{
	for r := range gloop.String("CAT") {
		fmt.Println(string(r))
	}

}
```

</details>

<details>
<summary><h4>Transform</h2></summary>

`Transform` runs a given function on each value over an [iter.Seq] sequence and allows looping over the returned values. 

```go
{
	values := []string{"CaT", "dOg"}
	for s := range gloop.Transform(gloop.Slice(values), strings.ToUpper) {
		fmt.Println(s)
	}

}
```

</details>

<details>
<summary><h4>Transform2</h2></summary>

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

</details>

<details>
<summary><h4>Values</h2></summary>

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

</details>

<details>
<summary><h4>Zip</h2></summary>

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

</details>

<details>
<summary><h4>Zip2</h2></summary>

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

</details>

### Vector Iterators

<details>
<summary><h4>Batch</h2></summary>

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

</details>

<details>
<summary><h4>Batch2</h2></summary>

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

</details>

<details>
<summary><h4>CartesianProduct</h2></summary>

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

</details>

<details>
<summary><h4>CartesianProduct2</h2></summary>

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

</details>

<details>
<summary><h4>Combinations</h2></summary>

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

</details>

<details>
<summary><h4>Combinations2</h2></summary>

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

</details>

<details>
<summary><h4>Permutations</h2></summary>

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

</details>

<details>
<summary><h4>Permutations2</h2></summary>

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

</details>

<details>
<summary><h4>Window</h2></summary>

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

</details>

<details>
<summary><h4>Window2</h2></summary>

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

</details>

<details>
<summary><h4>ZipN</h2></summary>

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

</details>

<details>
<summary><h4>ZipN2</h2></summary>

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

</details>

### Aggregators

<details>
<summary><h4>All</h2></summary>

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

</details>

<details>
<summary><h4>Any</h2></summary>

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

</details>

<details>
<summary><h4>Equal</h2></summary>

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

</details>

<details>
<summary><h4>Equal2</h2></summary>

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

</details>

<details>
<summary><h4>Equivalent</h2></summary>

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

</details>

<details>
<summary><h4>Equivalent2</h2></summary>

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

</details>

<details>
<summary><h4>Fold</h2></summary>

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

</details>

<details>
<summary><h4>Fold2</h2></summary>

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

</details>

<details>
<summary><h4>Max</h2></summary>

`Max` computes the maximum value over an [iter.Seq] sequence. 

```go
{
	values := []int{3, 1, 4, 2}
	maxValue := gloop.Max(gloop.Slice(values))
	fmt.Println(maxValue)

}
```

</details>

<details>
<summary><h4>Mean</h2></summary>

`Mean` computes the mean value over an [iter.Seq] sequence. 

```go
{
	values := []int{3, 1, 4, 2}
	mean := gloop.Mean(gloop.Slice(values))
	fmt.Println(mean)

}
```

</details>

<details>
<summary><h4>Min</h2></summary>

`Min` computes the minimum value over an [iter.Seq] sequence. 

```go
{
	values := []int{3, 1, 4, 2}
	minValue := gloop.Min(gloop.Slice(values))
	fmt.Println(minValue)

}
```

</details>

<details>
<summary><h4>Product</h2></summary>

`Product` computes the product of values over an [iter.Seq] sequence. 

```go
{
	values := []int{3, 1, 4}
	prod := gloop.Product(gloop.Slice(values))
	fmt.Println(prod)

}
```

</details>

<details>
<summary><h4>Reduce</h2></summary>

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

</details>

<details>
<summary><h4>Reduce2</h2></summary>

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

</details>

<details>
<summary><h4>Sum</h2></summary>

`Sum` computes summation over an [iter.Seq] sequence. 

```go
{
	values := []int{3, 1, 4}
	sum := gloop.Sum(gloop.Slice(values))
	fmt.Println(sum)

}
```

</details>

<details>
<summary><h4>ToList</h2></summary>

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

</details>

<details>
<summary><h4>ToList2</h2></summary>

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

</details>

<details>
<summary><h4>ToSlice</h2></summary>

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

</details>

<details>
<summary><h4>ToSlice2</h2></summary>

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

</details>

<details>
<summary><h4>ToString</h2></summary>

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

</details>

### Miscellaneous

<details>
<summary><h4>DeferLoop</h2></summary>

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

</details>

<details>
<summary><h4>Parallelize</h2></summary>

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

</details>

<details>
<summary><h4>Parallelize2</h2></summary>

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

</details>
