package main

import (
	"fmt"
	"solution/solution"
)

func main() {
	fmt.Println("hello world!")
	// 使用指针数组以便能够表示 nil 值
	binaryTree := []*int{solution.NewInt(5), solution.NewInt(4), solution.NewInt(8), solution.NewInt(11), nil, solution.NewInt(13), solution.NewInt(4), solution.NewInt(7), solution.NewInt(2), nil, nil, nil, solution.NewInt(1)}
	tree := solution.CreateTree(binaryTree)
	solution.PrintTree(tree, 0, "root")
	fmt.Printf("result: %v", solution.HasPathSum(tree, 22))
}
