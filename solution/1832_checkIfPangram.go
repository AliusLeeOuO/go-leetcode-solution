package solution

func CheckIfPangram(sentence string) bool {
	mapping := make(map[string]bool)
	for i := 0; i < len(sentence); i++ {
		char := string(sentence[i])
		if !mapping[char] {
			mapping[char] = true
		}
	}
	count := 0
	for i := 0; i < len(mapping); i++ {
		count++
	}
	return count == 26
}
