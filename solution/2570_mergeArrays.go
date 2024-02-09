package solution

import (
	"sort"
)

func MergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	numsHashList := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		numsHashList[nums1[i][0]] += nums1[i][1]
	}
	for i := 0; i < len(nums2); i++ {
		numsHashList[nums2[i][0]] += nums2[i][1]
	}
	// 提取map的所有键到切片keys
	var numListSort []int
	for i, _ := range numsHashList {
		numListSort = append(numListSort, i)
	}
	sort.Ints(numListSort)
	var numsList [][]int
	for _, k := range numListSort {
		numsList = append(numsList, []int{k, numsHashList[k]})
	}
	return numsList
}
