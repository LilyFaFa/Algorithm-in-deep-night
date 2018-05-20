package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	itemAsbrand := make(map[int]int)

	for _, num := range nums {
		_, exist := itemAsbrand[num]
		if exist {
			continue
		}
		// 是否存在左边邻居
		brand1, ok1 := itemAsbrand[num-1]
		// 是否存在右边邻居
		brand2, ok2 := itemAsbrand[num+1]
		//只有左邻居,更新边界，自己是右边的边界
		if ok1 && !ok2 {
			itemAsbrand[num-brand1] = itemAsbrand[num-brand1] + 1
			itemAsbrand[num] = itemAsbrand[num-brand1]
			fmt.Println("hello1")
			fmt.Println(num, itemAsbrand[num])
			fmt.Println(num-brand1, itemAsbrand[num-brand1])
		} else if !ok1 && ok2 {
			itemAsbrand[num+brand2] = itemAsbrand[num+brand2] + 1
			itemAsbrand[num] = itemAsbrand[num+brand2]
			fmt.Println("hello2")
			fmt.Println(num, itemAsbrand[num])
			fmt.Println(num+brand2, itemAsbrand[num+brand2])
		} else if ok1 && ok2 {
			itemAsbrand[num-brand1] = itemAsbrand[num-brand1] + itemAsbrand[num+brand2] + 1
			itemAsbrand[num+brand2] = itemAsbrand[num-brand1]
			//ensure num in itemAsbrand,the value is not import
			itemAsbrand[num] = 1
			fmt.Println("hello3")
			fmt.Println(num-brand1, itemAsbrand[num-brand1])
			fmt.Println(num+brand2, itemAsbrand[num+brand2])
		} else {
			itemAsbrand[num] = 1
			fmt.Println("hello4")
			fmt.Println(num, itemAsbrand[num])

		}

	}
	max := 0
	for _, value := range itemAsbrand {
		if value > max {
			max = value
		}
	}
	return max

}
