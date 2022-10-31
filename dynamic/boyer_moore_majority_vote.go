/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Boyer–Moore Majority Vote Algorithm
 */

package main

import "fmt"

func BoyerMooreMajorityVote(arr []int) int {
	var candidate int
	var count int

	for _, item := range arr {
		if count == 0 {
			candidate = item
		}
		if candidate == item {
			count++
		} else {
			count--
		}
	}

	return candidate
}

func main() {
	fmt.Println("Boyer–Moore Majority Vote Algorithm")

	arr := []int{1, 2, 3, 1, 1, 2, 1}
	fmt.Println(BoyerMooreMajorityVote(arr))
}
