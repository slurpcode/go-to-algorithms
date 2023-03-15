package main

import "fmt"

func PowerSet(nums []int) {
	size := (1<<len(nums))

	for counter := 0; counter < size; counter++ {
		for j := 0; j < len(nums); j++ {
			if (counter & (1 << j)) != 0 {
				fmt.Printf("%d ", nums[j])
			}
		}
		fmt.Println()
	}
}


func main() {
	nums := []int{1, 2, 7, 6, 5, 4, 3}
	PowerSet(nums)
}
