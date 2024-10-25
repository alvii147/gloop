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
