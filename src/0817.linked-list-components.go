package leetcode

/**
 * 给定一个链表（链表结点包含一个整型值）的头结点 head。
 *
 * 同时给定列表 G，该列表是上述链表中整型值的一个子集。
 *
 * 返回列表 G 中组件的个数，这里对组件的定义为：链表中一段最长连续结点的值（该值必须在列表 G 中）构成的集合。
 *
 * 示例 1：
 * 输入:
 * head: 0->1->2->3
 * G = [0, 1, 3]
 * 输出: 2
 * 解释:
 * 链表中,0 和 1 是相连接的，且 G 中不包含 2，所以 [0, 1] 是 G 的一个组件，同理 [3] 也是一个组件，故返回 2。
 *
 * 示例 2：
 * 输入:
 * head: 0->1->2->3->4
 * G = [0, 3, 1, 4]
 * 输出: 2
 * 解释:
 * 链表中，0 和 1 是相连接的，3 和 4 是相连接的，所以 [0, 1] 和 [3, 4] 是两个组件，故返回 2。
 *
 * @see https://leetcode-cn.com/problems/linked-list-components/description/
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, G []int) int {
	m := make(map[int]bool)
	for _, v := range G {
		m[v] = true
	}
	p := head
	res := 0
	flag := 0
	for p != nil {
		if m[p.Val] == true {
			if flag == 0 {
				res++
				flag = 1
			}
		} else {
			flag = 0
		}
		p = p.Next
	}
	return res
}
