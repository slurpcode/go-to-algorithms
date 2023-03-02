package main

import "fmt"



func main() {
    head := InitNode(1)
    second := InitNode(2)
    third := InitNode(4)
    last := InitNode(5)

    head.next = second
    second.next = third
    third.next = last

    TraverseList(head)
    fmt.Println()
    TraverseListRecursive(head)
    fmt.Println()

    head = DeleteNode(head, 2)
    TraverseList(head)
    DeleteNodeCool(&head, 1)
    TraverseList(head)

    PrintReverse(head)
}
