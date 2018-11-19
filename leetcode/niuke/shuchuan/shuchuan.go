package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var result []string
	for {
		var num int
		//使用control+D结束输入循环
		n, _ := fmt.Scan(&num)
		if n == 0 {
			break
		}
		var numbers []string
		var max string
		for i := 0; i < num; i++ {
			var n string
			fmt.Scanf("%s", &n)
			numbers = append(numbers, n)
		}
		sort.Strings(numbers)
		flag := 0
		for i := 0; i < len(numbers)-1; i++ {
			//if strings.HasPrefix(numbers[i+1], numbers[i]) && (len(numbers[i+1]) == len(numbers[i]) || (len(numbers[i+1]) > len(numbers[i]) && numbers[i+1][len(numbers[i])] < numbers[i][0])) {
			if strings.HasPrefix(numbers[i+1], numbers[i]) {
				if len(numbers[i+1]) == len(numbers[i]) {
					flag++
				} else {
					for k := len(numbers[i]); k < len(numbers[i+1]); k++ {
						if numbers[i+1][k] < numbers[i][k%len(numbers[i])] {
							flag++
							break
						} else if numbers[i+1][k] > numbers[i][k%len(numbers[i])] {
							break
						}
					}
				}

			} else if flag != 0 {
				f := flag
				for j := 0; j < (f+1)/2; j++ {
					tmp := numbers[i]
					numbers[i] = numbers[i-flag]
					numbers[i-flag] = tmp
					flag--
					i--
				}
				flag = 0
			}
		}
		for i := num - 1; i >= 0; i-- {
			max = max + numbers[i]
		}
		fmt.Println(numbers)
		result = append(result, max)
	}
	for _, r := range result {
		fmt.Println(r)
	}

}
