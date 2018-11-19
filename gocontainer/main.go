package main

import (
	"container/list"

	"fmt"
)

func main() {
	list := list.New()
	list.PushBack(1)
	list.PushBack("4")
	fmt.Println(list.Front())
}
