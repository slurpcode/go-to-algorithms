/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Kruskal’s Algorithm
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

func Kruskal(graph Graph) []Edge {
	var result []Edge
	var parent []int

	for i := 0; i < graph.Vertices; i++ {
		parent = append(parent, i)
	}

	// Sort edges by weight
	for i := 0; i < len(graph.Edges); i++ {
		for j := i + 1; j < len(graph.Edges); j++ {
			if graph.Edges[i].Weight > graph.Edges[j].Weight {
				graph.Edges[i], graph.Edges[j] = graph.Edges[j], graph.Edges[i]
			}
		}
	}

	for i := 0; i < len(graph.Edges); i++ {
		source := graph.Edges[i].Source
		destination := graph.Edges[i].Destination

		if Find(source, parent) != Find(destination, parent) {
			result = append(result, graph.Edges[i])
			Union(source, destination, parent)
		}
	}

	return result
}

func Find(vertex int, parent []int) int {
	if parent[vertex] == vertex {
		return vertex
	}

	return Find(parent[vertex], parent)
}

func Union(source int, destination int, parent []int) {
	sourceSet := Find(source, parent)
	destinationSet := Find(destination, parent)

	parent[sourceSet] = destinationSet
}

func main() {
	fmt.Println("Kruskal’s Algorithm")

	graph := Graph{
		Vertices: 6,
		Edges: []Edge{
			{0, 1, 4},
			{0, 2, 3},
			{1, 2, 1},
			{1, 3, 2},
			{2, 3, 4},
			{3, 4, 2},
			{4, 5, 6},
		},
	}

	result := Kruskal(graph)
	fmt.Println(result)
}
