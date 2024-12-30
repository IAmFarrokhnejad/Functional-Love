// The program below adds two square matrices of the same size. 
//The program prompts the user to input the size of the matrices (less than 5), elements for each matrix, and then calculates the sum of the matrices. It displays the original matrices and their sum as output.

use std::io;

fn main() {
    const MAX_SIZE: usize = 5;

    println!("Enter the size of the square matrices (less than {}):", MAX_SIZE);
    let mut input = String::new();
    io::stdin().read_line(&mut input).unwrap();
    let n: usize = input.trim().parse().unwrap();

    if n >= MAX_SIZE {
        println!("The size must be less than {}.", MAX_SIZE);
        return;
    }

    let mut matrix1 = vec![vec![0; n]; n];
    let mut matrix2 = vec![vec![0; n]; n];
    let mut matrix_sum = vec![vec![0; n]; n];

    println!("Enter the elements of the first {}x{} matrix:", n, n);
    for i in 0..n {
        for j in 0..n {
            println!("Element [{}][{}]:", i , j);
            input.clear();
            io::stdin().read_line(&mut input).unwrap();
            matrix1[i][j] = input.trim().parse().unwrap();
        }
    }

    println!("Enter the elements of the second {}x{} matrix:", n, n);
    for i in 0..n {
        for j in 0..n {
            println!("Element [{}][{}]:", i, j);
            input.clear();
            io::stdin().read_line(&mut input).unwrap();
            matrix2[i][j] = input.trim().parse().unwrap();
        }
    }

    for i in 0..n {     // Calculate the sum of the matrices
        for j in 0..n {
            matrix_sum[i][j] = matrix1[i][j] + matrix2[i][j];
        }
    }

    println!("\nMatrix 1:");
    display_matrix(&matrix1);

    println!("\nMatrix 2:");
    display_matrix(&matrix2);

    println!("\nSum of Matrix 1 and Matrix 2:");
    display_matrix(&matrix_sum);
}

fn display_matrix(matrix: &Vec<Vec<i32>>) {
    for row in matrix {
        for &element in row {
            print!("{:5}", element);
        }
        println!();
    }
}