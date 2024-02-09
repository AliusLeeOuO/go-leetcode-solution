package solution

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	// 在左子树中查找p和q，返回左子树调用的结果。
	left := LowestCommonAncestor(root.Left, p, q)
	// 在右子树中查找p和q，返回右子树调用的结果。
	right := LowestCommonAncestor(root.Right, p, q)

	// 判断当前节点是否为LCA
	if left != nil && right != nil {
		// 如果在左右子树的搜索结果中，一个节点来自左子树而另一个来自右子树，
		// 这意味着当前节点就是p和q的LCA，返回当前节点。
		return root
	}
	if left == nil {
		// 如果左子树搜索结果为空，这意味着p和q都不在左子树中，返回右子树的搜索结果。
		return right
	}
	// 如果右子树搜索结果为空，这意味着p和q都不在右子树中，返回左子树的搜索结果。
	return left
}
