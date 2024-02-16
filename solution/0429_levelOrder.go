package solution

func levelOrder0429(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var queue = []*Node{root}
	var nowLevelSize = 1
	var nextLevelSize = 0
	for len(queue) != 0 {
		var buffer []int
		for i := 0; i < nowLevelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			buffer = append(buffer, current.Val)
			if current.Children != nil {
				nextLevelSize += len(current.Children)
				// 将子节点压入栈
				for j := 0; j < len(current.Children); j++ {
					queue = append(queue, current.Children[j])
				}
			}
		}
		res = append(res, buffer)
		nowLevelSize = nextLevelSize
		nextLevelSize = 0
	}
	return res
}
