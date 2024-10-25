package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	for i := range gloop.Interval(3, 9, 2) {
		fmt.Println(i)
	}
}
