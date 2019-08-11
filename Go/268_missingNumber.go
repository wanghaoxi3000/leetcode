/*
268. 缺失数字

给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。

示例 1:
输入: [3,0,1]
输出: 2

示例 2:
输入: [9,6,4,2,3,5,7,0,1]
输出: 8

说明:
你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现?
*/

package leetcode

/*
异或法
将少了一个数的数组与 0 到 n 之间完整的那个数组进行异或处理
因为相同的数字异或会变为了 0 ，那么全部数字异或后，剩下的就是少了的那个数字
*/
func missingNumberXor(nums []int) int {
	res := 0
	i := 0
	for i = 0; i < len(nums); i++ {
		res = res ^ i ^ nums[i]
	}
	return res ^ i
}

/*
求和法
求出 0 到 n 之间所有的数字之和
遍历数组计算出原始数组中数字的累积和
两和相减，差值就是丢失的那个数字
*/
func missingNumberSum(nums []int) int {
	length := len(nums)
	sum := length * (length + 1) / 2
	for i := range nums {
		sum -= nums[i]
	}
	return sum
}
