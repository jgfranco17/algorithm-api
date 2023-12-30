package maxsubarray

import (
	"context"

	"github.com/jgfranco17/algorithm-api/service/pkg/logging"
)

func MaxSubArray(nums []int) int {
	ctx := context.WithValue(context.Background(), "section", "MaxSubArray")
	log := logging.GetLogger(ctx)

	maxSum, currentSum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}

	log.Infof("Calculated maximum subarray for %v", nums)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
