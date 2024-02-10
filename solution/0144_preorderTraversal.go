package solution

// PreorderTraversal 迭代法实现前序遍历
func PreorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	// 将根节点推入栈中
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) != 0 {
		// 表示当前访问的节点。
		current := stack[len(stack)-1]
		// 访问该节点：按前序遍历的规则，我们首先访问节点本身。
		res = append(res, current.Val)
		// 弹出栈
		stack = stack[:len(stack)-1]
		// 推入右子节点（如果有）
		// 因为栈是后进先出的数据结构，
		// 为了保证左子节点先于右子节点被访问，
		// 我们先将右子节点推入栈中。
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		// 推入左子节点（如果有）
		// 然后将左子节点推入栈中。
		// 这样可以保证下一次循环访问的是左子节点。
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}
	return res
}

//func PreorderTraversal(root *TreeNode) []int {
//	var res []int
//	if root == nil {
//		return res
//	} else {
//		res = append(res, root.Val)
//		res = append(res, PreorderTraversal(root.Left)...)
//		res = append(res, PreorderTraversal(root.Right)...)
//	}
//	return res
//}
