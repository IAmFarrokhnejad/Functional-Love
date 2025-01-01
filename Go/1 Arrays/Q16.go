// The program below prompts the user to input the size of the matrix and its elements.
// It then calculates the sum of each row and column, and displays the original matrix followed by the sums of its rows and columns as output.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Input the size of the square matrix:")
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil || n <= 0 {
		fmt.Println("Please enter a valid positive integer for the size of the matrix.")
		return
	}

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	fmt.Println("Input elements in the matrix:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("element - [%d],[%d]: ", i, j)
			scanner.Scan()
			matrix[i][j], err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Please enter a valid integer.")
				return
			}
		}
	}

	fmt.Println("\nThe original matrix:")
	displayMatrix(matrix)

	rowSums := make([]int, n) 	// Sum of rows
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rowSums[i] += matrix[i][j]
		}
	}

	colSums := make([]int, n) 	// Sum of columns
	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			colSums[j] += matrix[i][j]
		}
	}

	fmt.Println("\nThe sum of rows and columns of the matrix is:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%4d", matrix[i][j])
		}
		fmt.Printf("%7d\n", rowSums[i])
	}

	fmt.Println("\nColumn sums:")
	for _, sum := range colSums {
		fmt.Printf("%4d", sum)
	}
	fmt.Println()
}

func displayMatrix(matrix [][]int) { // Function to display the matrix
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%4d", val)
		}
		fmt.Println()
	}
}