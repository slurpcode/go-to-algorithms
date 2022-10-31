/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Topological Sort Algorithm
 */

package main

import "fmt"

func TopologicalSort(arr [][]int) []int {
	var result []int
	var visited []bool
	for i := 0; i < len(arr); i++ {
		visited = append(visited, false)
	}
	for i := 0; i < len(arr); i++ {
		if !visited[i] {
			TopologicalSortRecursive(arr, i, visited, &result)
		}
	}
	return result
}

func TopologicalSortRecursive(arr [][]int, i int, visited []bool, result *[]int) {
	visited[i] = true
	for j := 0; j < len(arr); j++ {
		if arr[i][j] == 1 && !visited[j] {
			TopologicalSortRecursive(arr, j, visited, result)
		}
	}
	*result = append(*result, i)
}

func main() {
	fmt.Println("Topological Sort Algorithm")

	arr := [][]int{
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 0},
	}
	fmt.Println(TopologicalSort(arr))
}
