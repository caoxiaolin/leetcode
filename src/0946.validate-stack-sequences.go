package leetcode

/**
 * @title 验证栈序列
 *
 * 给定 pushed 和 popped 两个序列，只有当它们可能是在最初空栈上进行的推入 push 和弹出 pop 操作序列的结果时，返回 true；否则，返回 false 。
 *
 * 示例 1：
 * 输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
 * 输出：true
 * 解释：我们可以按以下顺序执行：
 * push(1), push(2), push(3), push(4), pop() -> 4,
 * push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
 *
 * 示例 2：
 * 输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
 * 输出：false
 * 解释：1 不能在 2 之前弹出。
 *
 * @see https://leetcode-cn.com/problems/validate-stack-sequences/
 * @difficulty Medium
 */

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	i := 0
	for _, v := range pushed {
		stack = append(stack, v)
		length := len(stack)
		for length > 0 && stack[length-1] == popped[i] {
			stack = stack[:length-1]
			length--
			i++
		}
	}
	return len(stack) == 0
}
