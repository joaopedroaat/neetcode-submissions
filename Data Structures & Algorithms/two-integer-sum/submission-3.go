func twoSum(nums []int, target int) []int {
	i, j := 0, 1

	for i <= len(nums)-2 {
		j = i + 1
		for j < len(nums) {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
			j++
		}
		i++
	}

	return []int{i, j}
}

