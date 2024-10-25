package main

import (
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 3
		ch <- 1
		ch <- 4
		close(ch)
	}()

	for i := range gloop.Channel(ch) {
		fmt.Println(i)
	}
}
