package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type PalindromeTestInput struct {
	word     string
	expected bool
}

func TestPalindromeCheck(t *testing.T) {
	tests := []PalindromeTestInput{
		{word: "abcdcba", expected: true},
		{word: "racecar", expected: true},
		{word: "abcdefg", expected: false},
	}

	for _, test := range tests {
		result := PalindromeCheck(test.word)
		assert.Equal(t, test.expected, result)
	}
}

func TestIsAlphanumeric(t *testing.T) {
	tests := []struct {
		char     byte
		expected bool
	}{
		{'a', true},
		{'1', true},
		{'$', false},
		{'#', false},
		{'.', false},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, isAlphanumeric(test.char))
	}
}
