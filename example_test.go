package gloop_test

import (
	"container/list"
	"fmt"
	"iter"
	"strings"
	"time"

	"github.com/alvii147/gloop"
)

func ExampleInterval() {
	for i := range gloop.Interval(3, 9, 2) {
		fmt.Println(i)
	}
	// Output:
	// 3
	// 5
	// 7
}

func ExampleLinspace() {
	for i := range gloop.Linspace(2, 3, 5) {
		fmt.Println(i)
	}
	// Output:
	// 2
	// 2.25
	// 2.5
	// 2.75
}

func ExampleRandomNormal() {
	for i := range gloop.RandomNormal(2, 2, 5) {
		fmt.Println(i)
	}
	// Output:
	// 3.6053409359773543
	// 4.8493077906535165
	// 1.321369004660313
	// 1.3549030774712296
	// -0.6521572615302738
}

func ExampleRandomUniform() {
	for i := range gloop.RandomUniform(0, 2, 5) {
		fmt.Println(i)
	}
	// Output:
	// 1.7336396942444041
	// 0.9684446802268123
	// 1.5762348358917075
	// 0.5000463191262544
	// 1.1113562403363295
}

func ExampleAll() {
	allTrue := []bool{true, true, true}
	someTrue := []bool{true, false, true}
	allFalse := []bool{false, false, false}

	fmt.Println(gloop.All(gloop.Slice(allTrue)))
	fmt.Println(gloop.All(gloop.Slice(someTrue)))
	fmt.Println(gloop.All(gloop.Slice(allFalse)))
	// Output:
	// true
	// false
	// false
}

func ExampleAny() {
	allTrue := []bool{true, true, true}
	someTrue := []bool{true, false, true}
	allFalse := []bool{false, false, false}

	fmt.Println(gloop.Any(gloop.Slice(allTrue)))
	fmt.Println(gloop.Any(gloop.Slice(someTrue)))
	fmt.Println(gloop.Any(gloop.Slice(allFalse)))
	// Output:
	// true
	// true
	// false
}

func ExampleBatch() {
	values := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	for seq := range gloop.Batch(gloop.Slice(values), 3) {
		batch := gloop.ToSlice(seq)
		fmt.Println(batch)
	}
	// Output:
	// [3 1 4]
	// [1 5 9]
	// [2 6 5]
}

func ExampleBatch2() {
	values := []string{"CAT", "DOG", "MOUSE", "CHICKEN", "BUNNY", "BEAR"}
	for seq := range gloop.Batch2(gloop.Enumerate(gloop.Slice(values)), 3) {
		batchKeys, batchValues := gloop.ToSlice2(seq)
		fmt.Println(batchKeys, batchValues)
	}
	// Output:
	// [0 1 2] [CAT DOG MOUSE]
	// [3 4 5] [CHICKEN BUNNY BEAR]
}

func ExampleCartesianProduct() {
	s := "CAT"
	for seq := range gloop.CartesianProduct(gloop.String(s), 2) {
		product := gloop.ToString(seq)
		fmt.Println(product)
	}
	// Output:
	// CC
	// CA
	// CT
	// AC
	// AA
	// AT
	// TC
	// TA
	// TT
}

func ExampleCartesianProduct2() {
	m := map[string]int{
		"CAT": 3,
		"DOG": 1,
	}
	for seq := range gloop.CartesianProduct2(gloop.Map(m), 2) {
		productKeys, productValues := gloop.ToSlice2(seq)
		fmt.Println(productKeys, productValues)
	}
	// Output:
	// [CAT CAT] [3 3]
	// [CAT DOG] [3 1]
	// [DOG CAT] [1 3]
	// [DOG DOG] [1 1]
}

func ExampleChain() {
	values1 := []int{3, 1, 4}
	values2 := []int{1, 6}
	for i := range gloop.Chain(gloop.Slice(values1), gloop.Slice(values2)) {
		fmt.Println(i)
	}
	// Output:
	// 3
	// 1
	// 4
	// 1
	// 6
}

func ExampleChain2() {
	m1 := map[string]int{
		"CAT": 3,
		"DOG": 1,
	}

	m2 := map[string]int{
		"MOUSE": 4,
	}

	for key, value := range gloop.Chain2(gloop.Map(m1), gloop.Map(m2)) {
		fmt.Println(key, value)
	}
	// Output:
	// CAT 3
	// DOG 1
	// MOUSE 4
}

func ExampleChannel() {
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
	// Output:
	// CAT
	// DOG
	// MOUSE
}

func ExampleCollect() {
	for i := range gloop.Collect(3, 1, 4) {
		fmt.Println(i)
	}
	// Output:
	// 3
	// 1
	// 4
}

func ExampleCombinations() {
	s := "CAT"
	for seq := range gloop.Combinations(gloop.String(s), 2) {
		comb := gloop.ToString(seq)
		fmt.Println(comb)
	}
	// Output:
	// CA
	// CT
	// AT
}

func ExampleCombinations2() {
	m := map[string]int{
		"CAT":   3,
		"DOG":   1,
		"MOUSE": 4,
	}
	for seq := range gloop.Combinations2(gloop.Map(m), 2) {
		combKeys, combValues := gloop.ToSlice2(seq)
		fmt.Println(combKeys, combValues)
	}
	// Output:
	// [CAT DOG] [3 1]
	// [CAT MOUSE] [3 4]
	// [DOG MOUSE] [1 4]
}

func ExampleDeferLoop() {
	values := []int{3, 1, 4}
	for i, deferLoop := range gloop.DeferLoop(gloop.Slice(values)) {
		deferLoop(func() {
			fmt.Println("defer loop", i)
		})

		fmt.Println("regular loop", i)
	}
	// Output:
	// regular loop 3
	// defer loop 3
	// regular loop 1
	// defer loop 1
	// regular loop 4
	// defer loop 4
}

func ExampleEnumerate() {
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
	// Output:
	// 0 3
	// 1 1
	// 2 4
}

func ExampleEqual() {
	values1 := []int{3, 1, 4}
	values2 := []int{3, 1, -4}
	values3 := []int{3, 1, 4}

	fmt.Println(gloop.Equal(gloop.Slice(values1), gloop.Slice(values2)))
	fmt.Println(gloop.Equal(gloop.Slice(values1), gloop.Slice(values3)))
	// Output:
	// false
	// true
}

func ExampleEqual2() {
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
	// Output:
	// false
	// true
}

func ExampleEquivalent() {
	values1 := []int{3, 1, 4}
	values2 := []int{3, 1, -4}
	values3 := []int{3, 4, 1}

	fmt.Println(gloop.Equivalent(gloop.Slice(values1), gloop.Slice(values2)))
	fmt.Println(gloop.Equivalent(gloop.Slice(values1), gloop.Slice(values3)))
	// Output:
	// false
	// true
}

func ExampleEquivalent2() {
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
	// Output:
	// false
	// true
}

func ExampleFilter() {
	isOdd := func(i int) bool {
		return i%2 == 1
	}

	values := []int{3, 1, 4}
	for i := range gloop.Filter(gloop.Slice(values), isOdd) {
		fmt.Println(i)
	}
	// Output:
	// 3
	// 1
}

func ExampleFilter2() {
	isProductPositive := func(i, j int) bool {
		return i*j >= 0
	}

	m := map[int]int{
		-3: 3,
		-1: -1,
		4:  4,
	}

	for i, j := range gloop.Filter2(gloop.Map(m), isProductPositive) {
		fmt.Println(i, j)
	}
	// Output:
	// -1 -1
	// 4 4
}

func ExampleFold() {
	add := func(a, b int) int {
		return a + b
	}

	values := []int{3, 1, 4}
	sum := gloop.Fold(gloop.Slice(values), add)
	fmt.Println(sum)
	// Output:
	// 8
}

func ExampleFold2() {
	addKeyValueProduct := func(acc, key, value int) int {
		return acc + (key * value)
	}

	m := map[int]int{
		3: 1,
		1: 5,
		4: 9,
	}

	sumOfProducts := gloop.Fold2(gloop.Map(m), addKeyValueProduct)
	fmt.Println(sumOfProducts)
	// Output:
	// 44
}

func ExampleKeys() {
	m := map[string]int{
		"CAT":   3,
		"DOG":   1,
		"MOUSE": 4,
	}

	for key := range gloop.Keys(gloop.Map(m)) {
		fmt.Println(key)
	}
	// Output:
	// CAT
	// DOG
	// MOUSE
}

func ExampleKeyValue() {
	pairs := []gloop.KeyValuePair[string, int]{
		{
			Key:   "CAT",
			Value: 3,
		},
		{
			Key:   "DOG",
			Value: 1,
		},
		{
			Key:   "MOUSE",
			Value: 4,
		},
	}

	for key, value := range gloop.KeyValue(gloop.Slice(pairs)) {
		fmt.Println(key, value)
	}
	// Output:
	// CAT 3
	// DOG 1
	// MOUSE 4
}

func ExampleKeyValue2() {
	m := map[string]int{
		"CAT":   3,
		"DOG":   1,
		"MOUSE": 4,
	}

	for pair := range gloop.KeyValue2(gloop.Map(m)) {
		fmt.Println(pair.Key, pair.Value)
	}
	// Output:
	// CAT 3
	// DOG 1
	// MOUSE 4
}

func ExampleList() {
	l := list.New()
	l.PushBack(3)
	l.PushBack(1)
	l.PushBack(4)

	for elem := range gloop.List(l) {
		fmt.Println(elem.Value)
	}
	// Output:
	// 3
	// 1
	// 4
}

func ExampleMap() {
	m := map[string]int{
		"CAT":   3,
		"DOG":   1,
		"MOUSE": 4,
	}

	for key, value := range gloop.Map(m) {
		fmt.Println(key, value)
	}
	// Output:
	// CAT 3
	// DOG 1
	// MOUSE 4
}

func ExampleMax() {
	values := []int{3, 1, 4, 2}
	maxValue := gloop.Max(gloop.Slice(values))
	fmt.Println(maxValue)
	// Output:
	// 4
}

func ExampleMean() {
	values := []int{3, 1, 4, 2}
	mean := gloop.Mean(gloop.Slice(values))
	fmt.Println(mean)
	// Output:
	// 2.5
}

func ExampleMin() {
	values := []int{3, 1, 4, 2}
	minValue := gloop.Min(gloop.Slice(values))
	fmt.Println(minValue)
	// Output:
	// 1
}

func ExampleParallelize() {
	printlnWithDelay := func(s string) {
		time.Sleep(time.Second)
		fmt.Println(s)
	}

	values := []string{"CAT", "DOG", "MOUSE"}
	timeElaped := time.Now()

	gloop.Parallelize(gloop.Slice(values), printlnWithDelay)
	fmt.Println("Time Elapsed", time.Since(timeElaped))
	// Output:
	// DOG
	// MOUSE
	// CAT
	// Time Elapsed 1.00134375s
}

func ExampleParallelize2() {
	printlnWithDelay := func(k string, v int) {
		time.Sleep(time.Second)
		fmt.Println(k, v)
	}

	m := map[string]int{
		"CAT":   3,
		"DOG":   1,
		"MOUSE": 4,
	}
	timeElaped := time.Now()

	gloop.Parallelize2(gloop.Map(m), printlnWithDelay)
	fmt.Println("Time Elapsed", time.Since(timeElaped))
	// Output:
	// MOUSE 4
	// CAT 3
	// DOG 1
	// Time Elapsed 1.00058975s
}

func ExamplePermutations() {
	s := "CAT"
	for seq := range gloop.Permutations(gloop.String(s), 2) {
		perm := gloop.ToString(seq)
		fmt.Println(perm)
	}
	// Output:
	// CA
	// CT
	// AC
	// AT
	// TC
	// TA
}

func ExamplePermutations2() {
	m := map[string]int{
		"CAT":   3,
		"DOG":   1,
		"MOUSE": 4,
	}
	for seq := range gloop.Permutations2(gloop.Map(m), 2) {
		permKeys, permValues := gloop.ToSlice2(seq)
		fmt.Println(permKeys, permValues)
	}
	// Output:
	// [CAT DOG] [3 1]
	// [CAT MOUSE] [3 4]
	// [DOG CAT] [1 3]
	// [DOG MOUSE] [1 4]
	// [MOUSE CAT] [4 3]
	// [MOUSE DOG] [4 1]
}

func ExampleProduct() {
	values := []int{3, 1, 4}
	prod := gloop.Product(gloop.Slice(values))
	fmt.Println(prod)
	// Output:
	// 12
}

func ExampleReduce() {
	values := []int{3, 1, 4}
	minValue := gloop.Reduce(gloop.Slice(values), func(value1 int, value2 int) int {
		return min(value1, value2)
	})
	fmt.Println(minValue)
	// Output:
	// 1
}

func ExampleReduce2() {
	minKeyValueFunc := func(k1 int, v1 int, k2 int, v2 int) (int, int) {
		if v1 < v2 {
			return k1, v1
		}

		return k2, v2
	}

	m := map[int]int{
		0: 3,
		1: 1,
		2: 4,
	}

	minValueKey, minValue := gloop.Reduce2(gloop.Map(m), minKeyValueFunc)
	fmt.Println(minValueKey, minValue)
	// Output:
	// 1 1
}

func ExampleReverse() {
	values := []int{3, 1, 4}
	for i := range gloop.Reverse(gloop.Slice(values)) {
		fmt.Println(i)
	}
	// Output:
	// 4
	// 1
	// 3
}

func ExampleReverse2() {
	values := []int{3, 1, 4}
	for i, value := range gloop.Reverse2(gloop.Enumerate(gloop.Slice(values))) {
		fmt.Println(i, value)
	}
	// Output:
	// 2 4
	// 1 1
	// 0 3
}

func ExampleSlice() {
	values := []int{3, 1, 4}
	for i := range gloop.Slice(values) {
		fmt.Println(i)
	}
	// Output:
	// 3
	// 1
	// 4
}

func ExampleSort() {
	values := []int{3, 1, 4, 1, 5, 9}
	for i := range gloop.Sort(gloop.Slice(values), true) {
		fmt.Println(i)
	}
	// Output:
	// 1
	// 1
	// 3
	// 4
	// 5
	// 9
}

func ExampleSortByComparison() {
	compareStringLens := func(s1, s2 string) bool {
		return len(s1) < len(s2)
	}

	values := []string{"CAT", "MOUSE", "DOG"}
	for s := range gloop.SortByComparison(gloop.Slice(values), compareStringLens, true) {
		fmt.Println(s)
	}
	// Output:
	// CAT
	// DOG
	// MOUSE
}

func ExampleSortByComparison2() {
	compareKeyValueConcatLen := func(k1, v1, k2, v2 string) bool {
		return len(k1+v1) < len(k2+v2)
	}

	values := map[string]string{
		"CAT":   "DOG",
		"MOUSE": "CHICKEN",
		"BUNNY": "BEAR",
	}
	for key, value := range gloop.SortByComparison2(gloop.Map(values), compareKeyValueConcatLen, true) {
		fmt.Println(key, value)
	}
	// Output:
	// CAT DOG
	// BUNNY BEAR
	// MOUSE CHICKEN
}

func ExampleSortByRank() {
	stringLen := func(s string) int {
		return len(s)
	}

	values := []string{"CAT", "MOUSE", "DOG"}
	for s := range gloop.SortByRank(gloop.Slice(values), stringLen, true) {
		fmt.Println(s)
	}
	// Output:
	// CAT
	// DOG
	// MOUSE
}

func ExampleSortByRank2() {
	stringConcatLen := func(k1, v1 string) int {
		return len(k1 + v1)
	}

	values := map[string]string{
		"CAT":   "DOG",
		"MOUSE": "CHICKEN",
		"BUNNY": "BEAR",
	}
	for key, value := range gloop.SortByRank2(gloop.Map(values), stringConcatLen, true) {
		fmt.Println(key, value)
	}
	// Output:
	// CAT DOG
	// BUNNY BEAR
	// MOUSE CHICKEN
}

func ExampleString() {
	for r := range gloop.String("CAT") {
		fmt.Println(string(r))
	}
	// Output:
	// C
	// A
	// T
}

func ExampleSum() {
	values := []int{3, 1, 4}
	sum := gloop.Sum(gloop.Slice(values))
	fmt.Println(sum)
	// Output:
	// 8
}

func ExampleToList() {
	seq := func(yield func(int) bool) {
		yield(3)
		yield(1)
		yield(4)
	}

	l := gloop.ToList(seq)
	fmt.Println(l.Remove(l.Front()))
	fmt.Println(l.Remove(l.Front()))
	fmt.Println(l.Remove(l.Front()))
	// Output:
	// 3
	// 1
	// 4
}

func ExampleToList2() {
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
	// Output:
	// CAT
	// DOG
	// MOUSE
	// 3
	// 1
	// 4
}

func ExampleToSlice() {
	seq := func(yield func(int) bool) {
		yield(3)
		yield(1)
		yield(4)
	}

	fmt.Println(gloop.ToSlice(seq))
	// Output:
	// [3 1 4]
}

func ExampleToSlice2() {
	seq := func(yield func(string, int) bool) {
		yield("CAT", 3)
		yield("DOG", 1)
		yield("MOUSE", 4)
	}

	keys, values := gloop.ToSlice2(seq)
	fmt.Println(keys, values)
	// Output:
	// [CAT DOG MOUSE] [3 1 4]
}

func ExampleToString() {
	seq := func(yield func(rune) bool) {
		yield('C')
		yield('A')
		yield('T')
	}

	fmt.Println(gloop.ToString(seq))
	// Output:
	// CAT
}

func ExampleTransform() {
	values := []string{"CaT", "dOg"}
	for s := range gloop.Transform(gloop.Slice(values), strings.ToUpper) {
		fmt.Println(s)
	}
	// Output:
	// CAT
	// DOG
}

func ExampleTransform2() {
	concat := func(s1, s2 string) string {
		return s1 + s2
	}

	m := map[string]string{
		"CAT":   "DOG",
		"MOUSE": "CHICKEN",
	}

	for s := range gloop.Transform2(gloop.Map(m), concat) {
		fmt.Println(s)
	}
	// Output:
	// CATDOG
	// MOUSECHICKEN
}

func ExampleValues() {
	m := map[string]int{
		"CAT":   3,
		"DOG":   1,
		"MOUSE": 4,
	}

	for value := range gloop.Values(gloop.Map(m)) {
		fmt.Println(value)
	}
	// Output:
	// 3
	// 1
	// 4
}

func ExampleWindow() {
	values := []int{3, 1, 4, 1, 5, 9}
	for seq := range gloop.Window(gloop.Slice(values), 3) {
		window := gloop.ToSlice(seq)
		fmt.Println(window)
	}
	// Output:
	// [3 1 4]
	// [1 4 1]
	// [4 1 5]
	// [1 5 9]
}

func ExampleWindow2() {
	values := []string{"CAT", "DOG", "MOUSE", "CHICKEN", "BUNNY", "BEAR"}
	for seq := range gloop.Window2(gloop.Enumerate(gloop.Slice(values)), 3) {
		windowKeys, windowValues := gloop.ToSlice2(seq)
		fmt.Println(windowKeys, windowValues)
	}
	// Output:
	// [0 1 2] [CAT DOG MOUSE]
	// [1 2 3] [DOG MOUSE CHICKEN]
	// [2 3 4] [MOUSE CHICKEN BUNNY]
	// [3 4 5] [CHICKEN BUNNY BEAR]
}

func ExampleZip() {
	values1 := []string{"CAT", "DOG", "MOUSE"}
	values2 := []int{3, 1, 4}
	for value1, value2 := range gloop.Zip(gloop.Slice(values1), gloop.Slice(values2)) {
		fmt.Println(value1, value2)
	}
	// Output:
	// CAT 3
	// DOG 1
	// MOUSE 4
}

func ExampleZip2() {
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
	// Output:
	// CAT 3 3 1.2
	// DOG 1 1 3.4
	// MOUSE 4 4 5.6
}

func ExampleZipN() {
	seq1 := gloop.Slice([]string{"CAT", "DOG"})
	seq2 := gloop.Slice([]string{"MOUSE", "CHICKEN"})
	seq3 := gloop.Slice([]string{"BUNNY", "BEAR"})

	for seq := range gloop.ZipN(gloop.Collect(seq1, seq2, seq3)) {
		fmt.Println(gloop.ToSlice(seq))
	}
	// Output:
	// [CAT MOUSE BUNNY]
	// [DOG CHICKEN BEAR]
}

func ExampleZipN2() {
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
	// Output:
	// [CAT MOUSE] [3 1]
	// [DOG BUNNY] [1 5]
	// [MOUSE BEAR] [4 9]
}
