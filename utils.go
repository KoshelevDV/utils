package utils

func Contains(a []string, b string) bool {
	for _, n := range a {
		if b == n {
			return true
		}
	}
	return false
}
