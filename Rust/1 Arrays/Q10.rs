// The program below inserts a new value into an already sorted array while maintaining the sorted order. 
// The program prompts the user with the number of elements to input, elements in ascending order, and the value to be inserted. 
//It then display the array before and after insertion.
use std::io;

fn main() {
    const MAX_SIZE: usize = 100;
    let mut arr: Vec<i32> = Vec::new();
    let mut input = String::new();

    println!("Input number of elements you want to insert (max {}): ", MAX_SIZE);
    io::stdin().read_line(&mut input).expect("Failed to read input");
    let n: usize = input.trim().parse().expect("Please enter a valid number");

    if n > MAX_SIZE {
        println!("The size of the array cannot exceed {}. Please try again.", MAX_SIZE);
        return;
    }

    println!("Input {} elements in the array in ascending order:", n);
    for i in 0..n {
        input.clear();
        println!("element - {}: ", i);
        io::stdin().read_line(&mut input).expect("Failed to read input");
        let value: i32 = input.trim().parse().expect("Please enter a valid integer");
        arr.push(value);
    }

    input.clear();
    println!("Input the value to be inserted: ");
    io::stdin().read_line(&mut input).expect("Failed to read input");
    let inval: i32 = input.trim().parse().expect("Please enter a valid integer");

    println!("\nThe existing array list is:");
    for value in &arr {
        print!("{:5}", value);
    }

    // Find the position to insert the value
    let position = arr.iter().position(|&x| inval < x).unwrap_or(arr.len());
    
    // Insert the value at the correct position
    arr.insert(position, inval);

    println!("\n\nAfter Insert the list is:");
    for value in &arr {
        print!("{:5}", value);
    }
    println!();
}