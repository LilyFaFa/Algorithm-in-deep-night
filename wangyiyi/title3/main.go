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
			for i := 0; i < a; i++ {
				result := 0
				x := 0
				y := 0
				fmt.Scan(&x)
				fmt.Scan(&y)
				if x%2 == 1 {
					result = x / 2
					tmp := x/2 + 1
					result = result + tmp - y
				} else {
					result = x/2 - 1
					tmp := result + 2
					result = result + tmp - y
				}
				fmt.Print(0)
				fmt.Print(" ")
				fmt.Println(result)
			}

		}
	}
}
