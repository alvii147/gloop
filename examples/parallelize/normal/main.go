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
	gloop.Parallelize(gloop.Slice(values), printlnWithDelay)
	fmt.Println("Time Elapsed", time.Since(timeElaped))
}
