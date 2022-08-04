package CodeTop

func minPathSum(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return -1
	}

	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid[0]); i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}
	for i := 1; i < len(grid); i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}
