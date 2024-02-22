package solution

func BuildTree0106(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	// 后序遍历的最后一个元素为根节点
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}

	// 在中序遍历中找到根节点的索引
	i := 0
	for ; inorder[i] != rootVal; i++ {
	}

	// 切分中序遍历和后序遍历的数组，构建左右子树
	root.Left = BuildTree0106(inorder[:i], postorder[:i])
	root.Right = BuildTree0106(inorder[i+1:], postorder[i:len(postorder)-1])

	return root
}
