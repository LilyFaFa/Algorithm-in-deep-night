package main

import (
	"fmt"
)

func main() {
	var n int
	for {
		l, _ := fmt.Scan(&n)
		if l == 0 {
			break
		} else {
			for i := 0; i < n; i++ {
				var a, b, c int
				fmt.Scan(&a)
				fmt.Scan(&b)
				fmt.Scan(&c)
				if a%b == 0 && c != 0 {
					fmt.Println("NO")
				} else {
					fmt.Println("YES")
				}
			}

		}
	}
}
