package main

import (
	"fmt"
)

func main() {
	var num int
	var n int
	var testcase []int
	fmt.Scanf("%d", &num)
	fmt.Scanf("%d", &n)
	for i := 0; i < num; i++ {
		var item int
		fmt.Scanf("%d", &item)
		testcase = append(testcase, item)
	}
	result := combinationSum(testcase, n)
	fmt.Println(result)
}
