/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Binary Search Algorithm
 */
package main

import "fmt"

func BinarySearch(arr []int, target int) int {
	// 1. Create a sorted array
	// 2. Find the middle element
	// 3. Compare the middle element with the target value
	// 4. If the target value is equal to the middle element, we found it. return the index.
	// 5. If the target value is less than the middle element, then the target value should be in the left subarray.
	// 6. If the target value is greater than the middle element, then the target value should be in the right subarray.
	// 7. Repeat step 2 to 6 until the target value is found or the subarray is empty.

	middle := len(arr) / 2
	if arr[middle] == target {
		return middle
	}
	else if arr[middle] > target {
		return BinarySearch(arr[:middle], target)
	}
	else if arr[middle] < target {
		return BinarySearch(arr[middle:], target)
	}
	return -1
}
func main() {
	fmt.Println("Binary Search Algorithm")

	// 1. Create a sorted array
	// 2. Find the middle element
	// 3. Compare the middle element with the target value
	// 4. If the target value is equal to the middle element, we found it. return the index.
	// 5. If the target value is less than the middle element, then the target value should be in the left subarray.
	// 6. If the target value is greater than the middle element, then the target value should be in the right subarray.
	// 7. Repeat step 2 to 6 until the target value is found or the subarray is empty.

	// 1. Create a sorted array
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// BinarySearch
	fmt.Println(BinarySearch(arr, 5))
	

	// 2. Find the middle element
	middle := len(arr) / 2

	// 3. Compare the middle element with the target value
	target := 5

	// 4. If the target value is equal to the middle element, we found it. return the index.
	if arr[middle] == target {
		fmt.Println("Found")
	}

	// 5. If the target value is less than the middle element, then the target value should be in the left subarray.
	if arr[middle] > target {
		fmt.Println("Search in left subarray")
	}

	// 6. If the target value is greater than the middle element, then the target value should be in the right subarray.
	if arr[middle] < target {
		fmt.Println("Search in right subarray")
	}

	// 7. Repeat step 2 to 6 until the target value is found or the subarray is empty.
}
