package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func addKeyValue(acc, key, value int) int {
	return acc + key + value
}

func main() {
	m := map[int]int{
		3: 1,
		1: 5,
		4: 9,
	}

	sum := gloop.Reduce2(gloop.Map(m), addKeyValue)
	fmt.Println(sum)
}
