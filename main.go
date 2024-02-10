package main

import (
	"fmt"
	"solution/solution"
)

func main() {
	fmt.Println("hello world!")
	values := []*int{
		solution.NewInt(3),
		solution.NewInt(5),
		solution.NewInt(1),
		solution.NewInt(6),
		solution.NewInt(2),
		solution.NewInt(0),
		solution.NewInt(8),
		nil,
		nil,
		solution.NewInt(7),
		solution.NewInt(4),
	}
	root := solution.CreateTree(values)

	fmt.Printf("\nresult: %v", solution.PreorderTraversal(root))
	//fmt.Printf("\nresult: %v", solution.MinOperations([]int{3, 1, 5, 4, 2}, 5))
	//fmt.Printf("\nresult: %v", solution.MinOperations([]int{3, 2, 5, 3, 1}, 3))
}
