package algos

func LongestSubstring(input string) int {
	longest := 0
	left := 0
	counter := make(map[rune]int)

	runes := []rune(input)

	for right, value := range runes {
		counter[value]++

		if counter[value] > 1 {
			counter[runes[left]]--
			left++
		}

		longest = max(longest, right-left+1)
	}

	return longest
}
