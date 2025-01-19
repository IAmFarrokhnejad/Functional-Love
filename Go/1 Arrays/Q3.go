// The program below involves copying the elements from one array to another.
// The program takes a specified number of integer inputs to store in the first array, then copies these elements to a second array, and finally displays the contents of both arrays.
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
	fmt.Println("Enter the number of elements to be stored in the array:")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr1 := make([]int, n)

	fmt.Printf("Input %d elements in the array:\n", n)
	for i := 0; i < n; i++ {
		fmt.Printf("element - %d: ", i+1)
		scanner.Scan()
		arr1[i], _ = strconv.Atoi(scanner.Text())
	}

	arr2 := make([]int, n) // Copying elements to another array
	copy(arr2, arr1)

	fmt.Println("\nThe elements stored in the first array are:")
	for _, value := range arr1 {
		fmt.Printf("%5d", value)
	}
	fmt.Println()

	fmt.Println("\nThe elements copied into the second array are:")
	for _, value := range arr2 {
		fmt.Printf("%5d", value)
	}
}