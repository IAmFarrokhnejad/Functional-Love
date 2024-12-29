// The program below counts the number of duplicate elements in an array. 
// The program takes a specified number of integer inputs, stores them in an array, and then determines and displays how many elements appear more than once.

use std::collections::HashSet;
use std::io;

fn main() {
    let mut input = String::new();

    println!("Enter the number of elements to be stored in the array:");
    io::stdin().read_line(&mut input).expect("Failed to read input");
    let n: usize = input.trim().parse().expect("Please enter a valid number");
    input.clear();

    let mut arr = Vec::new();

    println!("Input {} elements in the array:", n);
    for i in 0..n {
        println!("element - {}: ", i + 1);
        io::stdin().read_line(&mut input).expect("Failed to read input");
        let value: i32 = input.trim().parse().expect("Please enter a valid integer");
        arr.push(value);
        input.clear();
    }

    let mut seen = HashSet::new();
    let mut duplicates = HashSet::new();

    for &value in &arr {
        if !seen.insert(value) { // Try to insert the value into 'seen'. If it's already there, it means it's a duplicate.
            duplicates.insert(value); // Add the duplicate value to the 'duplicates' set
        }
    }

    println!("Total number of duplicates found in the array: {}", duplicates.len());
}