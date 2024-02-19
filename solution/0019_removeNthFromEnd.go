package solution

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	current := head
	slow := head
	var slowPrev = &ListNode{
		Val:  0,
		Next: head,
	}
	h := slowPrev
	for i := 0; i < n; i++ {
		current = current.Next
	}
	for current != nil {
		current = current.Next
		slow = slow.Next
		slowPrev = slowPrev.Next
	}
	slowPrev.Next = slow.Next
	slow = nil
	return h.Next
}
