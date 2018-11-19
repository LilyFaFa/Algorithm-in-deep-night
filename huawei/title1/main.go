package main

import (
	"fmt"
	"io"
	"sort"
)

func main() {
	var result []float64
	for {
		var num int
		_, err := fmt.Scanf("%d", &num)
		if err == io.EOF {
			break
		} else {
			var sum float64
			var g float64
			sum = 0
			var nums []float64
			for i := 0; i < num; i++ {
				fmt.Scanf("%f", &g)
				nums = append(nums, g)
			}
			sort.Float64s(nums)
			for i := 1; i < num-1; i++ {
				sum += nums[i]
			}
			result = append(result, sum/(float64(num)-float64(2)))
		}
	}

	for _, r := range result {
		l := fmt.Sprintf("%.2f", r)
		fmt.Println(l)
	}
}
