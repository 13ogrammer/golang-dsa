package algos

func MaxSumSubArray(nums []int, size int) int {
	windowSum := 0

	for i := 0; i < size; i++ {
		windowSum += nums[i]
	}

	maxSum := windowSum

	for right := size; right < len(nums); right++ {
		left := right - size
		windowSum -= nums[left]
		windowSum += nums[right]

		maxSum = max(maxSum, windowSum)
	}

	return maxSum
}
