func reverseList(head *ListNode) *ListNode {
	return shiftPointers(head, nil)
}

func shiftPointers(curr, prev *ListNode) *ListNode {
	if curr == nil {
		return prev
	}

	next := curr.Next
	curr.Next = prev

	return shiftPointers(next, curr)
}
