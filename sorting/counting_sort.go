/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Counting Sort Algorithm
 */

package main

import "fmt"

func CountingSort(arr []int) []int {
	var result []int
	var count []int
	for i := 0; i < len(arr); i++ {
		count = append(count, 0)
	}
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	for i := 0; i < len(count); i++ {
		for j := 0; j < count[i]; j++ {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	fmt.Println("Counting Sort Algorithm")

	arr := []int{1, 5, 2, 4, 3}
	fmt.Println(CountingSort(arr))

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(CountingSort(arr))

	arr = []int{5, 4, 3, 2, 1}
	fmt.Println(CountingSort(arr))

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(CountingSort(arr))

	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(CountingSort(arr))
}
