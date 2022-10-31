/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Selection Sort Algorithm
 */

package main

import "fmt"

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func main() {
	fmt.Println("Selection Sort Algorithm")

	arr := []int{1, 5, 2, 4, 3}
	fmt.Println(SelectionSort(arr))

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(SelectionSort(arr))

	arr = []int{5, 4, 3, 2, 1}
	fmt.Println(SelectionSort(arr))

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(SelectionSort(arr))

	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(SelectionSort(arr))
}
