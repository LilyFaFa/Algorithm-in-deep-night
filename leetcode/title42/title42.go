package title42

import (
	"fmt"
)

func trap(height []int) int {
	start := 0
	end := len(height) - 1
	water := 0
	bottom := 0
	for start < end-1 {
		for i := start + 1; i < end; i++ {
			if height[i] < min(height[start], height[end]) {
				if height[i] <= bottom {
					water += (min(height[start], height[end]) - bottom)
				} else {
					water += (min(height[start], height[end]) - height[i])
				}
				fmt.Println(water)
			}
		}
		bottom = min(height[start], height[end])
		for start < end && height[start] <= bottom {
			start++
			fmt.Println("1:", start)
		}
		for start < end && height[end] <= bottom {
			end--
			fmt.Println("2:", end)
		}
	}
	return water
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
