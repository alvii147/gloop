package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	m1 := map[string]int{
		"CAT": 3,
		"DOG": 1,
	}

	m2 := map[string]int{
		"MOUSE": 4,
	}

	for key, value := range gloop.Chain2(gloop.Map(m1), gloop.Map(m2)) {
		fmt.Println(key, value)
	}
}
