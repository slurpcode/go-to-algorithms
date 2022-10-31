/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: Huffman Coding Compression Algorithm
 */

package main

import (
	"container/heap"
	"fmt"
)

type HuffmanNode struct {
	Left  *HuffmanNode
	Right *HuffmanNode
	Freq  int
	Char  string
}

type HuffmanHeap []*HuffmanNode

func (h HuffmanHeap) Len() int {
	return len(h)
}

func (h HuffmanHeap) Less(i, j int) bool {
	return h[i].Freq < h[j].Freq
}

func (h HuffmanHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HuffmanHeap) Push(x interface{}) {
	*h = append(*h, x.(*HuffmanNode))
}

func (h *HuffmanHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func HuffmanCodingCompression(text string) string {
	freq := make(map[string]int)
	for _, c := range text {
		freq[string(c)]++
	}
	h := &HuffmanHeap{}
	heap.Init(h)
	for k, v := range freq {
		heap.Push(h, &HuffmanNode{Freq: v, Char: k})
	}
	for h.Len() > 1 {
		left := heap.Pop(h).(*HuffmanNode)
		right := heap.Pop(h).(*HuffmanNode)
		top := &HuffmanNode{Left: left, Right: right, Freq: left.Freq + right.Freq}
		heap.Push(h, top)
	}
	root := heap.Pop(h).(*HuffmanNode)
	var result string
	for _, c := range text {
		result += HuffmanCodingCompressionRecursive(root, string(c), "")
	}
	return result
}

func HuffmanCodingCompressionRecursive(root *HuffmanNode, char string, result string) string {
	if root == nil {
		return ""
	}
	if root.Char == char {
		return result
	}
	return HuffmanCodingCompressionRecursive(root.Left, char, result+"0") + HuffmanCodingCompressionRecursive(root.Right, char, result+"1")
}

func main() {
	fmt.Println("Huffman Coding Compression Algorithm")

	fmt.Println(HuffmanCodingCompression("ABABDABACDABABCABAB"))
	fmt.Println(HuffmanCodingCompression("ABRACADABRA"))
	fmt.Println(HuffmanCodingCompression("ABRACADABRAABRACADABRA"))
	fmt.Println(HuffmanCodingCompression("ABRACADABRAABRACADABRAABRACADABRA"))
}
