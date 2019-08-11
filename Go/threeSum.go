/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

package leetcode

import "sort"

/*
1. 将数组排序
2.定义三个指针，i，j，k。遍历i
3. 这个问题就可以转化为在i之后的数组中寻找 nums[j] + nums[k] = -nums[i]
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int

	for c := len(nums) - 1; c >= 2; {
		for a, b := 0, c-1; a < b; {
			abSum := nums[a] + nums[b]
			if abSum < -nums[c] {
				a++
			} else if abSum > -nums[c] {
				b--
			} else {
				ret = append(ret, []int{nums[a], nums[b], nums[c]})
				// 去重复 a b
				for {
					a++
					if a >= b || nums[a-1] != nums[a] {
						break
					}
				}
				for {
					b--
					if a >= b || nums[b+1] != nums[b] {
						break
					}
				}
			}
		}
		// 去重复 c
		for {
			c--
			if c < 2 || nums[c+1] != nums[c] {
				break
			}
		}
	}

	return ret
}
