package solution

// InorderTraversal 迭代法二叉树中序遍历
func InorderTraversal(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	// 当前节点
	var current = root
	for current != nil || len(stack) != 0 {
		// 将所有左子节点压入栈中
		for current != nil {
			// 压入栈中
			stack = append(stack, current)
			// 转向左子节点
			current = current.Left
		}
		// 当前节点为空，意味着已到达最左侧，弹出栈顶元素
		current = stack[len(stack)-1]
		// 将栈顶元素的值加入结果集
		result = append(result, current.Val)
		// 弹出栈顶元素
		stack = stack[:len(stack)-1]
		// 转向右子节点
		current = current.Right
	}
	return result
}

//func PreorderTraversal(root *TreeNode) []int {
//	var res []int
//	if root == nil {
//		return res
//	} else {
//		res = append(res, PreorderTraversal(root.Left)...)
//		res = append(res, root.Val)
//		res = append(res, PreorderTraversal(root.Right)...)
//	}
//	return res
//}
