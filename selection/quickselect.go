/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Quickselect Algorithm
 */

package main

import "fmt"

func Quickselect(arr []int, k int) int {
	return QuickselectRecursive(arr, 0, len(arr)-1, k)
}

func QuickselectRecursive(arr []int, left int, right int, k int) int {
	if left == right {
		return arr[left]
	}

	pivotIndex := Partition(arr, left, right)

	if k == pivotIndex {
		return arr[k]
	} else if k < pivotIndex {
		return QuickselectRecursive(arr, left, pivotIndex-1, k)
	} else {
		return QuickselectRecursive(arr, pivotIndex+1, right, k)
	}
}

func Partition(arr []int, left int, right int) int {
	pivot := arr[right]
	pivotIndex := left

	for i := left; i < right; i++ {
		if arr[i] < pivot {
			arr[i], arr[pivotIndex] = arr[pivotIndex], arr[i]
			pivotIndex++
		}
	}

	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	return pivotIndex
}

func main() {
	fmt.Println("Quickselect Algorithm")

	arr := []int{7, 10, 4, 3, 20, 15}
	fmt.Println(Quickselect(arr, 3))

	arr = []int{1, 5, 2, 4, 3}
	fmt.Println(Quickselect(arr, 3))

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(Quickselect(arr, 3))
}
