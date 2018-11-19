package main

import (
	"fmt"
	"io"
	"sort"
)

func main() {
	for {
		var point int
		_, err := fmt.Scanf("%d", &point)
		if err == io.EOF {
			break
		} else {
			var worker int
			fmt.Scanf("%d", &worker)

			var weights []int
			for i := 0; i < point; i++ {
				var a int
				fmt.Scanf("%d", &a)
				weights = append(weights, a)
			}

			sort.Ints(weights)
			first := 0
			last := len(weights) - 1

			result := 0
			capacity := worker
			for first < last {
				flag := false
				if weights[last] <= capacity {
					flag = true
					capacity = capacity - weights[last]
					last--
				}
				if weights[first] <= capacity {
					flag = true
					capacity = capacity - weights[first]
					first++
				}
				if !flag && first < last {
					capacity = worker
					result++
					//fmt.Println(first, last)
				}
			}

			if first == last {
				if capacity < weights[first] {
					result++
					//fmt.Println(result)
				}
			}
			if capacity == 0 {
				result = result + 1
			}
			fmt.Println(result)
		}

	}
}
