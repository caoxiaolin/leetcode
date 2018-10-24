package leetcode

/**
 * @title 两数相加
 *
 * 给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。
 *
 * 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
 *
 * 示例：
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
 *
 * @see https://leetcode-cn.com/problems/add-two-numbers/description/
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p := l1
	q := l2
	jw := 0 //进位
	res := &ListNode{0, nil}
	reshead := res
	for p != nil || q != nil || jw > 0 {
		sum := jw
		if p != nil {
			sum += p.Val
			p = p.Next
		}
		if q != nil {
			sum += q.Val
			q = q.Next
		}
		jw = sum / 10
		res.Next = &ListNode{sum % 10, nil}
		res = res.Next
	}
	return reshead.Next
}
