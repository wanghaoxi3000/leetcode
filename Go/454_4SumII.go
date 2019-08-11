package leetcode

/*
454. 四数相加 II

给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。
为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -228 到 228 - 1 之间，最终结果不会超过 231 - 1 。

例如:
输入:
A = [ 1, 2]
B = [-2,-1]
C = [-1, 2]
D = [ 0, 2]

输出:
2

解释:
两个元组如下:
1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
*/

/*
1. 将 A 和 B 之间两两之和都求出来，在哈希表中建立两数之和与次数的映射
2. 遍历 C 和 D 之间任意两数之和，与哈希表中存在相反数记录的值相加
*/
func fourSumCount(A []int, B []int, C []int, D []int) int {
	checkSum := make(map[int]int, len(A)*len(A))

	for _, itemA := range A {
		for _, itemB := range B {
			checkSum[itemA+itemB]++
		}
	}

	res := 0
	for _, itemC := range C {
		for _, itemD := range D {
			res += checkSum[-(itemC + itemD)]
		}
	}

	return res
}
