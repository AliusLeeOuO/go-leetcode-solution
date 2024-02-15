package solution

// TrainningPlan 合并有序链表
func TrainningPlan(l1 *ListNode, l2 *ListNode) *ListNode {
	// 边界处理
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	// 新建一个哨兵节点作为结果链表的起始节点，简化边界条件处理
	sentinel := &ListNode{Val: 0}
	tail := sentinel
	// 遍历l1和l2直到一个到达末尾
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else if l1.Val > l2.Val {
			tail.Next = l2
			l2 = l2.Next
		}
		// 将链表尾指针指向尾部节点
		tail = tail.Next
	}
	if l1 != nil {
		tail.Next = l1
	} else if l2 != nil {
		tail.Next = l2
	}
	return sentinel.Next
}
