package solution

func RearrangeCharacters(s string, target string) int {
	// finalCount
	finalCount := 0
	// 统计字符串 s 中每个字符的出现次数。
	sCount := make(map[uint8]int)
	// 统计目标字符串 target 中每个字符的出现次数。
	targetCount := make(map[uint8]int)
	for i := 0; i < len(target); i++ {
		targetCount[target[i]]++
	}
	for i := 0; i < len(s); i++ {
		sCount[s[i]]++
	}
	// 取值
	flag := true
	for flag == true {
		// 循环targetCount取出字符
		for index, value := range targetCount {
			if sCount[index]-value < 0 {
				flag = false
				break
			} else {
				sCount[index] -= value
			}
		}
		finalCount++
	}
	return finalCount - 1
}
