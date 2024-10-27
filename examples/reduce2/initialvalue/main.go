package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func multiplyKeyValue(acc, key, value int) int {
	return acc * key * value
}

func main() {
	m := map[int]int{
		3: 1,
		1: 5,
		4: 9,
	}

	product := gloop.Reduce2(gloop.Map(m), multiplyKeyValue, gloop.WithReduceInitialValue(1))
	fmt.Println(product)
}
