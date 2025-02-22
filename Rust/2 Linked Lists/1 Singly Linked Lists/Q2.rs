// The program below demosntrates implementation of a singly linked list. Following functions are included:

// 1. fn new_node(data: i32): Creates a new node with given data.
// 2. detect_and_remove_loop(head: &Link): Detects and removes loop in a linked list.
// 3. display_list: Prints a linked list.
// 4. display_list_char(head: &Link): Prints a linked list of characters.
// 5. is_palindrome(head: &Link): Checks if the linked list is palindrome.

// Author: Morteza Farrokhnejad

use std::rc::Rc;
use std::cell::RefCell;

type Link = Option<Rc<RefCell<Node>>>;

#[derive(Debug)]
struct Node {
    data: i32,
    next: Link,
}

/// Creates a new node with the given data.
fn new_node(data: i32) -> Rc<RefCell<Node>> {
    Rc::new(RefCell::new(Node { data, next: None }))
}

/// Displays the list as integers (assumes no cycle).
fn display_list(head: &Link) {
    let mut current = head.clone();
    while let Some(node_rc) = current {
        print!("{} ", node_rc.borrow().data);
        current = node_rc.borrow().next.clone();
    }
    println!();
}

/// Displays the list as characters (casting data to u8 then to char).
fn display_list_char(head: &Link) {
    let mut current = head.clone();
    while let Some(node_rc) = current {
        print!("{}", node_rc.borrow().data as u8 as char);
        current = node_rc.borrow().next.clone();
    }
    println!();
}

/// Detects and removes a loop in the linked list using Floyd’s cycle-finding algorithm.
/// All temporary borrows are confined to inner blocks to ensure they are dropped before we reassign.
fn detect_and_remove_loop(head: &Link) {
    let mut slow = head.clone();
    let mut fast = head.clone();
    let mut loop_found = false;

    // Loop detection.
    while let (Some(s), Some(f)) = (slow.clone(), fast.clone()) {
        slow = s.borrow().next.clone();
        fast = if let Some(next) = f.borrow().next.clone() {
            next.borrow().next.clone()
        } else {
            None
        };

        if let (Some(s_ptr), Some(f_ptr)) = (slow.clone(), fast.clone()) {
            if Rc::ptr_eq(&s_ptr, &f_ptr) {
                loop_found = true;
                break;
            }
        }
    }

    // Loop removal.
    if loop_found {
        slow = head.clone();
        while let (Some(s), Some(f)) = (slow.clone(), fast.clone()) {
            let s_next = s.borrow().next.clone();
            let f_next = f.borrow().next.clone();
            if let (Some(s_next_rc), Some(f_next_rc)) = (s_next, f_next) {
                if Rc::ptr_eq(&s_next_rc, &f_next_rc) {
                    // Break the loop. We limit the borrow in its own block.
                    {
                        let mut f_mut = f.borrow_mut();
                        f_mut.next = None;
                    }
                    break;
                }
            }
            // Use block scopes to drop the temporary borrows before assignment.
            slow = { s.borrow().next.clone() };
            fast = { f.borrow().next.clone() };
        }
    }
}

/// Reverses the linked list and returns the new head.
fn reverse_list(mut head: Link) -> Link {
    let mut prev: Link = None;
    while let Some(node_rc) = head {
        let next = node_rc.borrow_mut().next.take();
        node_rc.borrow_mut().next = prev;
        prev = Some(node_rc);
        head = next;
    }
    prev
}

/// Checks if the linked list is a palindrome.
/// This function mimics the C approach:
/// 1. Find the middle of the list using two pointers.
/// 2. For odd-length lists, skip the middle node.
/// 3. Reverse the second half and compare it with the first half.
/// 4. Restore the list before returning.
fn is_palindrome(head: &Link) -> bool {
    if head.is_none() {
        return true;
    }

    let mut slow = head.clone();
    let mut fast = head.clone();
    let mut prev_slow = head.clone();
    let mut midnode: Link = None;

    // Find middle.
    while fast.is_some() && fast.as_ref().unwrap().borrow().next.is_some() {
        fast = {
            let temp = fast.as_ref().unwrap().borrow().next.clone();
            if let Some(ref next) = temp {
                next.borrow().next.clone()
            } else {
                None
            }
        };
        prev_slow = slow.clone();
        slow = slow.unwrap().borrow().next.clone();
    }

    // If odd number of nodes, skip the middle node.
    if fast.is_some() {
        midnode = slow.clone();
        slow = slow.unwrap().borrow().next.clone();
    }

    // Reverse the second half.
    let reversed_second_half = reverse_list(slow);

    // Compare first half and reversed second half.
    let mut p1 = head.clone();
    let mut p2 = reversed_second_half.clone();
    let mut result = true;
    while p2.is_some() {
        if p1.as_ref().unwrap().borrow().data != p2.as_ref().unwrap().borrow().data {
            result = false;
            break;
        }
        p1 = p1.unwrap().borrow().next.clone();
        p2 = p2.unwrap().borrow().next.clone();
    }

    // Restore the original list structure.
    let restored = reverse_list(reversed_second_half);
    if midnode.is_some() {
        prev_slow.unwrap().borrow_mut().next = midnode.clone();
        midnode.unwrap().borrow_mut().next = restored;
    } else {
        prev_slow.unwrap().borrow_mut().next = restored;
    }

    result
}

fn main() {
    // ----- Example 1: List with a loop -----
    // Create list: 1 -> 2 -> 3 -> 4 -> 5 -> 1
    let node1 = new_node(1);
    let node2 = new_node(2);
    let node3 = new_node(3);
    let node4 = new_node(4);
    let node5 = new_node(5);
    let node6 = new_node(1);

    // Link the nodes.
    node1.borrow_mut().next = Some(node2.clone());
    node2.borrow_mut().next = Some(node3.clone());
    node3.borrow_mut().next = Some(node4.clone());
    node4.borrow_mut().next = Some(node5.clone());
    node5.borrow_mut().next = Some(node6.clone());

    // Create a loop: set node6's next to node3.
    node6.borrow_mut().next = Some(node3.clone());

    println!("Original list (cannot display due to loop):");
    println!("\nLoop created (displaying the list now would loop indefinitely).");
    println!("Removing loop...");
    detect_and_remove_loop(&Some(node1.clone()));
    println!("After removing loop:");
    display_list(&Some(node1.clone()));

    if is_palindrome(&Some(node1.clone())) {
        println!("Linked list is a palindrome.");
    } else {
        println!("Linked list is not a palindrome.");
    }

    // ----- Example 2: Integer palindrome list -----
    let head2 = new_node(1);
    let node2_2 = new_node(2);
    let node3_2 = new_node(2);
    let node4_2 = new_node(1);
    head2.borrow_mut().next = Some(node2_2.clone());
    node2_2.borrow_mut().next = Some(node3_2.clone());
    node3_2.borrow_mut().next = Some(node4_2.clone());

    println!("\nInteger palindrome list:");
    display_list(&Some(head2.clone()));
    if is_palindrome(&Some(head2.clone())) {
        println!("Linked list is a palindrome.");
    } else {
        println!("Linked list is not a palindrome.");
    }

    // ----- Example 3: Character palindrome list -----
    let head3 = new_node('M' as i32);
    let node2_3 = new_node('A' as i32);
    let node3_3 = new_node('D' as i32);
    let node4_3 = new_node('A' as i32);
    let node5_3 = new_node('M' as i32);
    head3.borrow_mut().next = Some(node2_3.clone());
    node2_3.borrow_mut().next = Some(node3_3.clone());
    node3_3.borrow_mut().next = Some(node4_3.clone());
    node4_3.borrow_mut().next = Some(node5_3.clone());

    println!("\nCharacter palindrome list:");
    display_list_char(&Some(head3.clone()));
    if is_palindrome(&Some(head3.clone())) {
        println!("Linked list is a palindrome.");
    } else {
        println!("Linked list is not a palindrome.");
    }
}
