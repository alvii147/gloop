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
