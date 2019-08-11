/*
75. 颜色分类
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

注意:
不能使用代码库中的排序函数来解决这道题。

示例:
输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]

进阶：
一个直观的解决方案是使用计数排序的两趟扫描算法。
首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
你能想出一个仅使用常数空间的一趟扫描算法吗？
*/

package leetcode

/*
三路快排思想
设定两个索引，一个从左往右滑动 zero，一个从从右往左滑动 two， 遍历 nums
当 nums[i] 的值为 1 时，i++
当 nums[i] 的值为 2 时，two 的值先减 1，而后交换 nums[i] 与 nums[two]， 此时再观察nums[i]的值
当 nums[i] 的值为 0 时，zero++，而后交换 nums[i] 与 nums[zero]
当 i = two时，结束循环

对整个数组只遍历一遍
时间复杂度：O(n)
空间复杂度：O(1)
*/
func sortColors(nums []int) {
	zero := -1       // [0...zero] == 0
	two := len(nums) // [two...n-1] == 2

	for i := 0; i < two; {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			two--
			nums[i], nums[two] = nums[two], nums[i]
		} else {
			zero++
			nums[zero], nums[i] = nums[i], nums[zero]
			i++
		}
	}
}
