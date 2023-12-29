package fibonacci

import (
	"context"

	"github.com/jgfranco17/algorithm-api/service/pkg/logging"
)

type FibonacciSequence struct {
	Length   int
	Sequence []uint64
}

func fibonacciValue(number int) int {
	if number <= 1 {
		return number
	}

	a, b := 0, 1
	for i := 2; i <= number; i++ {
		a, b = b, a+b
	}
	return b
}

func Fibonacci(limit int) *FibonacciSequence {
	ctx := context.WithValue(context.Background(), "section", "Fibonacci")
	log := logging.GetLogger(ctx)
	result := []uint64{0}
	hardCap := 93
	if limit == 1 {
		result = append(result, 1)
	} else {
		if limit >= hardCap {
			log.Warnf("Maximum limit is %d", hardCap)
			limit = hardCap - 1
		}
		result = []uint64{0, 1}
		for i := 2; i < limit; i++ {
			nextFib := result[i-1] + result[i-2]
			result = append(result, nextFib)
		}
	}

	log.Printf("Generated Fibonacci number sequence for %d", limit)
	return &FibonacciSequence{
		Length:   len(result),
		Sequence: result,
	}
}
