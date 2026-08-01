func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nodes := []*ListNode{}

	currentNode := head

	for currentNode != nil {
		nodes = append(nodes, currentNode)
		currentNode = currentNode.Next
	}

	head = nodes[len(nodes)-1]
	currentNode = head

	for i := len(nodes) - 2; i >= 0; i-- {
		currentNode.Next = nodes[i]
		currentNode = currentNode.Next
	}

	currentNode.Next = nil

	return head
}

