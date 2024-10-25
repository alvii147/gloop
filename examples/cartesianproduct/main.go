package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	s := "CAT"
	for seq := range gloop.CartesianProduct(gloop.String(s), 2) {
		product := gloop.ToString(seq)
		fmt.Println(product)
	}
}
