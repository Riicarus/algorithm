package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}

	tail := res

	n1 := list1
	n2 := list2
	for ; n1 != nil && n2 != nil; tail = tail.Next {
		if n1.Val <= n2.Val {
			addTail(*n1, tail)
			n1 = n1.Next
		} else {
			addTail(*n2, tail)
			n2 = n2.Next
		}
	}

	if n1 == nil {
		tail.Next = n2
	}

	if n2 == nil {
		tail.Next = n1
	}

	return res.Next
}

func addTail(v ListNode, tail *ListNode) {
	tail.Next = &ListNode{
		Val: v.Val,
		Next: nil,
	}
}