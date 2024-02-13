package solution

import "strings"

func MergeAlternately(word1 string, word2 string) string {
	var res strings.Builder
	word1Ptr := 0
	word2Ptr := 0
	for len(word1) > word1Ptr && len(word2) > word2Ptr {
		res.WriteByte(word1[word1Ptr])
		word1Ptr++
		res.WriteByte(word2[word2Ptr])
		word2Ptr++
	}
	if word1Ptr < len(word1) {
		res.WriteString(word1[word1Ptr:])
	}
	if word2Ptr < len(word2) {
		res.WriteString(word2[word2Ptr:])
	}
	return res.String()
}
