/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Lee Algorithm
 */

package main

import "fmt"

func Lee(arr [][]int) int {
	m := len(arr)
	n := len(arr[0])
	// fmt.Println(m, n)

	// 1. Create a matrix to store the results
	// of subproblems
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 2. Initialize the first row and column
	// of the dp matrix
	dp[0][0] = arr[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + arr[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + arr[0][j]
	}

	// 3. Compute the rest of the dp matrix
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + arr[i][j]
		}
	}

	// 4. Return the last cell of the dp matrix
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Lee Algorithm")

	arr := [][]int{
		{1, 2, 3},
		{4, 8, 2},
		{1, 5, 3},
	}
	fmt.Println(Lee(arr))

	arr = [][]int{
		{4, 8, 2},
		{1, 5, 3},
		{6, 3, 1},
	}
	fmt.Println(Lee(arr))
}
