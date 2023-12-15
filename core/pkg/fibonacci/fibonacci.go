package fibonacci

import (
	"log"
)

type FibonacciSequence struct {
	Length   int
	Sequence []int
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
	result := []int{0}
	if limit == 1 {
		result = append(result, 1)
	} else {
		// Generate Fibonacci numbers up to the limit
		result = []int{0, 1}
		for i := 2; i < limit; i++ {
			nextFib := result[i-1] + result[i-2]
			result = append(result, nextFib)
		}
	}

	log.Printf("Generated Fibonacci number sequence for %d: %v", limit, result)
	return &FibonacciSequence{
		Length:   len(result),
		Sequence: result,
	}
}
