package leetcode

/*
447. 回旋镖的数量

给定平面上 n 对不同的点，“回旋镖” 是由点表示的元组 (i, j, k) ，其中 i 和 j 之间的距离和 i 和 k 之间的距离相等（需要考虑元组的顺序）。
找到所有回旋镖的数量。你可以假设 n 最大为 500，所有点的坐标在闭区间 [-10000, 10000] 中。

示例:
输入:
[[0,0],[1,0],[2,0]]

输出:
2

解释:
两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]
*/

/*
思路:
1. 若点 a 和 b 的距离与点 a 和 c 的距离相等, 那么就有 abc 和 acb 两种排列方法
2. 若点 b c d 都分别和 a 之间的距离相等, 那么就有 6 中排列方法 abc acb acd adc abd adb
3. n 个点和 a 距离相等, 那么排列方法为 n(n-1)

时间复杂度: O(n^2)
*/
func numberOfBoomerangs(points [][]int) int {
	var ret int
	for i := 0; i < len(points); i++ {
		record := make(map[int]int)
		for j := 0; j < len(points); j++ {
			if i != j {
				record[distance(points[i], points[j])]++
			}
		}
		for _, v := range record {
			ret += v * (v - 1)
		}
	}

	return ret
}

func distance(a []int, b []int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}
