package twosum

func TwoSum(nums []int, target int) []int {
	numIndices := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, found := numIndices[complement]; found {
			return []int{j, i}
		}

		numIndices[num] = i
	}

	return nil
}
