package solution

func DifferenceOfSums(n int, m int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		if i%m == 0 {
			sum -= i
		} else {
			sum += i
		}
	}
	return sum
}
