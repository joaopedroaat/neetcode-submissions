func calPoints(operations []string) int {
	record := []int{}

	for _, op := range operations {
		switch op {
		case "+":
			record = append(record, record[len(record)-2]+record[len(record)-1])
		case "D":
			record = append(record, record[len(record)-1]*2)
		case "C":
			record = record[:len(record)-1]
		default:
			n, _ := strconv.Atoi(op)
			record = append(record, n)
		}
	}

	total := 0

	for _, n := range record {
		total += n
	}

	return total
}

