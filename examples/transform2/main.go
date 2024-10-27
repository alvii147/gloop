package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func concat(s1, s2 string) string {
	return s1 + s2
}

func main() {
	m := map[string]string{
		"CAT":   "DOG",
		"MOUSE": "CHICKEN",
	}

	for s := range gloop.Transform2(gloop.Map(m), concat) {
		fmt.Println(s)
	}
}
