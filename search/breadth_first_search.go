/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Breadth First Search (BFS) Algorithm
 */

package main

import "fmt"

func BFS(graph map[string][]string, start string) []string {
	var queue []string
	queue = append(queue, start)

	var visited []string

	for len(queue) > 0 {
		// 1. Dequeue a vertex from queue and add it to visited list
		vertex := queue[0]
		queue = queue[1:]
		visited = append(visited, vertex)

		// 2. Get all adjacent vertices of the dequeued vertex s. If a adjacent has not been visited, then mark it visited and enqueue it
		for _, neighbour := range graph[vertex] {
			if !contains(visited, neighbour) {
				queue = append(queue, neighbour)
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
	fmt.Println("Breadth First Search Algorithm")

	graph := make(map[string][]string)
	graph["A"] = []string{"B", "C"}
	graph["B"] = []string{"A", "D"}
	graph["C"] = []string{"A", "D"}
	graph["D"] = []string{"E", "B", "C"}
	graph["E"] = []string{"D"}

	fmt.Println(BFS(graph, "A"))
}
