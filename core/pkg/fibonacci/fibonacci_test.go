package fibonacci

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		limit    int
		expected *FibonacciSequence
	}{
		{1, &FibonacciSequence{Length: 2, Sequence: []uint64{0, 1}}},
		{5, &FibonacciSequence{Length: 6, Sequence: []uint64{0, 1, 1, 2, 3}}},
		{10, &FibonacciSequence{Length: 11, Sequence: []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}}},
	}

	for _, test := range tests {
		result := Fibonacci(test.limit)

		if len(result.Sequence) != len(test.expected.Sequence) {
			t.Errorf("For limit=%d, expected sequence length %d but got %d", test.limit, len(test.expected.Sequence), len(result.Sequence))
		}

		for i, val := range result.Sequence {
			if val != test.expected.Sequence[i] {
				t.Errorf("For limit=%d, at index %d, expected %d but got %d", test.limit, i, test.expected.Sequence[i], val)
			}
		}
	}
}

func generateFibonacci(limit int) []uint64 {
	result := []uint64{0, 1}
	for i := 2; i < limit; i++ {
		nextFib := result[i-1] + result[i-2]
		result = append(result, nextFib)
	}
	return result
}

func TestFibonacciValue(t *testing.T) {
	tests := []struct {
		number   int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{5, 5},
		{10, 55},
	}

	for _, test := range tests {
		result := fibonacciValue(test.number)
		if result != test.expected {
			t.Errorf("For number=%d, expected %d but got %d", test.number, test.expected, result)
		}
	}
}
