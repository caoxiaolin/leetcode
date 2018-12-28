package leetcode

/**
 * @title 简化路劲
 *
 * 给定一个文档 (Unix-style) 的完全路径，请进行路径简化。
 *
 * 例如，
 * path = "/home/", => "/home"
 * path = "/a/./b/../../c/", => "/c"
 *
 * 边界情况:
 * 你是否考虑了 路径 = "/../" 的情况？
 * 在这种情况下，你需返回 "/" 。
 * 此外，路径中也可能包含多个斜杠 '/' ，如 "/home//foo/" 。
 * 在这种情况下，你可忽略多余的斜杠，返回 "/home/foo" 。
 *
 * @see https://leetcode-cn.com/problems/simplify-path/
 * @difficulty Medium
 */

import "strings"

func simplifyPath(path string) string {
	stack := []string{}
	s := strings.Split(path, "/")
	for _, v := range s {
		if v == "" || v == "." {
			continue
		} else if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}
