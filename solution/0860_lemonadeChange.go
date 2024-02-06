package solution

func LemonadeChange(bills []int) bool {
	cash5Count := 0
	cash10Count := 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			cash5Count++
		} else if bills[i] == 10 {
			if cash5Count == 0 {
				return false
			}
			cash10Count++
			cash5Count--
		} else if bills[i] == 20 {
			if cash5Count == 0 || (cash5Count == 0 && cash10Count == 0) || (cash10Count == 0 && cash5Count < 3) {
				return false
			} else if cash10Count >= 1 && cash5Count >= 1 {
				cash5Count--
				cash10Count--
			} else if cash5Count >= 3 {
				cash5Count -= 3
			}
		}
	}
	return true
}
