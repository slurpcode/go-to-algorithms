/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Kadane’s Algorithm
 */

package main

import "fmt"

func Kadane(arr []int) int {
	max := 0
	current := 0

	for _, item := range arr {
		current = current + item
		if current < 0 {
			current = 0
		}
		if max < current {
			max = current
		}
	}

	return max
}

func main() {
	fmt.Println("Kadane’s Algorithm")

	arr := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	fmt.Println(Kadane(arr))
}
