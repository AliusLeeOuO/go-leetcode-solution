package solution

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val      int
	Children []*Node
}

// CreateTree 根据层序遍历的节点值列表创建二叉树，nil表示空节点
func CreateTree(values []*int) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}
	index := 1

	for index < len(values) && len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// 添加左子节点
		if index < len(values) && values[index] != nil {
			current.Left = &TreeNode{Val: *values[index]}
			queue = append(queue, current.Left)
		}
		index++

		// 添加右子节点
		if index < len(values) && values[index] != nil {
			current.Right = &TreeNode{Val: *values[index]}
			queue = append(queue, current.Right)
		}
		index++
	}

	return root
}

// NewInt 是一个辅助函数，用于创建 int 类型的指针
func NewInt(val int) *int {
	v := val
	return &v
}

// PrintTree 是一个辅助函数，用于打印二叉树结构
func PrintTree(node *TreeNode, level int, side string) {
	if node == nil {
		return
	}
	format := ""
	for i := 0; i < level; i++ {
		format += "   "
	}
	format += "- " + side + ": "
	fmt.Printf("%s%d\n", format, node.Val)
	PrintTree(node.Left, level+1, "L")
	PrintTree(node.Right, level+1, "R")
}
