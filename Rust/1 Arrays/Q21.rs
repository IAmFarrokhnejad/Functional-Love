// The program below checks whether an array is a subset of another array.


fn is_subset(arr1: &[i32], arr2: &[i32]) -> bool {
    for &elem in arr2 {
        if !arr1.contains(&elem) {
            return false;
        }
    }
    true
}

fn main() {
    let arr1 = vec![4, 8, 2, 4, 6, 9, 5, 0, 2];
    let arr2 = vec![5, 4, 2, 0, 6, 2];

    println!("The original array is: {:?}", arr1);
    println!("The second array is: {:?}", arr2);

    if is_subset(&arr1, &arr2) {
        println!("The second array is a subset of the first array.");
    } else {
        println!("The second array is not a subset of the first array.");
    }
}