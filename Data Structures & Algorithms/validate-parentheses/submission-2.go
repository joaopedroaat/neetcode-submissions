func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stck := []rune{}

	for _, p := range s {
		switch p {
		case ')', ']', '}':
			if len(stck) == 0 {
				return false
			}

			valid := false
			switch stck[len(stck)-1] {
			case '(':
				valid = p == ')'
			case '[':
				valid = p == ']'
			case '{':
				valid = p == '}'
			}

			if !valid {
				return false
			}

			stck = stck[:len(stck)-1]
		default:
			stck = append(stck, p)
		}
	}

	return len(stck) == 0
}

