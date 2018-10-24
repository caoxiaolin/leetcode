package leetcode

/**
 * @title Excel表列名称
 *
 * 给定一个正整数，返回它在 Excel 表中相对应的列名称。
 *
 * 例如，
 *     1 -> A
 *     2 -> B
 *     3 -> C
 *     ...
 *     26 -> Z
 *     27 -> AA
 *     28 -> AB
 *     ...
 *
 * 示例 1:
 * 输入: 1
 * 输出: "A"
 *
 * 示例 2:
 * 输入: 28
 * 输出: "AB"
 *
 * 示例 3:
 * 输入: 701
 * 输出: "ZY"
 *
 * @see https://leetcode-cn.com/problems/excel-sheet-column-title/description/
 *
 * 基本思路：
 * 10进制转26进制
 */

func convertToTitle(n int) string {
	if n == 0 {
		return ""
	}
	n--
	return convertToTitle(n/26) + string(n%26+'A')
}
