package leetcode

/**
 * @title 删除排序链表中的重复元素
 *
 * 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
 *
 * 示例 1:
 * 输入: 1->1->2
 * 输出: 1->2
 *
 * 示例 2:
 * 输入: 1->1->2->3->3
 * 输出: 1->2->3
 *
 * @see https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/description/
 * @difficulty Easy
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	for p != nil && p.Next != nil {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next //删除节点
		} else {
			p = p.Next //值不相等的时候，p后移一位
		}
	}
	return head
}
