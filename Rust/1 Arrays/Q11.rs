// The program below insert values in the array (unsorted list). It promots the user to enter an integer value and a value to be inserted and its position.
// The program will then display the array after inserting the new value at the given position.
// Author: Morteza Farrokhnejad
use std::io;

fn main() {
    const MAX_SIZE: usize = 100;
    let mut arr: Vec<i32> = Vec::new();
    let mut input = String::new();

    println!("Input the size of array (max {}): ", MAX_SIZE);
    io::stdin().read_line(&mut input).expect("Failed to read input");
    let n: usize = input.trim().parse().expect("Please enter a valid number");

    if n > MAX_SIZE {
        println!("The size of the array cannot exceed {}. Please try again.", MAX_SIZE);
        return;
    }
    
    println!("Input {} elements in the array:", n);
    for i in 0..n {
        input.clear();
        println!("element - {}: ", i + 1);
        io::stdin().read_line(&mut input).expect("Failed to read input");
        let value: i32 = input.trim().parse().expect("Please enter a valid integer");
        arr.push(value);
    }
    input.clear();
    println!("Input the value to be inserted: ");
    io::stdin().read_line(&mut input).expect("Failed to read input");
    let x: i32 = input.trim().parse().expect("Please enter a valid integer");

    input.clear();     // Input the position where the value should be inserted
    println!("Input the position where the value should be inserted: ");
    io::stdin().read_line(&mut input).expect("Failed to read input");
    let p: usize = input.trim().parse().expect("Please enter a valid position");

    if p == 0 || p > n + 1 {
        println!("Invalid position! Position should be between 1 and {}.", n + 1);
        return;
    }

    println!("\nThe current list of the array:");
    for value in &arr {
        print!("{:5}", value);
    }

    arr.insert(p - 1, x); // Insert the value at the specified position

    println!("\n\nAfter inserting the element, the new list is:");
    for value in &arr {
        print!("{:5}", value);
    }
    println!();
}