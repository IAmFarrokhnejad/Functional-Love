// The program below prompts the user to input the size of the matrix and its elements.
// It then calculates the sum of each row and column, and displays the original matrix followed by the sums of its rows and columns as output.
// Author: Morteza Farrokhnejad
use std::io;

fn main() {
    println!("Input the size of the square matrix:");
    let n = read_usize();

    println!("Input elements in the matrix:");
    let mut matrix = vec![vec![0; n]; n];
    for i in 0..n {
        for j in 0..n {
            println!("element - [{}],[{}]:", i, j);
            matrix[i][j] = read_i32();
        }
    }

    println!("\nThe original matrix:");
    display_matrix(&matrix);

    let row_sums: Vec<i32> = matrix.iter().map(|row| row.iter().sum()).collect();     // Sum of rows

    let mut col_sums = vec![0; n];     // Sum of columns
    for j in 0..n {
        col_sums[j] = matrix.iter().map(|row| row[j]).sum();
    }

    println!("\nThe sum of rows and columns of the matrix is:");
    for (i, row) in matrix.iter().enumerate() {
        for &val in row {
            print!("{:4}", val);
        }
        println!("{:7}", row_sums[i]);
    }

    println!("\nColumn sums:");
    for &sum in &col_sums {
        print!("{:4}", sum);
    }
    println!();
}

fn read_usize() -> usize { // Function to read a usize
    let mut input = String::new();
    io::stdin().read_line(&mut input).unwrap();
    input.trim().parse().unwrap()
}

fn read_i32() -> i32 { // Function to read an i32
    let mut input = String::new();
    io::stdin().read_line(&mut input).unwrap();
    input.trim().parse().unwrap()
}

fn display_matrix(matrix: &[Vec<i32>]) {
    for row in matrix {
        for &val in row {
            print!("{:4}", val);
        }
        println!();
    }
}