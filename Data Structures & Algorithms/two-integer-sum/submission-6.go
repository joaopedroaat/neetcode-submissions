func twoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int)

	for i, v := range nums {
		diff := target - v

		j, exists := numToIndex[diff]
		if exists {
			return []int{j, i}
		}

		numToIndex[v] = i
	}

	return nil
}

