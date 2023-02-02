package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := make([]*ListNode, 0)

	var temp *ListNode
	for temp = head; temp != nil; temp = temp.Next {
		nodes = append(nodes, temp)
	}

	nth := nodes[len(nodes) - n]

	// 如果删除第一个元素，就直接返回其下一个元素作为 head
	if len(nodes) - n == 0 {
		return nth.Next
	}

	pre := nodes[len(nodes) - n - 1]
	pre.Next = nth.Next

	return head
}
