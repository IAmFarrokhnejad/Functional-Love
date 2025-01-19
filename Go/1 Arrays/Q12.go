// The program below deletes an element at a desired position from an array.
// Author: Morteza Farrokhnejad

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Input the size of the array: ")
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil || n < 1 {
		fmt.Println("Please enter a valid positive integer.")
		return
	}

	arr := make([]int, n)
	fmt.Printf("Input %d elements in the array:\n", n)
	for i := 0; i < n; i++ {
		fmt.Printf("element - %d: ", i+1)
		scanner.Scan()
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Please enter a valid integer.")
			return
		}
		arr[i] = value
	}

	fmt.Println("Input the position where to delete: ")
	scanner.Scan()
	pos, err := strconv.Atoi(scanner.Text())
	if err != nil || pos < 1 || pos > n {
		fmt.Printf("Invalid position! Position should be between 1 and %d.\n", n)
		return
	}


	arr = append(arr[:pos-1], arr[pos:]...) 	// Remove the element at the specified position

	fmt.Println("\nThe new list is:")
	for _, value := range arr {
		fmt.Printf("%5d", value)
	}
	fmt.Println()
}