package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	for i := range gloop.Linspace(2, 3, 5) {
		fmt.Println(i)
	}
}
