/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var head *ListNode
	curr1, curr2 := list1, list2

	if curr1.Val <= curr2.Val {
		head = curr1
		curr1 = curr1.Next
	} else {
		head = curr2
		curr2 = curr2.Next
	}

	tail := head

	for curr1 != nil && curr2 != nil {
		if curr1.Val <= curr2.Val {
			tail.Next = curr1
			curr1 = curr1.Next
		} else {
			tail.Next = curr2
			curr2 = curr2.Next
		}

		tail = tail.Next
	}

	if curr1 != nil && curr2 != nil {
		return head
	}

	if curr1 != nil {
		for curr1 != nil {
			tail.Next = curr1
			curr1 = curr1.Next
			tail = tail.Next
		}
	} else {
		for curr2 != nil {
			tail.Next = curr2
			curr2 = curr2.Next
			tail = tail.Next
		}
	}

	return head
}