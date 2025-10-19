package algos

import (
	"strings"
	"unicode"
)

func IsPalindrome(input string) bool {
	runes := []rune(input)
	left, right := 0, len(runes)-1

	for left < right {
		for left < right && !isAlphaNumeric(runes[left]) {
			left++
		}
		for left < right && !isAlphaNumeric(runes[right]) {
			right--
		}

		if strings.ToLower(string(runes[left])) != strings.ToLower(string(runes[right])) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}
