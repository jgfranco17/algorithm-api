package twosum

import (
	"context"

	"github.com/jgfranco17/algorithm-api/service/pkg/logging"
)

func TwoSum(nums []int, target int) (bool, []int) {
	ctx := context.WithValue(context.Background(), "section", "TwoSum")
	log := logging.GetLogger(ctx)
	numIndices := make(map[int]int)
	log.Infof("Finding target=%d in %v", nums, target)

	for i, num := range nums {
		complement := target - num
		if j, found := numIndices[complement]; found {
			return true, []int{j, i}
		}
		numIndices[num] = i
	}
	log.Infof("No valid result found.")
	return false, nil
}
