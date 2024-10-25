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
