// The program below moves all zeroes to the end of a given array. 


fn move_zeros_to_end(arr: &mut Vec<i32>) {
    let mut left = 0;
    let mut right = arr.len() - 1;

    while left < right {
        // Move the left pointer until a zero is found
        while left < arr.len() && arr[left] != 0 {
            left += 1;
        }

        // Move the right pointer until a non-zero is found
        while right > 0 && arr[right] == 0 {
            right -= 1;
        }

        // Swap the elements if the pointers haven't crossed
        if left < right {
            arr.swap(left, right);
        }
    }
}

fn main() {
    let mut arr = vec![24, 0, 7, 0, 4, 0, 7, -5, 8, 2];

    println!("The original array is: {:?}", arr);

    move_zeros_to_end(&mut arr);

    println!("The new array is: {:?}", arr);
}