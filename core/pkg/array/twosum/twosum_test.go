package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums            []int
		target          int
		expectedFound   bool
		expectedIndices []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, expectedFound: true, expectedIndices: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, expectedFound: true, expectedIndices: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, expectedFound: true, expectedIndices: []int{0, 1}},
		{nums: []int{1, 2, 3, 4, 5}, target: 9, expectedFound: true, expectedIndices: []int{3, 4}},
		{nums: []int{1, 2, 3, 4, 5}, target: 20, expectedFound: false, expectedIndices: nil},
	}

	for _, test := range tests {
		found, indices := TwoSum(test.nums, test.target)
		assert.Equal(t, found, test.expectedFound)
		assert.Equal(t, indices, test.expectedIndices)
	}
}
