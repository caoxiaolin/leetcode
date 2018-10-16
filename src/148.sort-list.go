package leetcode

/**
 * 在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
 *
 * 示例 1:
 *
 * 输入: 4->2->1->3
 * 输出: 1->2->3->4
 * 示例 2:
 *
 * 输入: -1->5->3->4->0
 * 输出: -1->0->3->4->5
 *
 * @see https://leetcode-cn.com/problems/sort-list/description/
 *
 * 基本思路：
 * O(n log n)的时间复杂度，归并排序
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	q := head.Next
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next.Next
	}
	l := head
	r := p.Next
	p.Next = nil
	return merge(sortList(l), sortList(r))
}

//合并两个有序链表（pls see 021.merge-two-sorted-lists）
func merge(l, r *ListNode) *ListNode {
	res := new(ListNode) //初始化一个新链表
	tmp := res
	for l != nil && r != nil {
		if l.Val < r.Val {
			tmp.Next = l
			l = l.Next
			tmp = tmp.Next
		} else {
			tmp.Next = r
			r = r.Next
			tmp = tmp.Next
		}
	}
	if l != nil {
		tmp.Next = l
	}
	if r != nil {
		tmp.Next = r
	}
	return res.Next
}
