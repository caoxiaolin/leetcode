package leetcode

/**
 * Init singly-linked list
 */
func InitListNode(list []int) *ListNode {
	nodehead := &ListNode{Val: 0, Next: nil}
	listnode := nodehead
	for _, v := range list {
		listnode.Next = &ListNode{Val: v, Next: nil}
		listnode = listnode.Next
	}
	return nodehead.Next
}

/**
 * ListNode to Slice
 */
func ListNode2Slice(list *ListNode) []int {
	res := []int{}
	for list != nil {
		res = append(res, list.Val)
		list = list.Next
	}
	return res
}

/**
 * Equal two slice
 */
func SliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
