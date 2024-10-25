package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	values := []int{3, 1, 4}
	for i := range gloop.Slice(values) {
		fmt.Println(i)
	}
}
