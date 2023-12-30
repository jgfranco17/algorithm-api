package utils

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jgfranco17/algorithm-api/service/pkg/logging"
)

func StringToIntArray(sequence []string) ([]int, error) {
	ctx := context.WithValue(context.Background(), "section", "ParseArray")
	log := logging.GetLogger(ctx)
	var numbers []int
	for _, numStr := range sequence {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Warnf("Failed to parse string '%s': %v", numStr, err)
			return nil, fmt.Errorf("Failed to parse string '%s': %v", numStr, err)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}
