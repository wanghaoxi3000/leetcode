/*
167. 两数之和 II - 输入有序数组

给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。

说明:
返回的下标值（index1 和 index2）不是从零开始的。
你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。

示例:
输入: numbers = [2, 7, 11, 15], target = 9
输出: [1,2]
解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
*/

package leetcode

// 通过已排序的特性使用左右两个指针寻找
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func twoSum2(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	sum := 0
	ret := []int{l, r}
	for l < r {
		sum = numbers[l] + numbers[r]
		if sum == target {
			ret[0] = l + 1
			ret[1] = r + 1
			break
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return ret
}
