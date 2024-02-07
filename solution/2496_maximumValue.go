package solution

import "strconv"

func MaximumValue(strs []string) int {
	count := 0
	for i := 0; i < len(strs); i++ {
		atoi, err := strconv.Atoi(strs[i])
		if err != nil {
			// 如果转换不了数字，则为字符串长度
			if count < len(strs[i]) {
				count = len(strs[i])
			}
			continue
		}
		if count < atoi {
			count = atoi
		}
	}
	return count
}
