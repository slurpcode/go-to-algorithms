/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Depth First Search (DFS) Algorithm
 */

package main

import "fmt"

func DFS(graph map[string][]string, start string) []string {
	var stack []string
	stack = append(stack, start)

	var visited []string

	for len(stack) > 0 {
		// 1. Pop a vertex from stack and add it to visited list
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited = append(visited, vertex)

		// 2. Get all adjacent vertices of the popped vertex s. If a adjacent has not been visited, then mark it visited and push it
		for _, neighbour := range graph[vertex] {
			if !contains(visited, neighbour) {
				stack = append(stack, neighbour)
			}
		}
	}

	return visited
}

func contains(arr []string, str string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Depth First Search Algorithm")

	graph := make(map[string][]string)
	graph["A"] = []string{"B", "C"}
	graph["B"] = []string{"A", "D"}
	graph["C"] = []string{"A", "D"}
	graph["D"] = []string{"E", "B", "C"}
	graph["E"] = []string{"D"}

	fmt.Println(DFS(graph, "A"))
}
