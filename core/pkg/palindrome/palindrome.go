package palindrome

import (
	"strings"
)

type PalindromeCheckOutput struct {
	Word         string
	IsPalindrome bool
}

func isAlphanumeric(c byte) bool {
	return ('a' <= c && c <= 'z') || ('0' <= c && c <= '9')
}

func PalindromeCheck(word string) *PalindromeCheckOutput {
	word = strings.ToLower(word)
	left, right := 0, len(word)-1
	result := &PalindromeCheckOutput{
		Word:         word,
		IsPalindrome: false,
	}

	for left < right {
		for left < right && !isAlphanumeric(word[left]) {
			left++
		}
		for left < right && !isAlphanumeric(word[right]) {
			right--
		}
		if word[left] != word[right] {
			return result
		}
		left++
		right--
	}

	result.IsPalindrome = true
	return result
}
