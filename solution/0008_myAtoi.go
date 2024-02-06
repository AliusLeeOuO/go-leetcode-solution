package solution

import "math"

func MyAtoi(s string) int {
	var result int
	var negative bool
	var hasSignOrDigit bool

	for i := 0; i < len(s); i++ {
		// 跳过前导空格
		if s[i] == ' ' && !hasSignOrDigit {
			continue
		}
		//检测正号或负号
		if (s[i] == '+' || s[i] == '-') && !hasSignOrDigit {
			negative = s[i] == '-'
			hasSignOrDigit = true
			continue
		}
		// 读取数字并构建结果
		if s[i] >= '0' && s[i] <= '9' {
			hasSignOrDigit = true
			result = result*10 + int(s[i]-'0')
			// 检查越界
			if negative {
				if -result < math.MinInt32 {
					return math.MinInt32
				}
			} else {
				if result > math.MaxInt32 {
					return math.MaxInt32
				}
			}
		} else {
			break
		}
	}
	if negative {
		result = -result
	}
	// 检查越界
	if -result < math.MinInt32 {
		return math.MinInt32
	} else if result > math.MaxInt32 {
		return math.MaxInt32
	}
	return result
}
