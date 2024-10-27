package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	values := []string{"CAT", "DOG", "MOUSE", "CHICKEN", "BUNNY", "BEAR"}
	for seq := range gloop.Window2(gloop.Enumerate(gloop.Slice(values)), 3) {
		windowKeys, windowValues := gloop.ToSlice2(seq)
		fmt.Println(windowKeys, windowValues)
	}
}
