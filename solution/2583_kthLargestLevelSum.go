package solution

import "sort"

func KthLargestLevelSum(root *TreeNode, k int) int64 {
	var c []int
	// 层序遍历
	// 创建一个队列，用于遍历，将root入队
	var queue = []*TreeNode{root}
	for len(queue) != 0 {
		sum := 0
		var levelSize = len(queue)
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			sum += current.Val
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		c = append(c, sum)
	}
	sort.Ints(c)
	if len(c)-k < 0 {
		return -1
	}
	return int64(c[len(c)-k])
}
