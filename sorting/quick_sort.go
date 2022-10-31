/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Quicksort Algorithm
 */

package main

import "fmt"

func Quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left []int
	var right []int

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = Quicksort(left)
	right = Quicksort(right)

	return append(append(left, pivot), right...)
}

func main() {
	fmt.Println("Quicksort Algorithm")

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(Quicksort(arr))
}
