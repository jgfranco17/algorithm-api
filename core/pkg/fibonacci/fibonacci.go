package fibonacci

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
	var result []int
	if limit <= 0 {
		result = []int{0}
	} else if limit == 1 {
		result = []int{0, 1}
	} else {
		// Initialize the result array with the first two Fibonacci numbers
		result := []int{0, 1}

		// Generate Fibonacci numbers up to the limit
		for i := 2; ; i++ {
			nextFib := result[i-1] + result[i-2]

			// Break the loop if the next Fibonacci number exceeds the limit
			if nextFib > limit {
				break
			}

			result = append(result, nextFib)
		}
	}

	return &FibonacciSequence{
		Length:   limit + 1,
		Sequence: result,
	}
}
