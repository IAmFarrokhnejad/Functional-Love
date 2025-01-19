// The program below counts and displays the frequency of each element in an array. 
// The program accepts a specified number of integer inputs, stores them in an array, and then determines and prints how many times each element appears in the array.
// Author: Morteza Farrokhnejad
use std::collections::HashMap;
use std::io;

fn main() {
    let mut input = String::new();

    // Read the number of elements
    println!("Enter the number of elements to be stored in the array:");
    io::stdin().read_line(&mut input).expect("Failed to read input");
    let n: usize = input.trim().parse().expect("Please enter a valid number");
    input.clear();

    let mut arr = Vec::new();
    println!("Enter the elements of the array:");
    for i in 0..n {
        println!("element - {}:", i + 1);
        io::stdin().read_line(&mut input).expect("Failed to read input");
        let element: i32 = input.trim().parse().expect("Please enter a valid integer");
        arr.push(element);
        input.clear();
    }

    // Count the frequency of each element using a HashMap
    let mut frequency = HashMap::new();
    for &element in &arr {
        *frequency.entry(element).or_insert(0) += 1;
    }

    println!("\nThe frequency of all elements of the array:");
    for (element, count) in &frequency {
        println!("{} occurs {} times", element, count);
    }
}   