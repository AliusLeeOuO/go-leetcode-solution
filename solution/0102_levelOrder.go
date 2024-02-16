package solution

func LevelOrder0102(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// 结果集
	var res = make([][]int, 0)
	// 创建一个队列，用于遍历，将root入队
	var queue = []*TreeNode{root}
	// 循环出队
	for len(queue) > 0 {
		var buffer []int
		// 当前层的节点数
		var levelSize = len(queue)
		for i := 0; i < levelSize; i++ {
			// 出队操作
			current := queue[0]
			queue = queue[1:]
			buffer = append(buffer, current.Val)
			// 将子节点入队
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		res = append(res, buffer)
	}
	return res
}
