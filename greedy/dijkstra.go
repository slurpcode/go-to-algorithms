/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Dijkstra’s Algorithm
 */

package main

import "fmt"

func Dijkstra(graph map[string]map[string]int, costs map[string]int, parents map[string]string) map[string]int {
	processed := []string{}

	node := findLowestCostNode(costs, processed)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]
		for n := range neighbors {
			newCost := cost + neighbors[n]
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}
		processed = append(processed, node)
		node = findLowestCostNode(costs, processed)
	}

	return costs
}

func findLowestCostNode(costs map[string]int, processed []string) string {
	lowestCost := 1000000
	lowestCostNode := ""

	for node := range costs {
		cost := costs[node]
		if cost < lowestCost && !contains(processed, node) {
			lowestCost = cost
			lowestCostNode = node
		}
	}

	return lowestCostNode
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
	fmt.Println("Dijkstra’s Algorithm")

	graph := make(map[string]map[string]int)
	graph["start"] = make(map[string]int)
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"] = make(map[string]int)
	graph["a"]["fin"] = 1

	graph["b"] = make(map[string]int)
	graph["b"]["a"] = 3
	graph["b"]["fin"] = 5

	graph["fin"] = make(map[string]int)

	costs := make(map[string]int)
	costs["a"] = 6
	costs["b"] = 2
	costs["fin"] = 1000000

	parents := make(map[string]string)
	parents["a"] = "start"
	parents["b"] = "start"
	parents["fin"] = ""

	fmt.Println(Dijkstra(graph, costs, parents))
}
