/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Flood Fill Algorithm
 */

package main

import "fmt"

func FloodFill(arr [][]int, x int, y int, newColor int) [][]int {
	oldColor := arr[x][y]
	if oldColor == newColor {
		return arr
	}
	return FloodFillRecursive(arr, x, y, oldColor, newColor)
}

func FloodFillRecursive(arr [][]int, x int, y int, oldColor int, newColor int) [][]int {
	if x < 0 || y < 0 || x >= len(arr) || y >= len(arr[0]) || arr[x][y] != oldColor {
		return arr
	}
	arr[x][y] = newColor
	FloodFillRecursive(arr, x+1, y, oldColor, newColor)
	FloodFillRecursive(arr, x-1, y, oldColor, newColor)
	FloodFillRecursive(arr, x, y+1, oldColor, newColor)
	FloodFillRecursive(arr, x, y-1, oldColor, newColor)
	return arr
}

func main() {
	fmt.Println("Flood Fill Algorithm")

	arr := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	fmt.Println(FloodFill(arr, 1, 1, 2))
}
