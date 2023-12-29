package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected *TwoSumResult
	}{
		{[]int{2, 7, 11, 15}, 9, &TwoSumResult{Found: true, Indices: []int{0, 1}}},
		{[]int{3, 2, 4}, 6, &TwoSumResult{Found: true, Indices: []int{1, 2}}},
		{[]int{3, 3}, 6, &TwoSumResult{Found: true, Indices: []int{0, 1}}},
		{[]int{1, 2, 3, 4, 5}, 9, &TwoSumResult{Found: true, Indices: []int{3, 4}}},
		{[]int{1, 2, 3, 4, 5}, 20, &TwoSumResult{Found: false, Indices: nil}},
	}

	for _, test := range tests {
		result := TwoSum(test.nums, test.target)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For nums=%v, target=%d, expected %v but got %v", test.nums, test.target, test.expected, result)
		}
	}
}
