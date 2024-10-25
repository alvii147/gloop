package main

import (
	"context"
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

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	gloop.Parallelize(gloop.Slice(values), printlnWithDelay, gloop.WithParallelizeContext(ctx))
	fmt.Println("Time Elapsed", time.Since(timeElaped))
}
