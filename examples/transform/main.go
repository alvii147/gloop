package main

import (
	"fmt"
	"strings"

	"github.com/alvii147/gloop"
)

func main() {
	values := []string{"CaT", "dOg"}
	for s := range gloop.Transform(gloop.Slice(values), strings.ToLower) {
		fmt.Println(s)
	}
}
