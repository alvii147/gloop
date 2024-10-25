package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	s := "CAT"
	for seq := range gloop.Combinations(gloop.String(s), 2) {
		comb := gloop.ToString(seq)
		fmt.Println(comb)
	}
}
