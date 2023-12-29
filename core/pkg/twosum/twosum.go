package twosum

type TwoSumResult struct {
	Found   bool
	Indices []int
}

func TwoSum(nums []int, target int) *TwoSumResult {
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

	return &TwoSumResult{
		Found:   false,
		Indices: nil,
	}
}
