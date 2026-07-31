func removeElement(nums []int, val int) int {
	insertPointer, checkPointer := 0, 0
	k := 0

	for checkPointer < len(nums) {
		if nums[checkPointer] != val {
			nums[insertPointer] = nums[checkPointer]
			if checkPointer != insertPointer {
				nums[checkPointer] = -1
			}
			k++
			insertPointer++
		} else {
			nums[checkPointer] = -1
		}

		checkPointer++
	}

	return k
}
