package solution

func KthToLast(head *ListNode, k int) int {
	current := head
	back := current
	for i := 0; i < k; i++ {
		current = current.Next
	}
	for current != nil {
		back = back.Next
		current = current.Next
	}
	return back.Val
}
