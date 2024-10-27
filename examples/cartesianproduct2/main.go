package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	m := map[string]int{
		"CAT": 3,
		"DOG": 1,
	}
	for seq := range gloop.CartesianProduct2(gloop.Map(m), 2) {
		productKeys, productValues := gloop.ToSlice2(seq)
		fmt.Println(productKeys, productValues)
	}
}
