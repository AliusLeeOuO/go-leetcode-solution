package solution

import "fmt"

func HammingWeight(num uint32) int {
	c := 0
	// 这种方法通过不断地将数字右移并检查最低位的值来计算位1的个数。
	// 每次迭代中，通过与1进行AND操作（num & 1）来检查最低位是否为1，
	// 如果是，计数器增加。然后将数字右移一位（num >>= 1），直到数字变为0。
	//
	// AND操作是一种基本的位运算，用于处理二进制数中的对应位。
	// 在AND操作中，只有当两个操作数的对应位都为1时，结果的那一位才为1；
	// 否则，结果的那一位为0。这可以用下面的真值表来表示：
	//
	// A (操作数1)		B (操作数2)		A AND B (结果)
	// 	  0					0				0
	//    0					1				0
	//    1					0				0
	//    1					1				1
	// 这里，A和B是操作数的位，A AND B是执行AND操作后的结果位。
	//
	// 例如，考虑以下两个二进制数进行AND操作：
	//
	//   10101010
	// & 11001100
	// ----------
	//   10001000
	// 在这个例子中，只有第1位和第7位在两个操作数中都是1，因此在结果中这两位也是1。
	// 其余位在至少一个操作数中为0，所以在结果中这些位都是0。
	//
	// AND操作在计算机科学中有多种用途，包括：
	//
	// 清零操作：通过与一个特定的掩码进行AND操作，可以将数字的某些位清零。
	// 位掩码：检查数字的特定位是否为1。
	// 权限控制：在权限设置中，通过AND操作检查是否具有某些权限。
	// AND操作是位运算的基础之一，
	// 与OR（或运算）、XOR（异或运算）、NOT（非运算）一起构成了位运算的基本集合。
	for num != 0 {
		fmt.Println(num)
		if num&1 == 1 {
			c++
		}
		num >>= 1
	}
	return c
}