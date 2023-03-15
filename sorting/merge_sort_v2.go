package main

import "fmt"

func MergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	m := l + (r-l)/2
	MergeSort(arr, l, m)
	MergeSort(arr, m+1, r)
	merge := func() {
		nl := m - l + 1
		nr := r - m
		L := make([]int, nl)
		R := make([]int, nr)

		copy(L, arr[l:m+1])
		copy(R, arr[m+1:r+1])
		i, j, k := 0, 0, l

		for i < nl && j < nr {
			if L[i] < R[j] {
				arr[k] = L[i]
				i++
			} else {
				arr[k] = R[j]
				j++
			}
			k++
		}

		for i < nl {
			arr[k] = L[i]
			k++
			i++
		}
		for j < nr {
			arr[k] = R[j]
			k++
			j++
		}
	}
	merge()
}

func main() {
	fmt.Println("Merge Sort Algorithm")

	arr := []int{1, 8, 2, 9, 11, 3, 92, 7, 10, 5}
	MergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
