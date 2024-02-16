package solution

func ZigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var queue = []*TreeNode{root}
	var destination = false
	for len(queue) != 0 {
		var buffer []int
		currentLevel := len(queue)
		for i := 0; i < currentLevel; i++ {
			current := queue[0]
			queue = queue[1:]
			buffer = append(buffer, current.Val)
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		// 如果是反方向，则翻转
		if destination {
			// 使用双指针进行翻转
			// 分别从切片的两端开始，向中心移动，并在过程中交换元素。
			for i, j := 0, len(buffer)-1; i < j; i, j = i+1, j-1 {
				buffer[i], buffer[j] = buffer[j], buffer[i]
			}
		}
		res = append(res, buffer)
		destination = !destination
	}
	return res
}
