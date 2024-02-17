package solution

func MakeEqual(words []string) bool {
	if len(words) == 1 {
		return true
	}
	h := make(map[uint8]int)
	c := 0
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			h[words[i][j]]++
			c++
		}
	}
	if c%len(words) != 0 {
		return false
	}
	a := -1
	// 遍历哈希表，如果
	//     哪个字母的数值不够分配给每个字符串（value < len(words)）
	//     哪个字母的数值不够分配给每个字符串（value < len(words)）
	for _, v := range h {
		if a != -1 {
			if v%len(words) != 0 || v < len(words) {
				return false
			}
		} else {
			a = v
		}
	}
	return true
}
