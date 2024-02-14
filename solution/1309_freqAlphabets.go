package solution

import (
	"strconv"
)

func FreqAlphabets(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			newChar := s[i-2 : i]
			i = i - 2
			rep, _ := strconv.Atoi(newChar)
			res = string(rep+96) + res
		} else {
			res = string(s[i]+48) + res
		}
	}
	return res
}
