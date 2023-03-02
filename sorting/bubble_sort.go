package main

import (
	"fmt"
)

func BubbleSort(arr []int) {
    for i := 0; i < len(arr) - 1; i++ {
        for j := 0; j < len(arr) - i - 1; j++ {
            if arr[j] > arr[j + 1] {
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
            }
        }
    }
}

func main() {
    fmt.Println("Bubble sort algorithm")
    arr := []int{1, 3, 2, 88, 7, 11, 10, 29}
    fmt.Println(arr)
    fmt.Println("After sort")
    BubbleSort(arr)
    fmt.Println(arr)
}
