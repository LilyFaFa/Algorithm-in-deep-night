package main

import (
	"fmt"
)

func main() {
	var input string
	for {
		n, _ := fmt.Scan(&input)
		if n == 0 {
			break
		} else {
			result := lengthOfLongestSubstring(input)
			fmt.Println(result)
		}
	}
}

func lengthOfLongestSubstring(s string) int {
	tmp := make(map[byte]int)
	start := -1
	maxlen := 0

	for i := 0; i < len(s); i++ {
		last, ok := tmp[s[i]]
		if ok {
			// 如果在 start 之后(不包含start) 已经出现过一次了
			if last > start {
				if maxlen < (i - start - 1) {
					maxlen = i - start - 1
				}
				start = last
			}
		}
		tmp[s[i]] = i
		if i == len(s)-1 && (!ok || last <= start) {
			if maxlen < (i - start) {
				maxlen = i - start
			}
		}
	}

	return maxlen
}
