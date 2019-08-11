/*
695. 岛屿的最大面积
给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。
找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)

示例 1:
[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。

示例 2:
[[0,0,0,0,0,0,0,0]]
对于上面这个给定的矩阵, 返回 0。

注意: 给定的矩阵grid 的长度和宽度都不超过 50。
*/

package leetcode

func maxAreaOfIsland(grid [][]int) int {
	var islandDfs func([][]int, int, int) int
	islandDfs = func(grid [][]int, i int, j int) int {
		if i < 0 || i > len(grid)-1 ||
			j < 0 || j > len(grid[i])-1 ||
			grid[i][j] != 1 {
			return 0
		}
		res := 1
		grid[i][j] = 2
		res += islandDfs(grid, i-1, j)
		res += islandDfs(grid, i+1, j)
		res += islandDfs(grid, i, j-1)
		res += islandDfs(grid, i, j+1)
		return res
	}

	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				area := islandDfs(grid, i, j)
				if area > ans {
					ans = area
				}
			}
		}
	}

	return ans
}
