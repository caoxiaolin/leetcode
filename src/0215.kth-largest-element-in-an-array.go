package leetcode

/**
 * @title 数组中的第K个最大元素
 *
 * 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 *
 * 示例 1:
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 *
 * 示例 2:
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 *
 * 说明:
 * 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
 *
 * @see https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/
 * @difficulty Medium
 *
 * 基本思路：
 * 堆排序
 */

func findKthLargest(nums []int, k int) int {
	//建立大顶堆
	length := len(nums)
	for i := length/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, length)
	}
	n := 0
	res := 0
	for i := length - 1; i >= 0; i-- {
		n++
		if n == k {
			res = nums[0]
			break
		}
		nums[0], nums[i] = nums[i], nums[0]
		adjustHeap(nums, 0, i)
	}
	return res
}

func adjustHeap(nums []int, i, l int) {
	tmp := nums[i]
	for j := 2 * (i + 1); j <= l; j = 2 * j {
		if j < l && nums[j-1] < nums[j] {
			j++
		}
		if tmp > nums[j-1] {
			break
		}
		nums[i] = nums[j-1]
		i = j - 1
	}
	nums[i] = tmp
}
