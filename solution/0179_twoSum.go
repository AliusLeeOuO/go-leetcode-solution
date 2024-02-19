package solution

func TwoSum(price []int, target int) []int {
	for i := 0; i < len(price); i++ {
		// 折半查找
		searchInts := price[i+1:]
		left := 0
		// 需要查找的数字
		needSearch := target - price[i]
		right := len(searchInts) - 1
		for left <= right {
			mid := left + (right-left)/2
			if needSearch > searchInts[mid] {
				left = mid + 1
			} else if needSearch < searchInts[mid] {
				right = mid - 1
			} else {
				return []int{price[i], needSearch}
			}
		}
	}
	return nil
}
