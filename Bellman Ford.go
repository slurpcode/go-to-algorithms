/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Bellman Ford Algorithm
 */

package main

import "fmt"

type Edge struct {
	Source      int
	Destination int
	Weight      int
}

type Graph struct {
	Vertices int
	Edges    []Edge
}

func BellmanFord(graph Graph, source int) []int {
	var distance []int

	for i := 0; i < graph.Vertices; i++ {
		distance = append(distance, 999)
	}

	distance[source] = 0

	for i := 0; i < graph.Vertices-1; i++ {
		for j := 0; j < len(graph.Edges); j++ {
			u := graph.Edges[j].Source
			v := graph.Edges[j].Destination
			weight := graph.Edges[j].Weight

			if distance[u]+weight < distance[v] {
				distance[v] = distance[u] + weight
			}
		}
	}

	for i := 0; i < len(graph.Edges); i++ {
		u := graph.Edges[i].Source
		v := graph.Edges[i].Destination
		weight := graph.Edges[i].Weight

		if distance[u]+weight < distance[v] {
			fmt.Println("Graph contains negative weight cycle")
			return nil
		}
	}

	return distance
}

func main() {
	fmt.Println("Bellman Ford Algorithm")

	graph := Graph{
		Vertices: 5,
		Edges: []Edge{
			{0, 1, -1},
			{0, 2, 4},
			{1, 2, 3},
			{1, 3, 2},
			{1, 4, 2},
			{3, 2, 5},
			{3, 1, 1},
			{4, 3, -3},
		},
	}

	distance := BellmanFord(graph, 0)
	fmt.Println(distance)
}
