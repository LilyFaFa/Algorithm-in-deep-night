package main

import (
	"fmt"
)

func main() {
	var n int
	for {
		l, _ := fmt.Scan(&n)
		if l == 0 {
			break
		} else {
			sc := smallestCommons(1, n)
			m := n + 1
			flag := false

			sc1 := m
			sc2 := LCM(sc, m)

			for !flag {
				if sc1 == sc2 {
					flag = true
				} else {
					m++
					sc1 = LCM(m, sc1)
					sc2 = LCM(m, sc2)
				}
			}
			fmt.Println(m)
		}
	}
}

func smallestCommons(first int, last int) int {
	var nums []int
	for i := first; i < last+1; i++ {
		nums = append(nums, i)
	}

	multiple := nums[0]
	for _, n := range nums {
		multiple = LCM(n, multiple)
	}
	return multiple
}

func GCD(num1 int, num2 int) int {
	r := num1 % num2
	if r == 0 {
		return num2
	}
	return GCD(num2, r)
}

func LCM(num1 int, num2 int) int {
	if num1 > num2 {
		return num1 * num2 / GCD(num1, num2)
	} else {
		return num1 * num2 / GCD(num2, num1)
	}

}
