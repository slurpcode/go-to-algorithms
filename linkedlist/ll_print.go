package main

import "fmt"

func TraverseList(head *Node) {
    for ; head != nil; head = head.next {
        fmt.Printf("%d ", head.data)
    }
}

func TraverseListRecursive(head *Node) {
    if head == nil {
        return
    }
    fmt.Printf("%d ", head.data)
    TraverseListRecursive(head.next)
}
