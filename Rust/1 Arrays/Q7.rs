// The program below takes a specified number of integer inputs, stores them in an array, and then determines and prints the highest and lowest values among the elements.
// Author: Morteza Farrokhnejad

use std::io;

fn main() {
    let mut input = String::new();

    println!("Enter the number of elements:");
    io::stdin().read_line(&mut input).expect("Failed to read input");
    let n: usize = input.trim().parse().expect("Please enter a valid number");
    input.clear();

    if n == 0 {
        println!("No elements to process.");
        return;
    }

    let mut numbers = Vec::new();
    println!("Enter the elements of the array:");
    for i in 0..n {
        println!("element - {}:", i + 1);
        io::stdin().read_line(&mut input).expect("Failed to read input");
        let element: i32 = input.trim().parse().expect("Please enter a valid integer");
        numbers.push(element);
        input.clear();
    }

    let min = numbers.iter().min().unwrap();
    let max = numbers.iter().max().unwrap();

    println!("\nThe highest value is: {}", max);
    println!("The lowest value is: {}", min);
}
