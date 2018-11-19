package main

import (
	"fmt"
)

func main() {
	s := "hello"
	mdMap(s)
	fmt.Println(s)
}

func mdMap(s string) {
	s[1] = 'p'
}
