package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func optimalRoute(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if cols == 0 {
		return 0
	}

	// Create a DP array to store the maximum sums
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}

	// Initialize the starting point
	dp[rows-1][0] = grid[rows-1][0]

	// Fill the first column (moving up only)
	for i := rows - 2; i >= 0; i-- {
		dp[i][0] = dp[i+1][0] + grid[i][0]
	}

	// Fill the first row (moving right only)
	for j := 1; j < cols; j++ {
		dp[rows-1][j] = dp[rows-1][j-1] + grid[rows-1][j]
	}

	// Fill the rest of the DP table
	for i := rows - 2; i >= 0; i-- {
		for j := 1; j < cols; j++ {
			dp[i][j] = max(dp[i+1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[0][cols-1]
}

func main() {
	// Define your grid
	grid := [][]int{
		{3, 2, 1, 5},
		{1, 9, 2, 3},
		{0, 4, 7, 1},
	}

	// Call the optimalRoute function and print the result
	result := optimalRoute(grid)
	fmt.Println("Optimal route sum:", result)
}
