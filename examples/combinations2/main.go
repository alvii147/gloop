package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	m := map[string]int{
		"CAT":   3,
		"DOG":   1,
		"MOUSE": 4,
	}
	for seq := range gloop.Combinations2(gloop.Map(m), 2) {
		combKeys, combValues := gloop.ToSlice2(seq)
		fmt.Println(combKeys, combValues)
	}
}
