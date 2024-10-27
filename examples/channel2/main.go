package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	ch := make(chan string)
	go func() {
		ch <- "CAT"
		ch <- "DOG"
		ch <- "MOUSE"
		close(ch)
	}()

	for i, value := range gloop.Channel2(ch) {
		fmt.Println(i, value)
	}
}
