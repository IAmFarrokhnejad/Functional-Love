// This program reads a user-defined number of integer values into an array and then display these values in reverse order. 
// After storing the values, the program first prints them in the original order and then prints them in the reversed order.
// Author: Morteza Farrokhnejad
use std::io;

fn main() {
    let mut input = String::new();
    
    println!("Enter the number of elements to store in the array:");
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

    println!("\nThe values stored in the array are:");
    for value in &array {
        print!("{:5}", value);
    }
    println!();

    println!("\nThe values stored in the array in reverse are:");
    for value in array.iter().rev() {
        print!("{:5}", value);
    }
}