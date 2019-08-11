/*
209. 长度最小的子数组
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。

示例:
输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。

进阶:
如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解
*/

package leetcode

/*
定义两个指针left和right，分别记录子数组的左右的边界位置。
1. 让right向右移，直到子数组和大于等于给定值或者right达到数组末尾；
2. 更新最短距离，将left像右移一位，sum减去移去的值；
3. 重复 1, 2 步骤，直到right到达末尾，且left到达临界位置

时间复杂度：O(n)
空间复杂度：O(1)
*/
func minSubArrayLen(s int, nums []int) int {
	res := len(nums) + 1
	sum := 0
	var left, right int = 0, 0

	for left < len(nums) {
		if right < len(nums) && sum < s {
			sum += nums[right]
			right++
		} else {
			sum -= nums[left]
			left++
		}

		if sum >= s && right-left < res {
			res = right - left
		}
	}

	if res > len(nums) {
		return 0
	}
	return res
}
