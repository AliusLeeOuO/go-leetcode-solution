package main

import (
	"fmt"
	"solution/solution"
)

func main() {
	fmt.Println("hello world!")
	//values := []*int{
	//	solution.NewInt(3),
	//	solution.NewInt(5),
	//	solution.NewInt(1),
	//	solution.NewInt(6),
	//	solution.NewInt(2),
	//	solution.NewInt(0),
	//	solution.NewInt(8),
	//	nil,
	//	nil,
	//	solution.NewInt(7),
	//	solution.NewInt(4),
	//}
	//root := solution.CreateTree(values)

	fmt.Printf("result: %v", solution.MergeArrays([][]int{{1, 2}, {2, 3}, {4, 5}}, [][]int{{1, 4}, {3, 2}, {4, 1}}))
}
