package solution

func SortArrayByParityII(nums []int) []int {
	// 判断是否应该是偶数的方法：
	// 		下标%2的结果是0时该位应该是偶数
	for i := 0; i < len(nums); i++ {
		if i%2 != nums[i]%2 {
			// 寻找下一个能跟它交换的数字
			for j := i + 1; j < len(nums); j++ {
				if j%2 != nums[j]%2 && nums[i]%2 != nums[j]%2 {
					// 交换数字
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
	return nums
}
