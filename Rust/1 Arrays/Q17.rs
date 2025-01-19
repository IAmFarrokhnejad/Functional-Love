//The program below calculates the expected sum of the sequence up to the length of the array, then subtract the sum of the array's elements to find the missing number. 
// Author: Morteza Farrokhnejad

fn findMissingNum(arr: &[i32], arr_size: usize) -> i32 {
    // Function to find the missing number
    let n = arr_size as i32 + 1; // Total elements (including the missing number)
    let sum: i32 = arr.iter().sum(); // Sum of all elements in the array

    (n * (n + 1)) / 2 - sum     // Calculate the missing number
}

fn main() {
    let arr = vec![1, 3, 4, 7, 5, 2, 9, 8];
    let arr_size = arr.len();

    println!("The given array is: {:?}", arr);

    let missing_number = findMissingNum(&arr, arr_size);
    println!("The missing number is: {}", missing_number);
}