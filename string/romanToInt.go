package main

import "fmt"

func romanToInt(s string) int {
	var romanMap = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var result int = 0
	var prev int = 0

	for i := len(s) - 1; i >= 0; i-- {
		curr := romanMap[string(s[i])]
		if curr < prev {
			result -= curr
		} else {
			result += curr
		}
		prev = curr
	}
	return result
}

func main() {
	fmt.Println("Code Solve Question Roman To Integer")
	fmt.Println(romanToInt("MXI"))
}
