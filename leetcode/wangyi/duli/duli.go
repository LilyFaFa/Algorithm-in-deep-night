package main

import (
	"fmt"
)

func main() {
	var x int
	var f int
	var d int
	var p int
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &f)
	fmt.Scanf("%d", &d)
	fmt.Scanf("%d", &p)
	fmt.Println(duli(x, f, d, p))

}
func duli(x int, f int, d int, p int) int {
	if d/x < f {
		return d / x
	} else {
		return (d-f*x)/(x+p) + f
	}

}
