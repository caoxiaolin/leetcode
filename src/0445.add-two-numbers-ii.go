package leetcode

/**
 * @title 两数相加 II
 *
 * 给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。
 *
 * 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
 *
 * 进阶:
 * 如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
 *
 * 示例:
 * 输入: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出: 7 -> 8 -> 0 -> 7
 *
 * @see https://leetcode-cn.com/problems/add-two-numbers-ii/description/
 */

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	p := l1
	q := l2
	plen, qlen := 0, 0

	//先分别计算两个链表的长度
	for p != nil {
		p = p.Next
		plen++
	}

	for q != nil {
		q = q.Next
		qlen++
	}

	long := l1
	short := l2
	longlen := plen
	diff := plen - qlen
	if plen < qlen {
		long = l2
		short = l1
		longlen = qlen
		diff = qlen - plen
	}

	res := &ListNode{0, nil}
	reshead := res

	//长的链表先向后移动长度差值，这样两个链表就对齐了
	for i := 0; i < diff; i++ {
		res.Next = &ListNode{long.Val, nil}
		long = long.Next
		res = res.Next
	}

	//两个链表同时向后移动，val相加
	for long != nil {
		res.Next = &ListNode{long.Val + short.Val, nil}
		long = long.Next
		short = short.Next
		res = res.Next
	}

	result := reshead.Next
	//遍历新的链表，处理进位问题
	fmt.Println(longlen)
	for longlen > 0 {
		qq := result
		for qq.Next != nil {
			if qq.Next.Val > 9 {
				qq.Val += 1
				qq.Next.Val -= 10
			}
			qq = qq.Next
		}
		longlen--
	}

	//处理第一位的进位
	if result.Val >= 10 {
		node := &ListNode{result.Val / 10, result}
		result.Val %= 10
		return node
	}
	return result
}
