package main


// The easy and usual way
func DeleteNode(head *Node, data int) *Node {
    if data == head.data {
        return head.next
    }
    ret := head
    prev := head

    for ; head.data != data; {
        prev = head
        head = head.next
    }
    prev.next = head.next
    return ret
}

// Linus Torvalds way
func DeleteNodeCool(head **Node, data int) {
    indirect := head

    for ; (*indirect).data != data; {
        indirect = &(*indirect).next
    }
    *indirect = (*indirect).next
}
