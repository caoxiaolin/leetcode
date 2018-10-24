package leetcode

/**
 * @title 移除链表元素
 *
 * 删除链表中等于给定值 val 的所有节点。
 *
 * 示例:
 *
 * 输入: 1->2->6->3->4->5->6, val = 6
 * 输出: 1->2->3->4->5
 *
 * @see https://leetcode-cn.com/problems/remove-linked-list-elements/description/
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	p := &ListNode{Val: 0, Next: head} //构造头节点
	phead := p
	q := head
	for q != nil {
		if q.Val == val {
			p.Next = q.Next //删除节点
		} else {
			p = p.Next //如果不匹配，则p后移一步，否则p不动
		}
		q = q.Next
	}
	return phead.Next
}
