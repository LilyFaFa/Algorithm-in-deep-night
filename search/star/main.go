package main

import (
	"fmt"
	"strconv"
)

func main() {
	exam := "123456"
	result := 0
	search(exam, 0, &result)
	fmt.Println(result)
}

// cut 代表从 input 的第 cut 个位置，cur 是指目前组合数目
func search(input string, cur int, result *int) {
	// 最后一块
	if cur == 5 {
		if len(input) == 0 {
			return
		}
		num, _ := strconv.Atoi(input)
		if 0 <= num && num <= 600 {
			*result += 1
		}
	}

	for i := 1; i < 3; i++ {
		if len(input) < i {
			break
		}
		num, _ := strconv.Atoi(input[0:i])
		if num >= 0 && num <= 600 {
			search(input[i:], cur+1, result)
		}

	}

}
