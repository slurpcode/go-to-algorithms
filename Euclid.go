/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Euclid’s Algorithm
 */

package main

import "fmt"

func Euclid(a int, b int) int {
	if b == 0 {
		return a
	}
	return Euclid(b, a%b)
}

func main() {
	fmt.Println("Euclid’s Algorithm")

	fmt.Println(Euclid(12, 18))

	fmt.Println(Euclid(18, 12))

	fmt.Println(Euclid(12, 0))

	fmt.Println(Euclid(0, 12))

	fmt.Println(Euclid(0, 0))

	fmt.Println(Euclid(1, 0))

	fmt.Println(Euclid(0, 1))

	fmt.Println(Euclid(1, 1))

	fmt.Println(Euclid(2, 1))

	fmt.Println(Euclid(1, 2))
}
