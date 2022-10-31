/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Floyd Warshall Algorithm
 */

package main

import "fmt"

const INF = 999999

func FloydWarshall(graph [][]int) [][]int {
	// 1. Create a distance matrix dist[][] and initialize all entries in it as INFINITE.
	// 2. Assign values to the distance matrix
	// 3. If there is a direct edge from i to j, then dist[i][j] = weight of edge.
	// 4. If i is equal to j, then dist[i][j] = 0.
	// 5. If there is no direct edge from i to j, then dist[i][j] = INFINITE.
	// 6. For each vertex k, do
	// 7.     For each vertex i, do
	// 8.         For each vertex j, do
	// 9.             If dist[i][k] + dist[k][j] < dist[i][j]
	// 10.                dist[i][j] = dist[i][k] + dist[k][j]
	// 11. Return dist[][]
	// 12. End

	dist := make([][]int, len(graph))
	for i := range dist {
		dist[i] = make([]int, len(graph))
	}

	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph); j++ {
			dist[i][j] = graph[i][j]
		}
	}

	for k := 0; k < len(graph); k++ {
		for i := 0; i < len(graph); i++ {
			for j := 0; j < len(graph); j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}

func main() {
	fmt.Println("Floyd Warshall Algorithm")

	graph := [][]int{
		{0, 5, INF, 10},
		{INF, 0, 3, INF},
		{INF, INF, 0, 1},
		{INF, INF, INF, 0},
	}

	fmt.Println(FloydWarshall(graph))
}
