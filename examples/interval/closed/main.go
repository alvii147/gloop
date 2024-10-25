package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	for i := range gloop.Interval(3, 9, 2, gloop.WithIntervalClosed(true)) {
		fmt.Println(i)
	}
}
