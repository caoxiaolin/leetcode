package leetcode

/**
 * @title 螺旋矩阵 II
 *
 * 给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
 *
 * 示例:
 *
 * 输入: 3
 * 输出:
 * [
 *  [ 1, 2, 3 ],
 *  [ 8, 9, 4 ],
 *  [ 7, 6, 5 ]
 * ]
 *
 * @see https://leetcode-cn.com/problems/spiral-matrix-ii/description/
 */
//import "fmt"
func generateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}
	result := make([][]int, n, n)
	for i := 0; i < n; i++ {
		s := make([]int, n)
		result[i] = s
	}
	left, right, top, bottom := 0, n-1, 0, n-1
	step := 1
	sum := 1
	for right > 0 || bottom > 0 {
		if step == 1 {
			for i := left; i <= right; i++ {
				result[top][i] = sum
				sum++
			}
			top++
			step = 2
		}
		if step == 2 {
			for i := top; i <= bottom; i++ {
				result[i][right] = sum
				sum++
			}
			right--
			step = 3
		}
		if step == 3 {
			for i := right; i >= left; i-- {
				result[bottom][i] = sum
				sum++
			}
			bottom--
			step = 4
		}
		if step == 4 {
			for i := bottom; i >= top; i-- {
				result[i][left] = sum
				sum++
			}
			left++
			step = 1
		}
	}
	return result
}
