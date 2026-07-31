func findMaxConsecutiveOnes(nums []int) int {
	currentCount, maxCount := 0, 0

	for _, n := range nums {

		if n == 1 {
			currentCount = currentCount + 1
		} else {
			if currentCount > maxCount {
				maxCount = currentCount
			}

			currentCount = 0
		}
	}

	if currentCount > maxCount {
		maxCount = currentCount
	}

	return maxCount
}
