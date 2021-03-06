package leetcode

/**
 * @title 相交链表
 *
 * 编写一个程序，找到两个单链表相交的起始节点。
 *
 * 例如，下面的两个链表：
 *
 * A:          a1 → a2
 *                    ↘
 *                      c1 → c2 → c3
 *                    ↗
 * B:     b1 → b2 → b3
 *
 * 在节点 c1 开始相交。
 *
 * 注意：
 *     如果两个链表没有交点，返回 null.
 *     在返回结果后，两个链表仍须保持原有的结构。
 *     可假定整个链表结构中没有循环。
 *     程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
 *
 * @see https://leetcode-cn.com/problems/intersection-of-two-linked-lists/description/
 * @difficulty Easy
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getIntersectionNode(l1 *ListNode, l2 *ListNode) *ListNode {
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
	diff := plen - qlen
	if plen < qlen {
		long = l2
		short = l1
		diff = qlen - plen
	}

	//长的链表先向后移动长度差值，这样两个链表就对齐了
	for i := 0; i < diff; i++ {
		long = long.Next
	}

	//两个链表同时向后移动，判断节点是否相等
	for long != nil {
		if long == short {
			return long
		}
		long = long.Next
		short = short.Next
	}
	return nil
}
