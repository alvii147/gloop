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
