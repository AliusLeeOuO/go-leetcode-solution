package solution

func Postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	stack := []*Node{root}
	var stack2 []*Node
	var res []int
	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		stack2 = append(stack2, current)
		if current.Children != nil {
			for i := 0; i < len(current.Children); i++ {
				stack = append(stack, current.Children[i])
			}
		}
	}
	for len(stack2) != 0 {
		res = append(res, stack2[len(stack2)-1].Val)
		stack2 = stack2[:len(stack2)-1]
	}
	return res
}
