package algos

func TwoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		value, ok := numToIndex[complement]

		if ok {
			return []int{value, i}
		}

		numToIndex[num] = i
	}

	return nil
}
