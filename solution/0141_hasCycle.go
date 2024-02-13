package solution

func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	// 定义一个快指针，一个慢指针，如果是环，则快指针肯定会追上慢指针
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
