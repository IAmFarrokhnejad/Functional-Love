// The program below returns the counting sort on an array.
// Author: Morteza Farrokhnejad

fn countingSort(arr: &mut [usize], max_value: usize) {
    let mut count = vec![0; max_value + 1];

    // Count the occurrences of each value
    for &num in arr.iter() {
        count[num] += 1;
    }

    // Reconstruct the sorted array
    let mut index = 0;
    for (value, &freq) in count.iter().enumerate() {
        for _ in 0..freq {
            arr[index] = value;
            index += 1;
        }
    }
}

fn main() {
    let mut array = [4, 2, 0, 1, 24, 4, 1, 2, 8, 3, 3, 1];
    let max_value = *array.iter().max().unwrap();

    println!("Original array: {:?}", array);

    countingSort(&mut array, max_value);

    println!("Sorted array: {:?}", array);
}