package main

import (
	"fmt"
	"solution/solution"
)

func main() {
	fmt.Println("hello world!")
	values := []*int{
		solution.NewInt(5),
		solution.NewInt(8),
		solution.NewInt(9),
		solution.NewInt(2),
		solution.NewInt(1),
		solution.NewInt(3),
		solution.NewInt(7),
	}
	root := solution.CreateTree(values)
	var a = solution.KthLargestLevelSum(root, 4)
	fmt.Printf("\nresult: %v", a)
	//fmt.Printf("\nresult: %v", solution.MinOperations([]int{3, 2, 5, 3, 1}, 3))
}
