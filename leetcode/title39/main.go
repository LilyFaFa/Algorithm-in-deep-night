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

	var result1 [][]int
	var used []int
	var arr1 []int
	arrange(n, testcase, used, arr1, &result1)
	fmt.Println(result1)

	var arr2 []int
	var result2 [][]int
	combination(n, testcase, 0, arr2, &result2)
	fmt.Println(result2)

	result := combinationSum(testcase, n)
	fmt.Println(result)
	/*
	   以下仅用于测试slice的类型属性
	*/
	rr := []int{1, 2, 3}
	modify(rr, 0)
	fmt.Println(rr)
	modify(rr, 1)
	fmt.Println(rr)

}

func modify(i []int, j int) {
	i[j] = 3
}
