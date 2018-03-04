package title11

import (
	"fmt"
)

func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	maxA := (end - start) * min(height[start], height[end])
	for start < end {
		area := (end - start) * min(height[start], height[end])
		maxA = max(maxA, area)
		if height[start] < height[end] {
			start++
			fmt.Println("1:", start)
		} else if height[start] > height[end] {
			end--
			fmt.Println("2:", end)
		} else {
			start++
			end--
		}
	}
	return maxA
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
