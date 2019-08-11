/*
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

示例:
输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]
*/

package leetcode

import (
	"strconv"
	"strings"
)

// 使用 DFS 深度优先搜索算法
func restoreIPAddresses(s string) []string {
	res := make([]string, 0)
	rightInit := strings.Split(s, "")

	/*
		定义 DFS 搜索递归函数
		right: 当前段终点
		left: 每段的起始点
		depth: 当前是 IP 中第几段
	*/
	var findAddress func([]string, []string, int)
	findAddress = func(right []string, left []string, depth int) {
		if len(right) == 0 && depth == 0 { // 结束条件: IP 最后一段没有多余的字符串
			res = append(res, strings.Join(left, "."))
		}

		for i := range right {
			if ((len(right[:i+1])) > 1 && right[0] == "0") || len(right) > depth*3 { // 每段不得有前置零, 长度不超过 3
				break
			}

			chInt, _ := strconv.Atoi(strings.Join(right[:i+1], ""))
			if chInt >= 0 && chInt <= 255 {
				findAddress(right[i+1:], append(left, strings.Join(right[:i+1], "")), depth-1)
			} else {
				break
			}
		}
	}

	leftInit := make([]string, 0)
	findAddress(rightInit, leftInit, 4)
	return res
}
