package solution

func MinOperations(nums []int, k int) int {
	var res = make(map[int]bool)
	var count = 0
	var ptr = len(nums) - 1
	for len(res) < k {
		if nums[ptr] <= k {
			res[nums[ptr]] = true
		}
		count++
		ptr--
	}
	return count
}
