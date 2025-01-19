// The program below sorts an array of 0s, 1s and 2s using a three-way partitioning approach.
// Author: Morteza Farrokhnejad

fn swap(arr: &mut Vec<i32>, i: usize, j: usize) {
    let temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
}

fn sort_elements(arr: &mut Vec<i32>) {
    let mut start = 0;
    let mut mid = 0;
    let mut end = arr.len() - 1;
    let pivot = 1;

    while mid <= end {
        if arr[mid] < pivot {
            swap(arr, start, mid);
            start += 1;
            mid += 1;
        } else if arr[mid] > pivot {
            swap(arr, mid, end);
            if end > 0 { // To avoid underflow
                end -= 1;
            }
        } else {
            mid += 1;
        }
    }
}

fn main() {
    let mut arr = vec![0, 1, 2, 2, 1, 0, 0, 2, 0, 1, 1, 0, 2, 1, 2, 2, 0, 2];

    println!("The original array: {:?}", arr);

    sort_elements(&mut arr);

    println!("After sorting, the elements in the array are:");
    println!("{:?}", arr);
}