package leetcode

/**
 * @title 对链表进行插入排序
 *
 * 对链表进行插入排序。
 *
 * 插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
 * 每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。
 *
 * 插入排序算法：
 * > 插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
 * > 每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
 * > 重复直到所有输入数据插入完为止。
 *
 * 示例 1：
 * 输入: 4->2->1->3
 * 输出: 1->2->3->4
 *
 * 示例 2：
 * 输入: -1->5->3->4->0
 * 输出: -1->0->3->4->5
 *
 * @see https://leetcode-cn.com/problems/insertion-sort-list/description/
 * @difficulty Medium
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := &ListNode{0, head}
	cur := p.Next
	for cur != nil && cur.Next != nil {
		left := p
		flag := 0
		for left.Next != cur.Next {
			if cur.Next.Val < left.Next.Val {
				tmp := cur.Next //要移动的节点
				cur.Next = cur.Next.Next
				tmp1 := left.Next
				left.Next = tmp
				tmp.Next = tmp1
				flag = 1
				break
			}
			left = left.Next
		}
		if flag == 0 {
			cur = cur.Next
		}
	}
	return p.Next
}
