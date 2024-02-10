package solution

func CountOperations(num1 int, num2 int) int {
	if num1 == 0 || num2 == 0 {
		return 0
	} else if num1 == num2 {
		return 1
	}
	count := 0
	for num1 != 0 && num2 != 0 {
		count++
		if num1 == num2 {
			break
		} else if num1 > num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
	}
	return count
}
