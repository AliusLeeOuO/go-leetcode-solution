package main

import (
	"fmt"
	"solution/solution"
)

func main() {
	fmt.Println("hello world!")
	//values := []*int{
	//	solution.NewInt(3),
	//	solution.NewInt(9),
	//	solution.NewInt(20),
	//	nil,
	//	nil,
	//	solution.NewInt(15),
	//	solution.NewInt(7),
	//}
	//root := solution.CreateTree(values)
	fmt.Printf("\nresult: %v", solution.MakeEqual([]string{"caaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaa", "a", "aaa", "aa", "aa", "aa", "a", "bbbbbbbbbb", "bbbb", "b", "cccccc", "ccc", "c", "cccc"}))
	//fmt.Printf("\nresult: %v", solution.MinOperations([]int{3, 1, 5, 4, 2}, 5))
	//fmt.Printf("\nresult: %v", solution.MinOperations([]int{3, 2, 5, 3, 1}, 3))
}
