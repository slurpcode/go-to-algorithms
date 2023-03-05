package main

import "fmt"

func solve(arr []int, sum int) int {
	dp := make([]int, sum + 1)
	dp[0] = 0

	for i := 1; i <= sum; i++ {
		dp[i] = 1e8
		for j := 1; j < len(arr); j++ {
			if i - arr[j] >= 0 {
				ok := dp[i] > dp[i - arr[j]]+1; if ok {
					dp[i] = dp[i - arr[j]] + 1
				}
			}
		}
	}
	return dp[sum]
}

func main() {
	coins := []int{1, 3, 4}
	sum := 38

	fmt.Println(solve(coins, sum))
}
