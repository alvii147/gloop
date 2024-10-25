package main

import (
	"fmt"
	"time"

	"github.com/alvii147/gloop"
)

func printlnWithDelay(s string) {
	time.Sleep(time.Second)
	fmt.Println(s)
}

func main() {
	values := []string{"CAT", "DOG", "MOUSE"}
	timeElaped := time.Now()
	gloop.Parallelize(
		gloop.Slice(values),
		func(s string) {
			printlnWithDelay(s)
			fmt.Println(time.Since(timeElaped))
		},
		gloop.WithParallelizeMaxThreads(2),
	)
	fmt.Println("Time Elapsed", time.Since(timeElaped))
}
