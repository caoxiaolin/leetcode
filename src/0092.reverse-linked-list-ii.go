package leetcode

/**
 * @title 反转链表 II
 *
 * 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
 *
 * 说明:
 * 1 ≤ m ≤ n ≤ 链表长度。
 *
 * 示例:
 * 输入: 1->2->3->4->5->NULL, m = 2, n = 4
 * 输出: 1->4->3->2->5->NULL
 *
 * @see https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
 */

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	prev := head
	cur := head
	next := head
	start := head
	end := head
	i := 1
	for i < m {
		//i==m-1时，表示从下一个位置开始要反转，所以记录当前位置及下一个位置
		if i == m-1 {
			start = cur
			end = cur.Next
		}
		cur = cur.Next
		i++
	}
	//开始反转
	for i <= n {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		i++
	}
	start.Next = prev
	end.Next = cur
	if m == 1 {
		return prev
	}
	return head
}
