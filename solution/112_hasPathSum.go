package solution

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 深度优先遍历
	/*
	 * 1. 访问根节点，记录根节点数字 = 5
	 * 2. 访问其左子节点，计算总和 5+4=9 < 22，继续找左子树
	 * 3. 访问其左子节点，计算总和 5+4+11=20 < 22，继续找左子树
	 * 4. 访问其左子节点，计算总和 5+4+11+7=27 > 22，不满足条件，往回退一层
	 * 4. 访问其右子节点，计算总和 5+4+11+2=22 = 22，满足条件，结束
	 */
	// 如果当前节点是叶子节点，并且节点的值等于剩余的 targetSum，则找到了路径
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}
	// 递归地检查左子树和右子树，更新 targetSum 为 targetSum 减去当前节点的值
	return HasPathSum(root.Left, targetSum-root.Val) || HasPathSum(root.Right, targetSum-root.Val)
}
