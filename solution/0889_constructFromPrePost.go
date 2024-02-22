package solution

func ConstructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}
	// 确定左子树的边界
	i := 0
	for ; i < len(postorder); i++ {
		if postorder[i] == preorder[1] {
			break
		}
	}
	// 左子树后续遍历
	treeLeftPostOrder := postorder[:i+1]
	//左子树前序遍历
	treeLeftPreOrder := preorder[1 : len(treeLeftPostOrder)+1]
	// 右子树后续遍历
	treeRightPostOrder := postorder[i+1 : len(postorder)-1]
	// 右子树前序遍历
	treeRightPreOrder := preorder[len(treeLeftPostOrder)+1:]
	root.Left = ConstructFromPrePost(treeLeftPreOrder, treeLeftPostOrder)
	root.Right = ConstructFromPrePost(treeRightPreOrder, treeRightPostOrder)
	return root
}
