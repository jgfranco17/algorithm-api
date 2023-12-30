package utils

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jgfranco17/algorithm-api/service/pkg/logging"
)

func ParseArray(sequence string) ([]int, error) {
	ctx := context.WithValue(context.Background(), "section", "ParseArray")
	log := logging.GetLogger(ctx)
	var numbers []int
	numberStrings := strings.Split(sequence, "-")
	for _, numStr := range numberStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Warnf("Failed to parse string '%s': %v", numStr, err)
			return nil, fmt.Errorf("")
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}
