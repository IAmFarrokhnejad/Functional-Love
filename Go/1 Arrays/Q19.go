// The program below prints a matrix in spiral form.
// Author: Morteza Farrokhnejad

package main

import "fmt"

func spiralOfMatrix(matrix [][]int) {
	rows := len(matrix)
	if rows == 0 {
		return
	}
	cols := len(matrix[0])
	if cols == 0 {
		return
	}

	rowStart, rowEnd := 0, rows
	colStart, colEnd := 0, cols

	for rowStart < rowEnd && colStart < colEnd {
		// Print the first row
		for col := colStart; col < colEnd; col++ {
			fmt.Printf("%d ", matrix[rowStart][col])
		}
		rowStart++

		// Print the last column
		for row := rowStart; row < rowEnd; row++ {
			fmt.Printf("%d ", matrix[row][colEnd-1])
		}
		colEnd--

		// Print the last row if it's within bounds
		if rowStart < rowEnd {
			for col := colEnd - 1; col >= colStart; col-- {
				fmt.Printf("%d ", matrix[rowEnd-1][col])
			}
			rowEnd--
		}

		// Print the first column if it's within bounds
		if colStart < colEnd {
			for row := rowEnd - 1; row >= rowStart; row-- {
				fmt.Printf("%d ", matrix[row][colStart])
			}
			colStart++
		}
	}
	fmt.Println()
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}

	fmt.Println("The original matrix is:")
	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("The spiral form of the above matrix is:")
	spiralOfMatrix(matrix)
}