// The program below insert values in the array (unsorted list). It promots the user to enter an integer value and a value to be inserted and its position.
// The program will then display the array after inserting the new value at the given position.
// Author: Morteza Farrokhnejad
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	const MAX_SIZE = 100
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Input the size of array (max %d): ", MAX_SIZE)
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil || n < 1 || n > MAX_SIZE {
		fmt.Printf("Please enter a valid number between 1 and %d.\n", MAX_SIZE)
		return
	}

	arr := make([]int, 0, n)
	fmt.Printf("Input %d elements in the array:\n", n)
	for i := 0; i < n; i++ {
		fmt.Printf("element - %d: ", i+1)
		scanner.Scan()
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Please enter a valid integer.")
			return
		}
		arr = append(arr, value)
	}

	fmt.Println("Input the value to be inserted: ")
	scanner.Scan()
	x, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Please enter a valid integer.")
		return
	}

	fmt.Println("Input the position where the value should be inserted: ")
	scanner.Scan()
	p, err := strconv.Atoi(scanner.Text())
	if err != nil || p < 1 || p > n+1 {
		fmt.Printf("Invalid position! Position should be between 1 and %d.\n", n+1)
		return
	}

	fmt.Println("\nThe current list of the array:")
	for _, value := range arr {
		fmt.Printf("%5d", value)
	}
	fmt.Println()

	arr = append(arr[:p-1], append([]int{x}, arr[p-1:]...)...)	// Insert the value at the specified position

	fmt.Println("\nAfter inserting the element, the new list is:")
	for _, value := range arr {
		fmt.Printf("%5d", value)
	}
	fmt.Println()
}