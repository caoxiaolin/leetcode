package leetcode

/**
 * @title 删除排序链表中的重复元素 II
 *
 * 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
 *
 * 示例 1:
 *
 * 输入: 1->2->3->3->4->4->5
 * 输出: 1->2->5
 * 示例 2:
 *
 * 输入: 1->1->1->2->3
 * 输出: 2->3
 *
 * @see https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * 基本思路：
 * 双指针
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := &ListNode{Val: 0, Next: head}
	phead := p
	q := head
	flag := 0 //标志位，是否重复
	for q.Next != nil {
		if q.Val == q.Next.Val {
			flag = 1
		} else if flag == 0 {
			p = p.Next
		} else if flag == 1 {
			p.Next = q.Next
			flag = 0
		}
		q = q.Next
	}
	if flag == 1 {
		p.Next = q.Next
	}
	return phead.Next
}
