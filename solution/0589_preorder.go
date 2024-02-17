package solution

/*
6
5没有小孩
3(1children[0])
1 - 3 2 4
*/

func Preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	// 初始化一个栈，首先只包含根节点。
	var stack = []*Node{root}
	// 重复这个过程，直到栈为空。
	for len(stack) != 0 {
		// 在每次迭代中，从栈中弹出一个节点，记录它的值。
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, current.Val)
		if current.Children != nil {
			// 将该节点的所有子节点从右到左压入栈中。
			// 这样做是为了保证下一次弹出栈时
			// 可以按照前序遍历的顺序（根节点->左子节点->右子节点）访问这些子节点。
			for i := len(current.Children) - 1; i >= 0; i-- {
				stack = append(stack, current.Children[i])
			}
		}
	}
	return res
}
