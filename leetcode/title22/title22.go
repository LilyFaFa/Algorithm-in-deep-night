package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	result := []string{}
	backtrack(&result, "", 0, 0, n)
	return result
}

func backtrack(result *([]string), str string, open int, clos int, max int) {
	if len(str) == max*2 {
		*result = append(*result, str)
		return
	}
	if open < max {
		backtrack(result, str+"(", open+1, clos, max)
	}
	if clos < open {
		backtrack(result, str+")", open, clos+1, max)
	}
}

func main() {
	result := generateParenthesis(3)
	fmt.Println(result)
}
