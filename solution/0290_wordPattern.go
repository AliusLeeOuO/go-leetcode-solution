package solution

import "strings"

func WordPattern(pattern string, s string) bool {
	// 拆开字符串
	parts := strings.Split(s, " ")
	if len(pattern) != len(parts) {
		return false
	}
	h := make(map[uint8]string)
	for i := 0; i < len(pattern); i++ {
		// 如果找到了字符串
		if value, ok := h[pattern[i]]; ok {
			if value != parts[i] {
				return false
			}
		} else {
			for _, hashValue := range h {
				if hashValue == parts[i] {
					return false
				}
			}
			h[pattern[i]] = parts[i]
		}
	}
	return true
}
