// The program below multiplies two matrices as long as multiplication is possible.
// The program prompts the user to input the size of the matrices, elements for each matrix, and then calculates the result of multiplication.
// Then it displays the original matrices and the result of multiplication as output.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the number of rows for Matrix A:")
	scanner.Scan()
	rowsA, err := strconv.Atoi(scanner.Text())
	if err != nil || rowsA <= 0 {
		fmt.Println("Please enter a valid positive integer for rows.")
		return
	}

	fmt.Println("Enter the number of columns for Matrix A (and rows for Matrix B):")
	scanner.Scan()
	colsA, err := strconv.Atoi(scanner.Text())
	if err != nil || colsA <= 0 {
		fmt.Println("Please enter a valid positive integer for columns.")
		return
	}

	fmt.Println("Enter the number of columns for Matrix B:")
	scanner.Scan()
	colsB, err := strconv.Atoi(scanner.Text())
	if err != nil || colsB <= 0 {
		fmt.Println("Please enter a valid positive integer for columns.")
		return
	}

	matrixA := make([][]int, rowsA)
	matrixB := make([][]int, colsA)
	resultMatrix := make([][]int, rowsA)

	for i := 0; i < rowsA; i++ {
		matrixA[i] = make([]int, colsA)
		resultMatrix[i] = make([]int, colsB)
	}

	for i := 0; i < colsA; i++ {
		matrixB[i] = make([]int, colsB)
	}

	fmt.Printf("Input elements for Matrix A (%dx%d):\n", rowsA, colsA)
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsA; j++ {
			fmt.Printf("A[%d][%d]: ", i+1, j+1)
			scanner.Scan()
			matrixA[i][j], err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Please enter a valid integer.")
				return
			}
		}
	}

	fmt.Printf("Input elements for Matrix B (%dx%d):\n", colsA, colsB)
	for i := 0; i < colsA; i++ {
		for j := 0; j < colsB; j++ {
			fmt.Printf("B[%d][%d]: ", i+1, j+1)
			scanner.Scan()
			matrixB[i][j], err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Please enter a valid integer.")
				return
			}
		}
	}

	for i := 0; i < rowsA; i++ { 	// Multiply matrices
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				resultMatrix[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}

	fmt.Println("\nMatrix A:")
	displayMatrix(matrixA)

	fmt.Println("\nMatrix B:")
	displayMatrix(matrixB)

	fmt.Println("\nResultant Matrix (A x B):")
	displayMatrix(resultMatrix)
}

func displayMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%5d", val)
		}
		fmt.Println()
	}
}
// Author: Morteza Farrokhnejad