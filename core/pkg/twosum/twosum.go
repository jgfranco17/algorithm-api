package twosum

import (
	"context"

	"github.com/jgfranco17/algorithm-api/service/pkg/logging"
)

type TwoSumResult struct {
	Found   bool
	Indices []int
}

func TwoSum(nums []int, target int) *TwoSumResult {
	ctx := context.WithValue(context.Background(), "section", "TwoSum")
	log := logging.GetLogger(ctx)
	numIndices := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, found := numIndices[complement]; found {
			return &TwoSumResult{
				Found:   true,
				Indices: []int{j, i},
			}
		}
		numIndices[num] = i
	}
	log.Infof("No valid result found.")
	return &TwoSumResult{
		Found:   false,
		Indices: nil,
	}
}
