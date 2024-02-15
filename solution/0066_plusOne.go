package solution

import "fmt"

func PlusOne(digits []int) []int {
	// 先加一
	digits[len(digits)-1]++
	for i := len(digits) - 1; i >= 0; i-- {
		fmt.Printf("%v ", digits[i])
		// 如果有某一位进到10了，则进位，如果前面没有数字了，则添加
		if digits[i] == 10 {
			// 进位，将本位变为0
			digits[i] = 0
			// 检查i-1是否越界
			// 如果越界了
			if i-1 < 0 {
				digits = append([]int{1}, digits...)
				break
			} else {
				digits[i-1]++
			}
		} else {
			break
		}
	}
	return digits
}
