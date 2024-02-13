package solution

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// 结果集
	var res = make([][]int, 0)
	// 创建一个队列，用于遍历
	var stack []struct {
		value *TreeNode
		level int
	}
	// 将root入队
	stack = append(stack, struct {
		value *TreeNode
		level int
	}{value: root, level: 1})
	// 每层的标记
	resLevel := 1
	var resHorizon []int
	// 循环出队
	for len(stack) != 0 {
		current := stack[0]
		stack = stack[1:]
		// 添加到结果集，如果层次达不到
		if resLevel != current.level {
			res = append(res, resHorizon)
			resHorizon = make([]int, 0)
			resLevel = current.level
		}
		resHorizon = append(resHorizon, current.value.Val)
		// 如果有左子树，则添加到对队列中
		if current.value.Left != nil {
			stack = append(stack, struct {
				value *TreeNode
				level int
			}{value: current.value.Left, level: current.level + 1})
		}
		// 如果有右子树，则添加到对队列中
		if current.value.Right != nil {
			stack = append(stack, struct {
				value *TreeNode
				level int
			}{value: current.value.Right, level: current.level + 1})
		}
	}
	// 将最后一行加入结果集
	res = append(res, resHorizon)
	return res
}
