package palindrome

import (
	"context"
	"strings"

	"github.com/jgfranco17/algorithm-api/service/pkg/logging"
)

func isAlphanumeric(c byte) bool {
	return ('a' <= c && c <= 'z') || ('0' <= c && c <= '9')
}

func PalindromeCheck(word string) bool {
	ctx := context.WithValue(context.Background(), "section", "Palindrome")
	log := logging.GetLogger(ctx)
	word = strings.ToLower(word)
	left, right := 0, len(word)-1
	log.Infof("Checking if '%s' palindrome", word)
	for left < right {
		for left < right && !isAlphanumeric(word[left]) {
			left++
		}
		for left < right && !isAlphanumeric(word[right]) {
			right--
		}
		if word[left] != word[right] {
			return false
		}
		left++
		right--
	}

	return true
}
