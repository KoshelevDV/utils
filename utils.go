package utils

func InSlice(a []string, b string) bool {
	for _, n := range a {
		if b == n {
			return true
		}
	}
	return false
}

func ContainsInt(a []int, b int) bool {
	for _, n := range a {
		if b == n {
			return true
		}
	}
	return false
}
