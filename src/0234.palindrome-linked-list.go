package leetcode

/**
 * @title 回文链表
 *
 * 请判断一个链表是否为回文链表。
 *
 * 示例 1:
 * 输入: 1->2
 * 输出: false
 *
 * 示例 2:
 * 输入: 1->2->2->1
 * 输出: true
 *
 * 进阶：
 * 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
 *
 * @see https://leetcode-cn.com/problems/palindrome-linked-list/description/
 * @difficulty Easy
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	p, q := head, head
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next.Next
	}

	//分段
	right := p
	left := head

	//反转
	first := 1
	tmp := right
	for right != nil {
		rn := right.Next
		if first == 1 {
			right.Next = nil
			first = 0
		} else {
			right.Next = tmp
		}
		tmp = right
		right = rn
	}
	right = tmp

	//比较
	for right != nil {
		if right.Val != left.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}
