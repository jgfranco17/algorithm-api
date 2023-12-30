package maxsubarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MaxSubArrayTestInput struct {
	input    []int
	expected int
}

func TestMaxSubArray(t *testing.T) {
	testInput := []MaxSubArrayTestInput{
		{input: []int{1, 2, 3, 4, 5}, expected: 15},
		{input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, expected: 6},
		{input: []int{-1, -2, -3, -4, -5}, expected: -1},
		{input: []int{}, expected: 0},
		{input: []int{7}, expected: 7},
	}
	for _, test := range testInput {
		assert.Equal(t, test.expected, MaxSubArray(test.input))
	}
}
