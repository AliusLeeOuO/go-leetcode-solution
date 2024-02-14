package solution

func LevelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		var currentSize = len(queue)
		var buffer []int
		for i := 0; i < currentSize; i++ {
			// 将节点出队，并加入缓冲区
			current := queue[0]
			queue = queue[1:]
			buffer = append(buffer, current.Val)
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		// 将缓冲区加入结果集，从头部添加符合题目要求
		res = append([][]int{buffer}, res...)
	}
	return res
}
