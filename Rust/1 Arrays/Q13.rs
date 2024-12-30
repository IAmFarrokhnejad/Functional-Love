// The program below finds the largest, second largest, and the third largest elements of an array without sorting the said array.


use std::io;

fn main() {
    println!("Enter the number of elements in the array:");
    let mut input = String::new();
    io::stdin().read_line(&mut input).unwrap();
    let n: usize = input.trim().parse().unwrap();

    if n < 3 {
        println!("The array must have at least 3 elements.");
        return;
    }

    println!("Enter {} elements separated by spaces:", n);
    input.clear();
    io::stdin().read_line(&mut input).unwrap();
    let arr: Vec<i32> = input.trim().split_whitespace().map(|x| x.parse().unwrap()).collect();

    if arr.len() != n {
        println!("Number of elements entered does not match the expected count.");
        return;
    }

    // Initialize the largest, second largest, and third largest values
    let mut largest = i32::MIN;
    let mut second_largest = i32::MIN;
    let mut third_largest = i32::MIN;

    for &num in &arr {
        if num > largest {
            third_largest = second_largest;
            second_largest = largest;
            largest = num;
        } else if num > second_largest && num < largest {
            third_largest = second_largest;
            second_largest = num;
        } else if num > third_largest && num < second_largest {
            third_largest = num;
        }
    }

    // Check if distinct second and third largest elements exist
    if second_largest == i32::MIN || third_largest == i32::MIN {
        println!("The array doesn't have distinct second and third largest elements.");
    } else {
        println!("The largest element is: {}", largest);
        println!("The second largest element is: {}", second_largest);
        println!("The third largest element is: {}", third_largest);
    }
}