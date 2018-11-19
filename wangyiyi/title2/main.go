package main

import (
	"fmt"
)

func main() {
	var s string
	for {
		n, _ := fmt.Scan(&s)
		if n == 0 {
			break
		} else {
			result := 1
			i := 0
			for i < len(s) {
				flag := true
				for j := 1; j < len(s); j++ {
					if s[(i+j)%len(s)] == s[(i+j-1)%len(s)] {
						if result < j {
							result = j
						}
						i = i + j
						flag = false
						break
					}
				}
				if flag {
					result = len(s)
					break
				}
			}
			fmt.Println(result)
		}
	}
}
