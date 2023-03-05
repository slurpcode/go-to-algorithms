package main

import "fmt"

var dp [256]int

func fibo(nth int) int {
	if nth <= 2 {
		return 1
	}

	if dp[nth] != 0 {
		return dp[nth]
	}

	dp[nth] = fibo(nth - 2) + fibo(nth - 1)
	return dp[nth]
}



func main() {
	fmt.Println(fibo(5))
}
