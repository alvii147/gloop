package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	values := []int{3, 1, 4}
	for i, value := range gloop.Reverse2(gloop.Enumerate(gloop.Slice(values))) {
		fmt.Println(i, value)
	}
}
