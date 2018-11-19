package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	for {
		n, _ := fmt.Scan(&input)
		if n == 0 {
			break
		} else {
			result := 0
			str := []string{}
			buildIPs(&result, input, str, 0)
			fmt.Println(result)
		}
	}
}
func buildIPs(res *int, s string, str []string, index int) {
	if len(str) == 4 {
		if index == len(s) {
			*res++
		}
		return
	}

	for i := index; i < index+3 && i < len(s); i++ {
		candidate := s[index : i+1]
		if num, _ := strconv.Atoi(candidate); num > 255 {
			break
		}

		str = append(str, candidate)
		buildIPs(res, s, str, i+1)
		str = str[:len(str)-1]

		if candidate == "0" {
			break
		}
	}
}
