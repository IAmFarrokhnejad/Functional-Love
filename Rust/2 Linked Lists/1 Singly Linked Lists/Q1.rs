// The program below demosntrates implementation of a singly linked list. Following functions are included:

// 1. create_node_list(n: usize): Creates a new linked list with n nodes, where data for each node is input by the user.
// 2. display_list(head): Prints all elements in the linked list with arrow separators (e.g., "1 -> 2 -> 3").
// 3. delete_first_node(head): Removes the first node from the linked list and updates the head.
// 4. delete_middle_node(head): Finds and removes the middle node using the "slow and fast pointer" technique.
// 5. delete_last_node(head): Traverses to the end of the list and removes the last node.
// 6. reverse_and_display(head): Reverses the order of nodes in the list and displays the result.
// 7. insert_node_beginning(head, num): Adds a new node with value 'num' at the start of the list.
// 8. insert_node_end(head, num): Adds a new node with value 'num' at the end of the list.
// 9. insert_node_middle(head, num): Adds a new node with value 'num' in the middle of the list.
// 10. search_element(head, value): Searches for a value in the list and prints its position if found.

// Author: Morteza Farrokhnejad

use std::cell::RefCell;
use std::io;
use std::rc::Rc;

#[derive(Debug)]
struct Node {
    num: i32,
    next: Option<Rc<RefCell<Node>>>,
}

fn create_node_list(n: usize) -> Option<Rc<RefCell<Node>>> {
    let mut head: Option<Rc<RefCell<Node>>> = None;
    let mut tail: Option<Rc<RefCell<Node>>> = None;

    for i in 1..=n {
        let mut input = String::new();
        println!("Enter data for node {}: ", i);
        io::stdin()
            .read_line(&mut input)
            .expect("Failed to read input");
        let num: i32 = input.trim().parse().expect("Please enter a valid number");

        let new_node = Rc::new(RefCell::new(Node { num, next: None }));

        if head.is_none() {
            head = Some(Rc::clone(&new_node));
            tail = Some(new_node);
        } else if let Some(tail_node) = tail.take() {
            tail_node.borrow_mut().next = Some(Rc::clone(&new_node));
            tail = Some(new_node);
        }
    }

    head
}

fn display_list(head: Option<Rc<RefCell<Node>>>) {
    if head.is_none() {
        println!("List is empty.");
        return;
    }

    let mut current = head;
    println!("Linked List Content:");
    while let Some(node) = current {
        print!("{} ", node.borrow().num);
        current = node.borrow().next.clone();
        if current.is_some() {
            print!("-> ");
        }
    }
    println!();
}

fn delete_first_node(head: &mut Option<Rc<RefCell<Node>>>) {
    if let Some(first_node) = head.take() {
        let next = first_node.borrow_mut().next.take();
        *head = next;
        println!("First node deleted successfully.");
    } else {
        println!("List is empty, nothing to delete.");
    }
}

fn delete_middle_node(head: &mut Option<Rc<RefCell<Node>>>) {
    if head.is_none() || head.as_ref().unwrap().borrow().next.is_none() {
        println!("List is too short to delete middle node.");
        return;
    }

    let mut slow = head.clone();
    let mut fast = head.clone();
    let mut prev: Option<Rc<RefCell<Node>>> = None;

    // Find the middle node using slow and fast pointers
    while let (Some(_), Some(next_fast)) =
        (fast.clone(), fast.clone().unwrap().borrow().next.clone())
    {
        fast = next_fast.borrow().next.clone();
        if fast.is_none() {
            break;
        }
        prev = slow.clone();
        slow = slow.unwrap().borrow().next.clone();
    }

    // Delete the middle node
    if let Some(prev_node) = prev {
        let middle = prev_node.borrow().next.clone();
        if let Some(middle_node) = middle {
            let next = middle_node.borrow().next.clone();
            prev_node.borrow_mut().next = next;
            println!("Middle node deleted successfully.");
        }
    } else {
        // If prev is None, we need to delete the first node
        let next = head.as_ref().unwrap().borrow().next.clone();
        *head = next;
        println!("Middle node (first node) deleted successfully.");
    }
}

fn delete_last_node(head: &mut Option<Rc<RefCell<Node>>>) {
    if head.is_none() {
        println!("List is empty, nothing to delete.");
        return;
    }

    if head.as_ref().unwrap().borrow().next.is_none() {
        // Only one node in the list
        head.take();
        println!("Last node (only node) deleted successfully.");
        return;
    }

    let mut current = head.clone();
    let mut prev: Option<Rc<RefCell<Node>>> = None;

    // Traverse to the last node
    while let Some(curr_node) = current.clone() {
        if curr_node.borrow().next.is_none() {
            break;
        }
        prev = current;
        current = curr_node.borrow().next.clone();
    }

    // Delete the last node
    if let Some(prev_node) = prev {
        prev_node.borrow_mut().next = None;
        println!("Last node deleted successfully.");
    }
}

fn reverse_and_display(head: Option<Rc<RefCell<Node>>>) -> Option<Rc<RefCell<Node>>> {
    let mut prev: Option<Rc<RefCell<Node>>> = None;
    let mut current = head;

    while let Some(curr_node) = current {
        let next = curr_node.borrow_mut().next.take();
        curr_node.borrow_mut().next = prev;
        prev = Some(curr_node);
        current = next;
    }

    println!("\nReversed Linked List Content:");
    display_list(prev.clone());

    prev
}

fn insert_node_beginning(head: &mut Option<Rc<RefCell<Node>>>, num: i32) {
    let new_node = Rc::new(RefCell::new(Node {
        num,
        next: head.take(),
    }));

    *head = Some(new_node);
}

fn insert_node_end(head: &mut Option<Rc<RefCell<Node>>>, num: i32) {
    let new_node = Rc::new(RefCell::new(Node { num, next: None }));

    if let Some(first_node) = head {
        let current = Rc::clone(first_node);
        {
            let mut current_ref = current;
            while let Some(next) = current_ref.clone().borrow().next.clone() {
                current_ref = next;
            }
            current_ref.borrow_mut().next = Some(new_node);
        }
    } else {
        *head = Some(new_node);
    }
}

fn insert_node_middle(head: &mut Option<Rc<RefCell<Node>>>, num: i32) {
    if head.is_none() {
        *head = Some(Rc::new(RefCell::new(Node { num, next: None })));
        return;
    }

    let mut slow = head.clone();
    let mut fast = head.clone();

    // Find the middle using slow and fast pointers
    while let Some(fast_node) = fast.clone() {
        if fast_node.borrow().next.is_none() {
            break;
        }
        
        if let Some(next_fast) = fast_node.borrow().next.clone() {
            if next_fast.borrow().next.is_none() {
                break;
            }
            fast = next_fast.borrow().next.clone();
            slow = slow.unwrap().borrow().next.clone();
        } else {
            break;
        }
    }

    // Insert the new node after the slow pointer
    if let Some(slow_node) = slow {
        let new_node = Rc::new(RefCell::new(Node {
            num,
            next: slow_node.borrow().next.clone(),
        }));
        slow_node.borrow_mut().next = Some(new_node);
    }
}

fn search_element(head: Option<Rc<RefCell<Node>>>, value: i32) {
    if head.is_none() {
        println!("List is empty.");
        return;
    }

    let mut current = head;
    let mut position = 1;
    let mut found = false;

    while let Some(node) = current {
        if node.borrow().num == value {
            println!("Element found at node {}", position);
            found = true;
            break;
        }
        current = node.borrow().next.clone();
        position += 1;
    }

    if !found {
        println!("Element not found in the list.");
    }
}

fn main() {
    println!("Enter the number of nodes: ");
    let mut input = String::new();
    io::stdin()
        .read_line(&mut input)
        .expect("Failed to read input");
    let n: usize = input.trim().parse().expect("Please enter a valid number");

    let mut head = create_node_list(n);
    println!("\nOriginal Linked List Content:");
    display_list(head.clone());

    // Test insert operations
    insert_node_beginning(&mut head, 7);
    println!("\nAfter inserting 7 at the beginning:");
    display_list(head.clone());

    insert_node_end(&mut head, 99);
    println!("\nAfter inserting 99 at the end:");
    display_list(head.clone());

    insert_node_middle(&mut head, 50);
    println!("\nAfter inserting 50 in the middle:");
    display_list(head.clone());

    // Test delete operations
    delete_first_node(&mut head);
    println!("\nAfter deleting first node:");
    display_list(head.clone());

    delete_middle_node(&mut head);
    println!("\nAfter deleting middle node:");
    display_list(head.clone());

    delete_last_node(&mut head);
    println!("\nAfter deleting last node:");
    display_list(head.clone());

    let head = reverse_and_display(head); 
    println!("\nEnter the element to be searched: ");
    let mut search_input = String::new();
    io::stdin().read_line(&mut search_input).expect("Failed to read input");

    let search_value: i32 = search_input.trim().parse().expect("Please enter a valid number");

    search_element(head, search_value); // Use the reversed list head
}