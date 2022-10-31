/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Floyd’s Cycle Detection Algorithm
 */

package main

import "fmt"

func FloydCycleDetection(arr []int) bool {
	slow := arr[0]
	fast := arr[0]

	for {
		slow = arr[slow]
		fast = arr[arr[fast]]

		if slow == fast {
			break
		}
	}

	if slow == 0 {
		return false
	}

	slow = arr[0]

	for slow != fast {
		slow = arr[slow]
		fast = arr[fast]
	}

	return true
}

func main() {
	fmt.Println("Floyd’s Cycle Detection Algorithm")

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(FloydCycleDetection(arr))
}
