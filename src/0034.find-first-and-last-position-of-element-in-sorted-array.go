package leetcode

/**
 * @title 在排序数组中查找元素的第一个和最后一个位置
 *
 * 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
 *
 * 你的算法时间复杂度必须是 O(log n) 级别。
 *
 * 如果数组中不存在目标值，返回 [-1, -1]。
 *
 * 示例 1:
 * 输入: nums = [5,7,7,8,8,10], target = 8
 * 输出: [3,4]
 *
 * 示例 2:
 * 输入: nums = [5,7,7,8,8,10], target = 6
 * 输出: [-1,-1]
 *
 * @see https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 * @difficulty Medium
 *
 * 基本思路：
 * 二分查找
 */
func searchRange(nums []int, target int) []int {
	length := len(nums)
	result := []int{-1, -1}
	if length == 0 || nums[0] > target || nums[length-1] < target {
		return result
	}
	left := 0
	right := length - 1
	for left <= right {
		mid := left + (right-left)/2
		//如果找到了target，则分别向左右查找其起始点
		if nums[mid] == target {
			result[0] = getLeftPos(nums, left, mid, target)
			result[1] = getRightPos(nums, mid, right, target)
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = left + 1
		}
	}
	return result
}

func getLeftPos(nums []int, left, right, target int) int {
	res := 0
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

func getRightPos(nums []int, left, right, target int) int {
	res := 0
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}
