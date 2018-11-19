package search

import (
	"strconv"
)

var phone = [][]byte{
	{'a', 'b', 'c'},
	{'d', 'e', 'f'},
	{'g', 'h', 'i'},
	{'j', 'k', 'l'},
	{'m', 'n', 'o'},
	{'p', 'q', 'r', 's'},
	{'t', 'u', 'v'},
	{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	result := []string{}
	search(digits, "", &result)
	return result
}
func search(digits string, cur string, result *[]string) {
	// 1 递归终止条件
	if len(cur) == len(digits) {
		*result = append(*result, cur)
		return
	}
	// 2 处理递归
	i, _ := strconv.Atoi(string(digits[len(cur)]))
	for _, b := range phone[i-2] {
		search(digits, cur+string(b), result)
	}

}
