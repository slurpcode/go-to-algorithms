package main

import "fmt"

func PrintReverse(head *Node) {
	if head == nil {
		fmt.Println()
		return
	}

	PrintReverse(head.next)
	fmt.Printf("%d ", head.data)
}
