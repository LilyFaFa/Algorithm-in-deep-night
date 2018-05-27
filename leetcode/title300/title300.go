package main

import (
	"fmt"
	"sort"
)

// 这个是一个比较神奇的操作，利用了插入排序，这是leetcode中时间复杂度最快的解法
// 查看sort.SearchInts(tail[:res], n)查看n在切片tail[:res]中的位置，条件是切片已经升序排序
/*
func lengthOfLIS(nums []int) int {
	tail, res := make([]int, len(nums)), 0
	for _, n := range nums {
		i := sort.SearchInts(tail[:res], n)
		tail[i] = n
		if i == res {
			res++
		}
	}
	return res
}

*/
//这个算法会记录子问题的解，但是时间复杂度还是比较高，因为需要对比比自己所有
//小的问题的解，而这个常常有包含关系
/*
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var midresult []int
	for i := 0; i < len(nums); i++ {
		midresult = append(midresult, 1)
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && midresult[j] >= midresult[i] {
				midresult[i] = midresult[j] + 1
			}
		}
	}
	var max int
	for _, m := range midresult {
		if m > max {
			max = m
		}
	}
	return max
}
*/
//解法会超时，因为没有计算中间结果
/*
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := 0
	sumlength(&max, 0, -1000000000, 0, nums)
	return max

}

func sumlength(max *int, cout int, maxCur int, i int, nums []int) {
	if i == len(nums) {
		if *max < cout {
			*max = cout
		}
		return
	}

	if nums[i] <= maxCur {
		sumlength(max, cout, maxCur, i+1, nums)
	} else {
		sumlength(max, cout, maxCur, i+1, nums)
		sumlength(max, cout+1, nums[i], i+1, nums)
	}
}
*/

func lengthOfLIS(nums []int) int {
	tail, res := make([]int, len(nums)), 0
	for _, n := range nums {
		i := sort.SearchInts(tail[:res], n)
		tail[i] = n
		if i == res {
			res++
			fmt.Println("huhu")
		}
		fmt.Println(tail)
	}
	fmt.Println(tail)
	return res
}

func main() {
	nums := []int{10, 90, 45, -2, -1, 45, 30, 5, 7, 8, -1, 6, 3, 4, 8, 9}
	result := lengthOfLIS(nums)
	fmt.Println(result)

	fmt.Println("__________")
	s := []int{5, 2, 6, 1, 4} // 未排序的切片数据
	sort.Ints(s)              //排序后的s为[1 2  4 5 6]
	fmt.Println(sort.SearchInts(s, 3))
}
