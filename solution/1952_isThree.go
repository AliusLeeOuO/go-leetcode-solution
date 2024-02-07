package solution

func IsThree(n int) bool {
	if n == 0 {
		return false
	}
	count := 1
	for i := 2; i < n; i++ {
		if n%i == 0 {
			count++
		}
		if count > 2 {
			return false
		}
	}
	if count == 2 {
		return true
	}
	return false
}
