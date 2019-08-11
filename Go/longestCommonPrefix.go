/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:
输入: ["flower","flow","flight"]
输出: "fl"

示例 2:
输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。

说明:
所有输入只包含小写字母 a-z 。
*/

package leetcode

import "strings"

func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	start := strs[0]
	common := len(start)
	var same int

	for _, ch := range strs[1:] {
		same = 0
		for i := range ch {
			if i >= common {
				break
			}
			if ch[i] == start[i] {
				same++
			} else {
				break
			}
		}
		if same < common {
			common = same
		}

	}

	return start[0:common]
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	tt := strs[0]

	for _, v := range strs[1:] {
		for strings.Index(v, tt) != 0 {
			tt = tt[0 : len(tt)-1]
		}
	}

	return tt
}
