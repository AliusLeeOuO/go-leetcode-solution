package solution

func PostorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	var stack2 []*TreeNode
	// 处理根节点：将根节点推入第一个栈中。
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) != 0 {
		// 弹出栈顶元素：从第一个栈中弹出当前访问的节点。
		current := stack[len(stack)-1]
		// 正确的位置，在处理子节点之前，将当前节点从第一个栈中移除
		stack = stack[:len(stack)-1]
		//  推入第二个栈：将弹出的节点推入第二个栈中。第二个栈用于逆序存储访问的节点，以满足后序遍历的顺序。
		stack2 = append(stack2, current)
		// 推入子节点到第一个栈：
		//      首先将当前节点的左子节点（如果有）推入第一个栈，然后推入右子节点（如果有）。
		//      这样做是因为我们希望最后访问根节点，而栈的特性是后进先出，所以要先处理左节点再处理右节点，这与前序遍历的顺序相反。
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
	}
	for len(stack2) != 0 {
		current := stack2[len(stack2)-1]
		res = append(res, current.Val)
		stack2 = stack2[:len(stack2)-1]
	}
	return res
}

/*
func postorderTraversal(root *TreeNode) []int {
    var res []int
    if root == nil {
        return res
    } else {
        res = append(postorderTraversal(root.Left), postorderTraversal(root.Right)...)
        res = append(res, root.Val)
    }
    return res
}
*/
