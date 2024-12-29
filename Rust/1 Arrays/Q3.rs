// The program below involves copying the elements from one array to another. 
// The program takes a specified number of integer inputs to store in the first array, then copies these elements to a second array, and finally displays the contents of both arrays.

use std::io;

fn main() {
    let mut input = String::new();

    println!("Enter the number of elements to be stored in the array:");
    io::stdin().read_line(&mut input).expect("Failed to read input");
    let n: usize = input.trim().parse().expect("Please enter a valid number");
    input.clear();

    let mut arr1 = Vec::new();

    println!("Input {} elements in the array:", n);
    for i in 0..n {
        println!("element - {}: ", i + 1);
        io::stdin().read_line(&mut input).expect("Failed to read input");
        let value: i32 = input.trim().parse().expect("Please enter a valid integer");
        arr1.push(value);
        input.clear();
    }

    let arr2 = arr1.clone();

    println!("\nThe elements stored in the first array are:");
    for value in &arr1 {
        print!("{:5}", value);
    }
    println!();

    println!("\nThe elements copied into the second array are:");
    for value in &arr2 {
        print!("{:5}", value);
    }
}