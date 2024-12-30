//The following program separates odd and even integers from a given array into two separate arrays.
// The program will take a specified number of integer inputs, store them in an array, and then create and display two new arrays: one containing all even elements and the other containing all odd elements.

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
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < n; i++ {
		fmt.Printf("element - %d: ", i+1)
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	evenArr := []int{}
	oddArr := []int{}
	for _, num := range arr {
		if num%2 == 0 {
			evenArr = append(evenArr, num)
		} else {
			oddArr = append(oddArr, num)
		}
	}

	fmt.Println("\nThe even elements are:")
	for _, even := range evenArr {
		fmt.Printf("%d ", even)
	}
	fmt.Println()

	fmt.Println("\nThe odd elements are:")
	for _, odd := range oddArr {
		fmt.Printf("%d ", odd)
	}
}
