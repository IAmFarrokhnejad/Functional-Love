// The program below deletes an element at a desired position from an array.


use std::io;

fn main() {
    let mut arr = Vec::new();
    let mut input = String::new();

    println!("Input the size of the array: ");
    io::stdin().read_line(&mut input).expect("Failed to read input");
    let n: usize = input.trim().parse().expect("Please enter a valid number");

    println!("Input {} elements in the array:", n);
    for i in 1..=n {
        input.clear();
        println!("element - {}: ", i);
        io::stdin().read_line(&mut input).expect("Failed to read input");
        let value: i32 = input.trim().parse().expect("Please enter a valid integer");
        arr.push(value);
    }

    input.clear(); // Input the position to delete
    println!("Input the position where to delete: ");
    io::stdin().read_line(&mut input).expect("Failed to read input");
    let pos: usize = input.trim().parse().expect("Please enter a valid position");

    if (1..=n).contains(&pos) { // Validate position and remove element
        arr.remove(pos - 1);
        println!("\nThe new list is: {:?}", arr);
    } else {
        println!("Invalid position! Position should be between 1 and {}.", n);
    }
}