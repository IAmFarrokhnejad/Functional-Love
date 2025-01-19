// The program below reads a specified number of integers into an array and then calculate and print the sum of these elements. 
// After storing the input values, the program iterates through the array to compute the total sum and display the result.
// Author: Morteza Farrokhnejad
use std::io;

fn main() {
    let mut input = String::new();
    let mut sum = 0;

    println!("Enter the number of elements to be stored in the array:");
    io::stdin().read_line(&mut input).expect("Failed to read input");
    let n: usize = input.trim().parse().expect("Please enter a valid number");
    input.clear();

    let mut array = Vec::new();

    println!("Input {} elements in the array:", n);
    for i in 0..n {
        println!("element - {}: ", i + 1);
        io::stdin().read_line(&mut input).expect("Failed to read input");
        let value: i32 = input.trim().parse().expect("Please enter a valid integer");
        array.push(value);
        input.clear();
    }

    for &value in &array {
        sum += value;
    }

    println!("Sum of all elements stored in the array is: {}\n", sum);
}