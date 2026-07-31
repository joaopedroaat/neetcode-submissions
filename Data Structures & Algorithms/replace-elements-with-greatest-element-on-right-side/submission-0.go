func replaceElements(arr []int) []int {
	lp, rp := 0, 1
	arr[lp] = -1

	for lp < len(arr)-1 {
		if rp >= len(arr) {
			lp++
			arr[lp] = -1
			rp = lp + 1
			continue
		}

		if arr[rp] > arr[lp] {
			arr[lp] = arr[rp]
		}

		rp++
	}

	return arr
}
