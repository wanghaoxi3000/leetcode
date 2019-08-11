/*
9. 回文数

判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:
输入: 121
输出: true

示例 2:
输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

示例 3:
输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

进阶:
你能不将整数转为字符串来解决这个问题吗？
*/

package leetcode

/*
思路：取后半段的数字进行翻转后与前半段数字比对
1. 每次通过取余（% 10）操作，取出最低的数字 y = x % 10
2. 将最低的数字加到取出数字的末尾 revertNum = revertNum * 10 + y
3. 每取一个最低位的数字，x 都要自除以 10
4. 判断 x 是不是小于 revertNum，当它小于的时候，数字已经对半或过半了
5. 最后，判断奇偶数情况，偶数 revertNum 和 x 相等，奇数最中间的数字就在 revertNum 的最低位上，将它除以 10 后应该和 x 相等
*/
func isPalindrome(x int) bool {
	// x 小于 0 或末尾为 0 必定为 false
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertNum := 0
	for x > revertNum {
		revertNum = revertNum*10 + x%10
		x /= 10
	}
	return x == revertNum || x == revertNum/10
}
