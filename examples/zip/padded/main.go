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
