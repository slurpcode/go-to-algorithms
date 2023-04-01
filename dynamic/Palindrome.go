package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var y int = x
	var z int = 0
	for y != 0 {
		z = z*10 + y%10
		y = y / 10
	}
	if z == x {
		return true
	}
	return false
}

func main() {
	fmt.Println("Palindrome Algorithm")

	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
}
