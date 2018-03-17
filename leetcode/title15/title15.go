package main

import (
	"fmt"
)

func threeSum(nums []int) [][]int {
	sort := sortNum(nums)
	result := [][]int{}
	for middle := 0; middle < len(sort); middle++ {
		first := 0
		last := len(sort) - 1
		for first < middle && last > middle {
			if sort[first]+sort[last]+sort[middle] == 0 {
				insert := []int{sort[first], sort[middle], sort[last]}
				result = append(result, insert)
				last--
				first++
				// 避免中间值固定，两边值重复导致的重复子串
				for sort[first] == sort[first-1] && first < middle {
					first++
				}
				for sort[last] == sort[last+1] && last > middle {
					last--
				}
			} else if sort[first]+sort[last]+sort[middle] > 0 {
				last--
			} else {
				first++
			}

		}
		//避免中间值相同导致重复字符串
		//避免出现0,0,0,0和-1,-2,-2,3,4情况冲突采用下面的方式
		firstRepeat := false
		for middle < len(sort)-1 && sort[middle+1] == sort[middle] {
			if !firstRepeat {
				for i := middle + 2; i < len(sort); i++ {
					if sort[middle]+sort[middle+1]+sort[i] == 0 {
						fmt.Println("insert:", middle)
						insert := []int{sort[middle], sort[middle+1], sort[i]}
						result = append(result, insert)
						// 避免重复
						break
					}
				}
			}
			firstRepeat = true
			middle++
		}

	}
	return result
}

//冒泡排序
func sortNum(nums []int) []int {
	tmp := 0
	for i, _ := range nums {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j-1] > nums[j] {
				tmp = nums[j-1]
				nums[j-1] = nums[j]
				nums[j] = tmp
			}
		}
	}
	return nums
}

func main() {
	input := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	output := sortNum(input)
	fmt.Println(output)
	result := threeSum(input)
	fmt.Println(result)

}
