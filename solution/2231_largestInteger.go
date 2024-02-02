package solution

import (
	"fmt"
	"strconv"
)

// 交换两个整数的值
func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}

// 分区函数
func partition(arr []int, low int, high int) int {
	pivot := arr[high] // 选择最后一个元素作为枢纽
	i := low - 1       // i 是比枢纽小的元素的索引
	// 遍历数组，重新排列元素
	for j := low; j < high; j++ {
		if arr[j] > pivot {
			i++
			swap(&arr[i], &arr[j])
		}
	}
	swap(&arr[i+1], &arr[high]) // 将枢纽放到它最终的位置
	return i + 1                // 返回枢纽的索引
}

// 快速排序函数
func quickSort(arr []int, low int, high int) {
	if low < high {
		// 找到枢纽元素所在位置
		pi := partition(arr, low, high)

		// 递归地对枢纽左侧的元素进行快速排序
		quickSort(arr, low, pi-1)
		// 递归地对枢纽右侧的元素进行快速排序
		quickSort(arr, pi+1, high)
	}
}

func LargestInteger(num int) int {
	// 将int转换为字符串
	numStr := strconv.Itoa(num)
	// 使用一个数组，记录原数组的奇偶性 1是奇数 0是偶数
	var newNum []int
	// 偶数
	var evens []int
	// 奇数
	var odds []int
	for i := 0; i < len(numStr); i++ {
		char := int(numStr[i]) - '0'
		if char%2 == 0 {
			evens = append(evens, char)
		} else if char%2 == 1 {
			odds = append(odds, char)
		}
		newNum = append(newNum, char%2)
	}
	// 降序排列
	quickSort(evens, 0, len(evens)-1)
	quickSort(odds, 0, len(odds)-1)
	// 重新排列
	evensPtr := 0
	oddsPtr := 0
	newNumString := ""
	for i := 0; i < len(newNum); i++ {
		// 如果是奇数
		if newNum[i] == 1 {
			newNum[i] = odds[oddsPtr]
			oddsPtr++
		} else if newNum[i] == 0 {
			// 如果是偶数
			newNum[i] = evens[evensPtr]
			evensPtr++
		}
		newNumString += fmt.Sprintf("%v", newNum[i])
	}
	numResult, _ := strconv.ParseInt(newNumString, 10, 32)
	return int(numResult)
}
