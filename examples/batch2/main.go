package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	values := []string{"CAT", "DOG", "MOUSE", "CHICKEN", "BUNNY", "BEAR"}
	for seq := range gloop.Batch2(gloop.Enumerate(gloop.Slice(values)), 3) {
		batchKeys, batchValues := gloop.ToSlice2(seq)
		fmt.Println(batchKeys, batchValues)
	}
}
