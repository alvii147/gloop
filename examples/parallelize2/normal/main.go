package main

import (
	"fmt"
	"time"

	"github.com/alvii147/gloop"
)

func printlnWithDelay(k string, v int) {
	time.Sleep(time.Second)
	fmt.Println(k, v)
}

func main() {
	m := map[string]int{
		"CAT":   3,
		"DOG":   1,
		"MOUSE": 4,
	}
	timeElaped := time.Now()

	gloop.Parallelize2(gloop.Map(m), printlnWithDelay)
	fmt.Println("Time Elapsed", time.Since(timeElaped))
}
