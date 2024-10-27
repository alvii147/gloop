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

	for key, value := range gloop.Map(m) {
		fmt.Println(key, value)
	}
}
