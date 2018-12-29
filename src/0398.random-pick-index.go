package leetcode

/**
 * @title 随机数索引
 *
 * 给定一个可能含有重复元素的整数数组，要求随机输出给定的数字的索引。 您可以假设给定的数字一定存在于数组中。
 *
 * 注意：
 * 数组大小可能非常大。 使用太多额外空间的解决方案将不会通过测试。
 *
 * 示例:
 *
 * int[] nums = new int[] {1,2,3,3,3};
 * Solution solution = new Solution(nums);
 *
 * // pick(3) 应该返回索引 2,3 或者 4。每个索引的返回概率应该相等。
 * solution.pick(3);
 *
 * // pick(1) 应该返回 0。因为只有nums[0]等于1。
 * solution.pick(1);
 *
 * @see https://leetcode-cn.com/problems/random-pick-index/
 * @difficulty Medium
 */

import "math/rand"

type Solution0398 struct {
	Nums []int
}

func Constructor0398(nums []int) Solution0398 {
	return Solution0398{nums}
}

func (this *Solution0398) Pick(target int) int {
	res := -1
	n := 0
	for k, v := range this.Nums {
		if v == target {
			n++
			if n == 1 {
				res = k
			} else {
				r := rand.Float64()
				if r < 1/float64(n) {
					res = k
				}
			}
		}
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
