package main

import (
	"fmt"
	"solution/solution"
)

func main() {
	fmt.Println("hello world!")
	//values := []*int{solution.NewInt(1), solution.NewInt(2), solution.NewInt(3), nil, solution.NewInt(4)}
	//root := solution.CreateTree(values)

	fmt.Printf("result: %v", solution.FindCenter([][]int{{1, 2}, {2, 3}, {4, 2}}))
}
