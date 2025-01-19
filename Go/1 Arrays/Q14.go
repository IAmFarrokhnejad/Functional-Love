// The program below adds two square matrices of the same size.
// The program prompts the user to input the size of the matrices (less than 5), elements for each matrix, and then calculates the sum of the matrices. It displays the original matrices and their sum as output.
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

	fmt.Println("Input the size of the square matrices (less than 5):")
	scanner.Scan()
	size, err := strconv.Atoi(scanner.Text())
	if err != nil || size <= 0 || size >= 5 {
		fmt.Println("Please enter a valid size (1 to 4).")
		return
	}

	matrixA := make([][]int, size)
	matrixB := make([][]int, size)
	sumMatrix := make([][]int, size)

	for i := 0; i < size; i++ {
		matrixA[i] = make([]int, size)
		matrixB[i] = make([]int, size)
		sumMatrix[i] = make([]int, size)
	}

	fmt.Printf("Input elements for Matrix A (%dx%d):\n", size, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("A[%d][%d]: ", i, j)
			scanner.Scan()
			matrixA[i][j], err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Please enter a valid integer.")
				return
			}
		}
	}

	fmt.Printf("Input elements for Matrix B (%dx%d):\n", size, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("B[%d][%d]: ", i, j)
			scanner.Scan()
			matrixB[i][j], err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Please enter a valid integer.")
				return
			}
		}
	}

	for i := 0; i < size; i++ { 	// Calculate the sum of the matrices
		for j := 0; j < size; j++ {
			sumMatrix[i][j] = matrixA[i][j] + matrixB[i][j]
		}
	}

	fmt.Println("\nMatrix A:")
	displayMatrix(matrixA)

	fmt.Println("\nMatrix B:")
	displayMatrix(matrixB)

	// Display the Sum Matrix
	fmt.Println("\nSum Matrix:")
	displayMatrix(sumMatrix)
}

func displayMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%5d", val)
		}
		fmt.Println()
	}
}