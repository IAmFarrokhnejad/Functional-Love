// The program below counts the number of duplicate elements in an array.
// The program takes a specified number of integer inputs, stores them in an array, and then determines and displays how many elements appear more than once.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the number of elements to be stored in the array:")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)
	fmt.Printf("Input %d elements in the array:\n", n)
	for i := 0; i < n; i++ {
		fmt.Printf("element - %d: ", i+1)
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	// Use a map to track occurrences
	seen := make(map[int]bool)
	duplicates := make(map[int]bool)

	for _, value := range arr {
		if seen[value] {
			duplicates[value] = true
		} else {
			seen[value] = true
		}
	}

	fmt.Printf("Total number of duplicates found in the array: %d\n", len(duplicates))
}