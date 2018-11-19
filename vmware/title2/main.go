package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		var point int
		_, err := fmt.Scanf("%d", &point)
		if err == io.EOF {
			break
		} else {
			var intputs []int
			for i := 0; i < point; i++ {
				var a int
				fmt.Scanf("%d", &a)
				intputs = append(intputs, a)
			}

			min := 0
			max := 0
			for i := 0; i < len(intputs)-1; i++ {
				for j := i + 1; j < len(intputs); j++ {
					if intputs[i] > intputs[j] {
						min++
					} else {
						max++
					}
				}
			}
			if min > max {
				fmt.Println(max)
			} else {
				fmt.Println(min)
			}

		}
	}

}
