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
	//	solution.NewInt(7),
	//	solution.NewInt(4),
	//	solution.NewInt(2),
	//	nil,
	//	nil,
	//	nil,
	//	nil,
	//	nil,
	//	nil,
	//	solution.NewInt(9),
	//	solution.NewInt(11),
	//	nil,
	//	nil,
	//	solution.NewInt(8),
	//	solution.NewInt(10),
	//}
	//root := solution.CreateTree(values)

	fmt.Printf("\nresult: %v", solution.BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
	//fmt.Printf("\nresult: %v", solution.MinOperations([]int{3, 2, 5, 3, 1}, 3))
}
