// The program below sorts an array of 0s, 1s and 2s using a three-way partitioning approach.

package main

import "fmt"

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func sortElements(arr []int) {
	start, mid, end := 0, 0, len(arr)-1
	pivot := 1

	for mid <= end {
		if arr[mid] < pivot {
			swap(arr, start, mid)
			start++
			mid++
		} else if arr[mid] > pivot {
			swap(arr, mid, end)
			end--
		} else {
			mid++
		}
	}
}

func main() {
	arr := []int{0, 1, 2, 2, 1, 0, 0, 2, 0, 1, 1, 0, 2, 1, 2, 2, 0, 2}

	fmt.Println("The original array:", arr)

	sortElements(arr)

	fmt.Println("After sorting, the elements in the array are:")
	fmt.Println(arr)
}