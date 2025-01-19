// The program below checks whether one array is a subset of another array (all elements of the second array are present in the first array).
// It does it by converting the first array into a HashSet to allow efficient lookups. 
// It uses the all iterator method to check if every element in the second array exists in the HashSet derived from the first array.
use std::collections::HashSet;

fn checkIfSubset(arr1: &[i32], arr2: &[i32]) -> bool {
    let set1: HashSet<i32> = arr1.iter().cloned().collect();
    arr2.iter().all(|&item| set1.contains(&item))
}

fn main() {
    let arr1 = [1, 8, 2, 4, 6, 9, 5, 0, 2];
    let arr2 = [5, 4, 1, 0, 6, 2];

    println!("The first array is: {:?}", arr1);
    println!("The second array is: {:?}", arr2);

    if checkIfSubset(&arr1, &arr2) {
        println!("The second array is a subset of the first array.");
        } else {
            println!("The second array is not a subset of the first array.");
        } 
}
// Author: Morteza Farrokhnejad