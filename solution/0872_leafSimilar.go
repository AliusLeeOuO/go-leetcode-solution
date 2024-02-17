package solution

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return false
	}
	var stack = []*TreeNode{root1}
	var node1Leaf []int
	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.Left == nil && current.Right == nil {
			node1Leaf = append(node1Leaf, current.Val)
		} else {
			if current.Right != nil {
				stack = append(stack, current.Right)
			}
			if current.Left != nil {
				stack = append(stack, current.Left)
			}
		}
	}
	stack = []*TreeNode{root2}
	// 比较箭头
	pkPtr := 0
	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.Left == nil && current.Right == nil {
			// 如果下标越界，直接返回false
			if pkPtr == len(node1Leaf) || current.Val != node1Leaf[pkPtr] {
				return false
			}
			pkPtr++
		} else {
			if current.Right != nil {
				stack = append(stack, current.Right)
			}
			if current.Left != nil {
				stack = append(stack, current.Left)
			}
		}
	}
	// 如果长度不符，返回false
	if pkPtr != len(node1Leaf) {
		return false
	}
	return true
}
