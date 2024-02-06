package solution

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	// 将字符串去杂
	// 定义正则规则
	onlyAlphaRegx := regexp.MustCompile("[^a-zA-Z0-9]")
	b := onlyAlphaRegx.ReplaceAllString(s, "")
	onlyAlphaStr := strings.ToLower(b)
	// 去杂后的字符串进行计算长度
	alphaStringLength := len(onlyAlphaStr)
	if alphaStringLength == 0 {
		return true
	}
	// 设置前指针和后指针，用于对比
	frontPtr := 0
	backPtr := alphaStringLength - 1
	for i := 0; i < alphaStringLength/2; i++ {
		if onlyAlphaStr[frontPtr] != onlyAlphaStr[backPtr] {
			return false
		} else {
			frontPtr++
			backPtr--
		}
	}
	return true
}
