func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	paths := []int{1, 2}

	for n > 2 {
		paths = append(paths, paths[len(paths)-2]+paths[len(paths)-1])
		n--
	}

	return paths[len(paths)-1]
}
