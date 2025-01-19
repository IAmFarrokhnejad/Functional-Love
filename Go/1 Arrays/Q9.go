// The program below sorts an array in both ascending and descending order using the Quicksort algorithm.
// Quicksort is one of the most efficient sorting algorithms with an average time complexity of 𝑂(𝑛 log 𝑛).
// Author: Morteza Farrokhnejad
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func quicksort(arr []int, ascending bool) {
	if len(arr) <= 1 {
		return
	}

	pivotIndex := partition(arr, ascending)
	quicksort(arr[:pivotIndex], ascending)
	quicksort(arr[pivotIndex+1:], ascending)
}

func partition(arr []int, ascending bool) int {
	pivot := arr[len(arr)-1]
	i := 0

	for j := 0; j < len(arr)-1; j++ {
		if (ascending && arr[j] <= pivot) || (!ascending && arr[j] >= pivot) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the number of elements in the array:")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < n; i++ {
		scanner.Scan()
		element, _ := strconv.Atoi(scanner.Text())
		arr[i] = element
	}

	fmt.Println("Original array:", arr)

	// Sort in ascending order
	ascending := make([]int, len(arr))
	copy(ascending, arr)
	quicksort(ascending, true)
	fmt.Println("Sorted in ascending order:", ascending)

	// Sort in descending order
	descending := make([]int, len(arr))
	copy(descending, arr)
	quicksort(descending, false)
	fmt.Println("Sorted in descending order:", descending)
}