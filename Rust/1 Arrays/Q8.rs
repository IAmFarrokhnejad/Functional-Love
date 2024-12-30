//The following program separates odd and even integers from a given array into two separate arrays. 
// The program will take a specified number of integer inputs, store them in an array, and then create and display two new arrays: one containing all even elements and the other containing all odd elements.

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

    let mut evens = Vec::new();
    let mut odds = Vec::new();

    for &num in &numbers {
        if num % 2 == 0 {
            evens.push(num);
        } else {
            odds.push(num);
        }
    }

    println!("\nOriginal array: {:?}", numbers);
    println!("Even elements: {:?}", evens);
    println!("Odd elements: {:?}", odds);
}
