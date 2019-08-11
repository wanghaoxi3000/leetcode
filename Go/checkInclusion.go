/*
字符串的排列

给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:
输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").

示例2:
输入: s1= "ab" s2 = "eidboaoo"
输出: False

注意：
    输入的字符串只包含小写字母
    两个字符串的长度都在 [1, 10,000] 之间
*/

package leetcode

import (
	"strings"
)

// 从 s2 第一个字符开始, 每个字符 s1 长度的字符串各字符数量等于 s1 字符数量, 即为 true
func checkInclusion1(s1 string, s2 string) bool {
	s1Len := len(s1)
	s2Len := len(s2)
	if s1Len > s2Len {
		return false
	}

	chMap := make(map[byte]int)
	reduceMap := make(map[byte]int)
	for i := 0; i < s1Len; i++ {
		if _, ok := chMap[s1[i]]; ok {
			chMap[s1[i]]++
		} else {
			chMap[s1[i]] = 1
			reduceMap[s1[i]] = 0
		}
	}

	index := 0
	for i := 0; i < s2Len-s1Len+1; i++ {
		for index = range s1 {
			if _, ok := chMap[s2[i+index]]; ok && reduceMap[s2[i+index]] < chMap[s2[i+index]] {
				reduceMap[s2[i+index]]++
			} else {
				if index > 0 {
					for k := range reduceMap {
						reduceMap[k] = 0
					}
				}
				index--
				break
			}
		}
		if index == s1Len-1 {
			return true
		}
	}

	return false
}

// 使用一个 map 记录 s1 中各单词数量，另一个 map 记录 s2 在 s1 相同长度下的各单词数量，两 map 相同即包含
func checkInclusion2(s1 string, s2 string) bool {
	s1Len := len(s1)
	s2Len := len(s2)
	if s1Len > s2Len {
		return false
	}
	var target, window [26]int

	for i := range s1 {
		target[s1[i]-'a']++
	}

	for i := 0; i < s2Len; i++ {
		window[s2[i]-'a']++
		if i == s1Len-1 {
			if window == target {
				return true
			}
		} else if s1Len <= i {
			window[s2[i-s1Len]-'a']--
			if window == target {
				return true
			}
		}
	}

	return false
}

/*
permutation: 全排列回溯法, 此题使用全排列会存在超时的问题
1) n个元素的全排列 = (n-1个元素的全排列)+(另一个元素作为前缀)
2) 出口: 如果只有一个元素的全排列, 则说明已经排完, 则输出数组;
3) 不断将每个元素放左第一个元素, 然后将它作为前缀, 并将其余元素继续全排列

             a|b|c(状态A)
               |
               |swap(0,0)
               |
             a|b|c(状态B)
             /  \
   swap(1,1)/    \swap(1,2)
           /      \
		 a|b|c   a|c|b

            a|b|c(状态A)
               |
               |swap(0,1)
               |
             b|a|c(状态B)
             /  \
   swap(1,1)/    \swap(1,2)
           /      \
		 b|a|c   b|c|a

            a|b|c(状态A)
               |
               |swap(0,2)
               |
             c|b|a(状态B)
             /  \
   swap(1,1)/    \swap(1,2)
           /      \
         c|b|a   c|a|b
*/
func checkInclusion3(s1 string, s2 string) bool {
	chMap := make(map[string]bool)
	contain := false
	var permutation func(chs []byte, i int)
	permutation = func(chs []byte, i int) {
		if i == len(chs)-1 {
			str := string(chs)
			if _, ok := chMap[str]; !ok {
				if strings.Contains(s2, str) {
					contain = true
				}
				chMap[str] = true
			}
			return
		}
		for k := i; k < len(chs); k++ {
			chs[i], chs[k] = chs[k], chs[i]
			permutation(chs, i+1)
			if contain {
				break
			}
			chs[i], chs[k] = chs[k], chs[i]
		}
	}
	permutation([]byte(s1), 0)
	return contain
}
