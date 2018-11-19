package main

import (
	"fmt"
)

func main() {
	a := 0
	for {
		n, _ := fmt.Scan(&a)
		if n == 0 {
			break
		} else {
			m := f(a)
			fmt.Println(m)
		}
	}
}
func f(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return f(n-1) + f(n-2)

}
