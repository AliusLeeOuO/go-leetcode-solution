package main

import (
	"fmt"
	"solution/solution"
)

func main() {
	fmt.Println("hello world!")
	fmt.Printf("result: %v", solution.CountGoodRectangles([][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}))
}
