/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Kahn’s Topological Sort Algorithm
 */

package main

import "fmt"

func KahnTopologicalSort(arr [][]int) []int {
	var result []int
	var visited []bool
	var queue []int
	for i := 0; i < len(arr); i++ {
		visited = append(visited, false)
	}
	for i := 0; i < len(arr); i++ {
		if !visited[i] {
			queue = append(queue, i)
			visited[i] = true
		}
		for len(queue) > 0 {
			item := queue[0]
			queue = queue[1:]
			result = append(result, item)
			for j := 0; j < len(arr); j++ {
				if arr[item][j] == 1 && !visited[j] {
					queue = append(queue, j)
					visited[j] = true
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println("Kahn’s Topological Sort Algorithm")

	arr := [][]int{
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{0, 1, 1, 0, 0, 0},
	}
	fmt.Println(KahnTopologicalSort(arr))
}
