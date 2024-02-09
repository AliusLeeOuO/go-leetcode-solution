package solution

import (
	"math"
)

func MaximumStrongPairXor(nums []int) int {
	maxXor := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if int(math.Abs(float64(nums[i]-nums[j]))) <= min(nums[i], nums[j]) {
				if nums[i]^nums[j] > maxXor {
					maxXor = nums[i] ^ nums[j]
				}
			}
		}
	}
	return maxXor
}
