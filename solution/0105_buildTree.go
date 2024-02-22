package solution

func BuildTree0105(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{
			Val:   preorder[0],
			Left:  nil,
			Right: nil,
		}
	}
	root := preorder[0]
	rootTree := &TreeNode{
		Val:   root,
		Left:  nil,
		Right: nil,
	}
	var i = 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root {
			break
		}
	}
	// build 左子树先序遍历
	leftPreOrder := preorder[1 : i+1]
	rightPreOrder := preorder[i+1:]
	rootTree.Left = BuildTree0105(leftPreOrder, inorder[:i])
	rootTree.Right = BuildTree0105(rightPreOrder, inorder[i+1:])
	return rootTree
}
