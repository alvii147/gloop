package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	for r := range gloop.String("CAT") {
		fmt.Println(string(r))
	}
}
