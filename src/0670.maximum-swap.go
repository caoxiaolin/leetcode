package leetcode

/**
 * @title 最大交换
 *
 * 给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。
 *
 * 示例 1 :
 * 输入: 2736
 * 输出: 7236
 * 解释: 交换数字2和数字7。
 *
 * 示例 2 :
 * 输入: 9973
 * 输出: 9973
 * 解释: 不需要交换。
 *
 * 注意:
 * 给定数字的范围是 [0, 10^8]
 *
 * @see https://leetcode-cn.com/problems/maximum-swap/description/
 * @difficulty Medium
 */

import "strconv"

func maximumSwap(num int) int {
	str1 := strconv.Itoa(num)
	str := str1[:]
	//m,n是两个要交换的下标，x是分割点
	m, n, x := 0, 0, 0
	flag := false //是否需要交换
	len := len(str)
	for i := 0; i < len-1; i++ {
		if str[i] < str[i+1] {
			x = i
			flag = true
			break
		}
	}
	if flag {
		//查找x后面的最大数，记录下标 n
		max := str[x+1]
		n = x + 1
		for i := x + 1; i < len; i++ {
			if str[i] >= max {
				n = i
				max = str[i]
			}
		}
		//从前往后查找x前面比max小的值，记录下标 m
		for i := 0; i <= x; i++ {
			if str[i] < max {
				m = i
				break
			}
		}
	}
	if m > 0 || n > 0 {
		slice := []byte(str)
		slice[m], slice[n] = str[n], str[m]
		str = string(slice)
	}
	res, _ := strconv.Atoi(str)
	return res
}
