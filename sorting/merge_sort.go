/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Merge Sort Algorithm
 */

package main

import "fmt"

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])

	return Merge(left, right)
}

func Merge(left []int, right []int) []int {
	var result []int

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) > 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) > 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

func main() {
	fmt.Println("Merge Sort Algorithm")

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(MergeSort(arr))
}
