package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	s := "CAT"
	for seq := range gloop.Permutations(gloop.String(s), 2) {
		perm := gloop.ToString(seq)
		fmt.Println(perm)
	}
}
