package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeKLists(lists []*ListNode) *ListNode {
	return binaryMerge(lists)
}

func binaryMerge(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	newLists := make([]*ListNode, 0)
	for i := 0; i < len(lists) / 2; i++ {
		newLists = append(newLists, MergeTwoLists(lists[2 * i], lists[2 * i + 1]))
	}
	if len(lists) % 2 == 1 {
		newLists = append(newLists, lists[len(lists) - 1])
	}

	return binaryMerge(newLists)
}