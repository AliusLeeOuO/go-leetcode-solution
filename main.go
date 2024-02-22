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
	var a = solution.ConstructFromPrePost([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1})
	fmt.Printf("\nresult: %v", a)
	//fmt.Printf("\nresult: %v", solution.MinOperations([]int{3, 2, 5, 3, 1}, 3))
}
