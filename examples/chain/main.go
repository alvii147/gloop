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
