package solution

func ClosestNodes(root *TreeNode, queries []int) [][]int {
	var values []int
	// 通过中序遍历得到BST的所有节点值，存储于values数组
	inorderTraversal(root, &values)

	ans := make([][]int, 0, len(queries))
	for _, query := range queries {
		mini, maxi := -1, -1 // 初始化mini和maxi为-1，表示找不到符合条件的值

		// 使用二分搜索在有序数组values中找到最接近query的值的索引
		index := binarySearch(values, query)

		// 根据二分搜索返回的索引，确定mini和maxi
		if index < len(values) {
			if values[index] == query { // 如果找到准确匹配的值
				mini, maxi = values[index], values[index]
			} else {
				maxi = values[index] // 设置maxi为找到的大于等于query的最小值
				if index > 0 {
					mini = values[index-1] // 如果存在，设置mini为找到的小于query的最大值
				}
			}
		} else if index == len(values) { // 如果query大于所有节点值
			mini = values[len(values)-1] // 设置mini为数组中的最后一个值
		}

		// 添加找到的mini和maxi到答案数组
		ans = append(ans, []int{mini, maxi})
	}
	return ans
}

// inorderTraversal 中序遍历二叉搜索树，将节点值以升序方式存储到values切片中
func inorderTraversal(root *TreeNode, values *[]int) {
	if root == nil {
		return // 如果当前节点为空，则返回
	}
	// 递归遍历左子树
	inorderTraversal(root.Left, values)
	// 将当前节点的值添加到values中
	*values = append(*values, root.Val)
	// 递归遍历右子树
	inorderTraversal(root.Right, values)
}

// binarySearch 二分搜索，返回在有序数组values中大于等于query的最小值的索引
func binarySearch(values []int, query int) int {
	left, right := 0, len(values)
	for left < right {
		mid := left + (right-left)/2
		if values[mid] < query { // 如果中间值小于查询值，搜索右半部分
			left = mid + 1
		} else { // 否则，搜索左半部分
			right = mid
		}
	}
	return left // 返回大于等于query的最小值的索引
}
