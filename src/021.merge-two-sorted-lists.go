package leetcode

/**
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

@see https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p := l1
	q := l2
	res := new(ListNode)
	tmp := res
	for p != nil && q != nil {
		if p.Val < q.Val {
			tmp.Next = p
			p = p.Next
			tmp = tmp.Next
		} else {
			tmp.Next = q
			q = q.Next
			tmp = tmp.Next
		}
	}
	if p != nil {
		tmp.Next = p
	}
	if q != nil {
		tmp.Next = q
	}
	return res.Next
}
