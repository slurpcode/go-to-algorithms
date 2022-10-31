/*
 * Author: Max Base
 * Date: 2022-10-28
 * Name: KMP Algorithm
 */

package main

import "fmt"

func KMP(text string, pattern string) int {
	lps := LPS(pattern)
	i := 0
	j := 0
	for i < len(text) {
		if text[i] == pattern[j] {
			i++
			j++
		}
		if j == len(pattern) {
			return i - j
		} else if i < len(text) && text[i] != pattern[j] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return -1
}

func LPS(pattern string) []int {
	lps := make([]int, len(pattern))
	i := 1
	j := 0
	for i < len(pattern) {
		if pattern[i] == pattern[j] {
			j++
			lps[i] = j
			i++
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func main() {
	fmt.Println("KMP Algorithm")

	fmt.Println(KMP("ABABDABACDABABCABAB", "ABABCABAB"))
}
