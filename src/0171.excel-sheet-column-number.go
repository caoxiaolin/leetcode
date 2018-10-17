package leetcode

/**
 * 给定一个Excel表格中的列名称，返回其相应的列序号。
 *
 * 例如，
 *     A -> 1
 *     B -> 2
 *     C -> 3
 *     ...
 *     Z -> 26
 *     AA -> 27
 *     AB -> 28
 *     ...
 *
 * 示例 1:
 * 输入: "A"
 * 输出: 1
 *
 * 示例 2:
 * 输入: "AB"
 * 输出: 28
 *
 * 示例 3:
 * 输入: "ZY"
 * 输出: 701
 *
 * @see https://leetcode-cn.com/problems/excel-sheet-column-number/description/
 *
 * 基本思路：
 * 26进制转10进制
 */

import (
	"math"
)

func titleToNumber(s string) int {
	res := 0
	len := len(s)
	for k, v := range s {
		res = res + (int(v)-'A'+1)*int(math.Pow(float64(26), float64(len-1-k)))
	}
	return res
}
