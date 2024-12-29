// In this program, we will enter two sorted arrays. While merging them, we will compare the elements of both the arrays and merge them in a sorted manner.


use std::io;

fn main() {
    let mut input = String::new();

    println!("Enter the size of the first array:");
    io::stdin().read_line(&mut input).expect("Failed to read input");
    let n1: usize = input.trim().parse().expect("Please enter a valid number");
    input.clear();

    let mut a = Vec::new();
    println!("Enter the elements of the first array:");
    for _ in 0..n1 {
        io::stdin().read_line(&mut input).expect("Failed to read input");
        let element: i32 = input.trim().parse().expect("Please enter a valid integer");
        a.push(element);
        input.clear();
    }

    println!("Enter the size of the second array:");
    io::stdin().read_line(&mut input).expect("Failed to read input");
    let n2: usize = input.trim().parse().expect("Please enter a valid number");
    input.clear();

    let mut b = Vec::new();
    println!("Enter the elements of the second array:");
    for _ in 0..n2 {
        io::stdin().read_line(&mut input).expect("Failed to read input");
        let element: i32 = input.trim().parse().expect("Please enter a valid integer");
        b.push(element);
        input.clear();
    }

    let mut c = Vec::new();
    let mut i = 0;
    let mut j = 0;

    while i < n1 && j < n2 {
        if a[i] < b[j] {
            c.push(a[i]);
            i += 1;
        } else {
            c.push(b[j]);
            j += 1;
        }
    }

    while i < n1 {
        c.push(a[i]);
        i += 1;
    }

    while j < n2 {
        c.push(b[j]);
        j += 1;
    }

    println!("Final array after merging:");
    for element in &c {
        print!("{} ", element);
    }
    println!();
}