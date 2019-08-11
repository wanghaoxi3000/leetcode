/*
219. 给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值不大于 k。

示例 1:
输入: nums = [1,2,3,1], k = 3
输出: true

示例 2:
输入: nums = [1,0,1,1], k = 1
输出: true

示例 3:
输入: nums = [1,2,3,1,2,3], k = 2
输出: false
*/

package leetcode

/*思路:
1. 设置查找表 record, 遍历数组 nums, 每次遍历时在 record 中查找是否存在相同元素, 存在则返回 true
2. 若 record 中未找到, 则将该元素插入到 record 中
3. 查看 record 中长度是否为 k+1, 若已到达 k+1, 则删除 record 中 nums[i-k] 元素
4. 遍历完 nums 未查到则返回 false

时间复杂度 O(n)
空间复杂度 O(k)
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}

	if k < 0 {
		return false
	}

	record := make(map[int]bool, k)
	for i, v := range nums {
		if ok, _ := record[v]; ok {
			return true
		}

		record[v] = true

		// 保持record中最多有k个元素
		if len(record) == k+1 {
			delete(record, nums[i-k])
		}
	}

	return false
}
