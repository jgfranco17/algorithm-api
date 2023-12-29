package palindrome

import (
	"context"
	"strings"

	"github.com/jgfranco17/algorithm-api/service/pkg/logging"
)

type PalindromeCheckOutput struct {
	Word         string
	IsPalindrome bool
}

func isAlphanumeric(c byte) bool {
	return ('a' <= c && c <= 'z') || ('0' <= c && c <= '9')
}

func PalindromeCheck(word string) *PalindromeCheckOutput {
	ctx := context.WithValue(context.Background(), "section", "Palindrome")
	log := logging.GetLogger(ctx)
	word = strings.ToLower(word)
	left, right := 0, len(word)-1
	result := &PalindromeCheckOutput{
		Word:         word,
		IsPalindrome: false,
	}
	log.Infof("Checking if '%s' palindrome", word)
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
