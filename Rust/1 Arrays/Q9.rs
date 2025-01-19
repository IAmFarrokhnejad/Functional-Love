// The program below sorts an array in both ascending and descending order using the Quicksort algorithm.
// Quicksort is one of the most efficient sorting algorithms with an average time complexity of 𝑂(𝑛 log 𝑛).
// Author: Morteza Farrokhnejad
use std::io;

fn quicksort(arr: &mut [i32], ascending: bool) {
    if arr.len() <= 1 {
        return;
    }

    let pivot_index = partition(arr, ascending);
    quicksort(&mut arr[0..pivot_index], ascending);
    quicksort(&mut arr[pivot_index + 1..], ascending);
}

fn partition(arr: &mut [i32], ascending: bool) -> usize {
    let pivot = arr[arr.len() - 1];
    let mut i = 0;

    for j in 0..arr.len() - 1 {
        if (ascending && arr[j] <= pivot) || (!ascending && arr[j] >= pivot) {
            arr.swap(i, j);
            i += 1;
        }
    }
    arr.swap(i, arr.len() - 1);
    i
}

fn main() {
    let mut input = String::new();
    println!("Enter the number of elements in the array:");
    io::stdin().read_line(&mut input).expect("Failed to read input");
    let n: usize = input.trim().parse().expect("Please enter a valid number");

    let mut arr = Vec::new();

    println!("Enter the elements of the array:");
    for _ in 0..n {
        input.clear();
        io::stdin().read_line(&mut input).expect("Failed to read input");
        let num: i32 = input.trim().parse().expect("Please enter a valid integer");
        arr.push(num);
    }

    println!("Original array: {:?}", arr);

    // Sort in ascending order
    let mut ascending = arr.clone();
    quicksort(&mut ascending, true);
    println!("Sorted in ascending order: {:?}", ascending);

    // Sort in descending order
    let mut descending = arr.clone();
    quicksort(&mut descending, false);
    println!("Sorted in descending order: {:?}", descending);
}