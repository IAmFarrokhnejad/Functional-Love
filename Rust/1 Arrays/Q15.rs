// The program below multiplies two matrices as long as multiplication is possible. 
//The program prompts the user to input the size of the matrices, elements for each matrix, and then calculates the result of multiplication.
// Then it displays the original matrices and the result of multiplication as output.
use std::io;

fn main() {
    println!("Enter the number of rows for the first matrix:");
    let rows1 = read_usize();
    println!("Enter the number of columns for the first matrix:");
    let cols1 = read_usize();

    println!("Enter the number of rows for the second matrix:");
    let rows2 = read_usize();
    println!("Enter the number of columns for the second matrix:");
    let cols2 = read_usize();

    if cols1 != rows2 {     // Check if matrix multiplication is possible
        println!("Matrix multiplication is not possible. The number of columns of the first matrix must equal the number of rows of the second matrix.");
        return;
    }

    println!("Enter the elements of the first {}x{} matrix:", rows1, cols1);
    let matrix1 = read_matrix(rows1, cols1);
    println!("Enter the elements of the second {}x{} matrix:", rows2, cols2);
    let matrix2 = read_matrix(rows2, cols2);

    let matrix_product = multiply_matrices(&matrix1, &matrix2, rows1, cols1, cols2);

    println!("\nFirst Matrix:");
    display_matrix(&matrix1);

    println!("\nSecond Matrix:");
    display_matrix(&matrix2);

    println!("\nResultant Matrix after Multiplication:");
    display_matrix(&matrix_product);
}

fn read_usize() -> usize {
    let mut input = String::new();
    io::stdin().read_line(&mut input).unwrap();
    input.trim().parse().unwrap()
}

fn read_matrix(rows: usize, cols: usize) -> Vec<Vec<i32>> {
    let mut matrix = vec![vec![0; cols]; rows];
    for i in 0..rows {
        for j in 0..cols {
            println!("Element [{}][{}]:", i + 1, j + 1);
            let mut input = String::new();
            io::stdin().read_line(&mut input).unwrap();
            matrix[i][j] = input.trim().parse().unwrap();
        }
    }
    matrix
}

fn multiply_matrices(
    matrix1: &Vec<Vec<i32>>,
    matrix2: &Vec<Vec<i32>>,
    rows1: usize,
    cols1: usize,
    cols2: usize,
) -> Vec<Vec<i32>> {
    let mut product = vec![vec![0; cols2]; rows1];
    for i in 0..rows1 {
        for j in 0..cols2 {
            product[i][j] = (0..cols1).map(|k| matrix1[i][k] * matrix2[k][j]).sum();
        }
    }
    product
}

fn display_matrix(matrix: &Vec<Vec<i32>>) {
    for row in matrix {
        for &element in row {
            print!("{:5}", element);
        }
        println!();
    }
}
// Author: Morteza Farrokhnejad