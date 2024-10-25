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
