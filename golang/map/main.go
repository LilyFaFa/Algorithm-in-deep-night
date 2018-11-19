package main

import (
	"fmt"
)

func main() {
	f := make(map[int]string)
	f[1] = "one"
	f[2] = "two"
	fmt.Println(f[2])
	delete(f, 1)
	fmt.Println(f)
	delete(f, 3)
}
