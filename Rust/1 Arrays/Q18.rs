// The program below finds two array elements whose sum is closest to zero.


fn find_min_sum_pair(arr: &[i32]) -> Option<(i32, i32)> {
    // Check if the array is valid and has at least two elements
    if arr.len() < 2 {
        return None;
    }

    let mut min1_pair = arr[0];
    let mut min2_pair = arr[1];
    let mut min_sum = min1_pair + min2_pair;


    for i in 0..arr.len() - 1 {     // Find the pair of elements with the minimum sum closest to zero
        for j in i + 1..arr.len() {
            let sum = arr[i] + arr[j];
            if sum.abs() < min_sum.abs() {
                min_sum = sum;
                min1_pair = arr[i];
                min2_pair = arr[j];
            }
        }
    }

    Some((min1_pair, min2_pair))
}

fn main() {
    let arr = vec![9, 7, 22, -10, -2, 19, 84, -77, 4, -1];

    println!("The original array: {:?}", arr);

    match find_min_sum_pair(&arr) {
        Some((a, b)) => {
            println!("The pair of elements whose sum is minimum are: {} and {}", a, b);
        }
        None => {
            println!("The array does not contain enough elements to find a pair.");
        }
    }
}