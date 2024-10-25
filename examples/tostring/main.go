package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	seq := func(yield func(rune) bool) {
		yield('C')
		yield('A')
		yield('T')
	}

	fmt.Println(gloop.ToString(seq))
}
