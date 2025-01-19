// The program below prints a matrix in spiral form.
// Author: Morteza Farrokhnejad

fn spiral_of_matrix(matrix: &[Vec<i32>]) {
    let rows = matrix.len();
    if rows == 0 {
        return;
    }
    let cols = matrix[0].len();
    if cols == 0 {
        return;
    }

    let (mut row_start, mut row_end) = (0, rows);
    let (mut col_start, mut col_end) = (0, cols);

    while row_start < row_end && col_start < col_end {
        // Print the first row
        for col in col_start..col_end {
            print!("{} ", matrix[row_start][col]);
        }
        row_start += 1;

        // Print the last column
        for row in row_start..row_end {
            print!("{} ", matrix[row][col_end - 1]);
        }
        if col_end > col_start {
            col_end -= 1;
        }

        // Print the last row if it's within bounds
        if row_start < row_end {
            for col in (col_start..col_end).rev() {
                print!("{} ", matrix[row_end - 1][col]);
            }
            row_end -= 1;
        }

        // Print the first column if it's within bounds
        if col_start < col_end {
            for row in (row_start..row_end).rev() {
                print!("{} ", matrix[row][col_start]);
            }
            col_start += 1;
        }
    }
    println!();
}

fn main() {
    let matrix = vec![
        vec![1, 2, 3, 4, 5],
        vec![6, 7, 8, 9, 10],
        vec![11, 12, 13, 14, 15],
        vec![16, 17, 18, 19, 20],
    ];

    println!("The original matrix is: ");
    for row in &matrix {
        for value in row {
            print!("{} ", value);
        }
        println!();
    }
    println!();

    println!("The spiral form of the above matrix is:");
    spiral_of_matrix(&matrix);
}