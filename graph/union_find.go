/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Union Find Algorithm
 */

package main

import "fmt"

type Graph struct {
	Vertices int
	Edges    []Edge
}

type Edge struct {
	Source      int
	Destination int
}

func UnionFind(graph Graph) []Edge {
	var result []Edge
	var parent []int

	for i := 0; i < graph.Vertices; i++ {
		parent = append(parent, i)
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
	fmt.Println("Union Find Algorithm")

	graph := Graph{
		Vertices: 3,
		Edges: []Edge{
			{Source: 0, Destination: 1},
			{Source: 1, Destination: 2},
			{Source: 0, Destination: 2},
		},
	}

	fmt.Println(UnionFind(graph))
}
