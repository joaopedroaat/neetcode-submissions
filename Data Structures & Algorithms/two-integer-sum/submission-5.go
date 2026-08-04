func twoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int)

	for i, n := range nums {
		numToIndex[n] = i
	}

	for i := range nums {
		diff := target - nums[i]

		j, exists := numToIndex[diff]
		if exists && j != i {
			return []int{min(i, j), max(i, j)}
		}
	}

	return []int{0, 1}
}

