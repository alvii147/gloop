package main

import (
	"container/list"
	"fmt"

	"github.com/alvii147/gloop"
)

func main() {
	l := list.New()
	l.PushBack(3)
	l.PushBack(1)
	l.PushBack(4)

	for elem := range gloop.List(l) {
		fmt.Println(elem.Value)
	}
}
