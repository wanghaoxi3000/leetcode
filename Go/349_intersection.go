package leetcode

/*
349. 两个数组的交集

给定两个数组，编写一个函数来计算它们的交集。

示例 1:
输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2]

示例 2:
输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [9,4]

说明:
输出结果中的每个元素一定是唯一的。
我们可以不考虑输出结果的顺序。
*/

func intersection(nums1 []int, nums2 []int) []int {
	record := make(map[int]int)
	var result []int

	for _, v := range nums1 {
		record[v] = 0
	}

	for _, v := range nums2 {
		if _, ok := record[v]; ok && record[v] == 0 {
			record[v]++
			result = append(result, v)
		}
	}

	return result
}
