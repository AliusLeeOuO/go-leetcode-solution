package solution

func IsCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	// 创建一个队列，用于存储节点以及它们所在的层次
	var queue []struct {
		node   *TreeNode
		level  int
		parent int
	}
	// 将二叉树的根节点以及它所在的层次（1）放入队列中
	queue = append(queue, struct {
		node   *TreeNode
		level  int
		parent int
	}{root, 1, -1})
	// 循环直到队列为空：
	// 分别用于记录x和y的层级和父节点
	xLevel, yLevel, xParent, yParent := -1, -1, -1, -1
	for len(queue) != 0 {
		// 从队列中取出一个节点。
		top := queue[0]
		queue = queue[1:]

		if top.node.Val == x {
			xLevel = top.level
			xParent = top.parent
		} else if top.node.Val == y {
			yLevel = top.level
			yParent = top.parent
		}

		// 如果已经找到x和y，则判断它们是否为堂兄弟节点
		if xLevel != -1 && yLevel != -1 {
			return xLevel == yLevel && xParent != yParent
		}
		// 如果该节点有左子节点，则将右子节点入队。
		if top.node.Left != nil {
			queue = append(queue, struct {
				node   *TreeNode
				level  int
				parent int
			}{node: top.node.Left, level: top.level + 1, parent: top.node.Val})
		}
		// 如果该节点有右子节点，则将右子节点入队。
		if top.node.Right != nil {
			queue = append(queue, struct {
				node   *TreeNode
				level  int
				parent int
			}{node: top.node.Right, level: top.level + 1, parent: top.node.Val})
		}
	}
	return false
}
