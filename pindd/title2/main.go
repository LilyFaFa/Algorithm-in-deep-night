package main

import (
	"fmt"
)

func main() {
	var number int
	for {
		n, _ := fmt.Scan(&number)
		if n == 0 {
			break
		} else {
			inputs := []int{}

			for i := 0; i < number; i++ {
				var a int
				fmt.Scan(&a)
				inputs = append(inputs, a)
			}

			prehit := 0
			curhit := -1
			size := 0
			for prehit != curhit {
				prehit = curhit
				curhit = 0

				size++
				cache := make([]int, size)
				cache[0] = -1

				for i := 0; i < len(inputs); i++ {
					if cache[inputs[i]%size] == inputs[i] {
						curhit++
					} else {
						cache[inputs[i]%size] = inputs[i]
					}
				}
				fmt.Println(prehit, curhit, size)
			}

			if size-1 <= 0 {
				fmt.Println(1)
			} else {
				fmt.Println(size - 1)
			}

		}
	}
}
