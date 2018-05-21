package main

func main() {

}

func maxSubArray(nums []int) int {
	maxSum := make([]int, len(nums))
	maxSum[0] = nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if maxSum[i-1] > 0 {
			maxSum[i] = maxSum[i-1] + nums[i]
		} else {
			maxSum[i] = nums[i]
		}
		if maxSum[i] > max {
			max = maxSum[i]
		}
	}
}
