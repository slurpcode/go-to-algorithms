/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Heap Sort Algorithm
 */

package main

import "fmt"

func HeapSort(arr []int) []int {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		HeapSortRecursive(arr, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		HeapSortRecursive(arr, i, 0)
	}
	return arr
}

func HeapSortRecursive(arr []int, n int, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		HeapSortRecursive(arr, n, largest)
	}
}

func main() {
	fmt.Println("Heap Sort Algorithm")

	arr := []int{1, 5, 2, 4, 3}
	fmt.Println(HeapSort(arr))

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(HeapSort(arr))

	arr = []int{5, 4, 3, 2, 1}
	fmt.Println(HeapSort(arr))

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(HeapSort(arr))

	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(HeapSort(arr))
}
