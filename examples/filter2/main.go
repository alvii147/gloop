package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func isProductPositive(i int, j int) bool {
	return i*j >= 0
}

func main() {
	m := map[int]int{
		-3: 3,
		-1: -1,
		4:  4,
	}
	for i, j := range gloop.Filter2(gloop.Map(m), isProductPositive) {
		fmt.Println(i, j)
	}
}
