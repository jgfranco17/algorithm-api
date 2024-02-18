package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		limit    int
		expected []uint64
	}{
		{limit: 1, expected: []uint64{0, 1}},
		{limit: 5, expected: []uint64{0, 1, 1, 2, 3}},
		{limit: 10, expected: []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}},
	}

	for _, test := range tests {
		result, err := Fibonacci(test.limit)
		assert.NoError(t, err)
		assert.Equal(t, test.expected, result)
	}
}
