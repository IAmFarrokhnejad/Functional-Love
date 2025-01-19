// The program below creates and displays a singly linked list. "reverse_and_display" function reverses the list and displays it. "insert_node_beginning" function inserts a new node at the beginning of the list.
// "insertNodeAtBeginning" function inserts a new node at the end of the linked list.
// Author: Morteza Farrokhnejad
use std::io;
use std::rc::Rc;
use std::cell::RefCell;

// Structure for a node in a linked list
#[derive(Debug)]
struct Node {
    num: i32,                    // Data of the node
    next: Option<Rc<RefCell<Node>>>, // Reference to the next node
}

fn create_node_list(n: usize) -> Option<Rc<RefCell<Node>>> {
    let mut head: Option<Rc<RefCell<Node>>> = None;
    let mut tail: Option<Rc<RefCell<Node>>> = None;

    for i in 1..=n {
        let mut input = String::new();
        println!("Enter data for node {}: ", i);
        io::stdin().read_line(&mut input).expect("Failed to read input");
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

fn display_list(mut head: Option<Rc<RefCell<Node>>>) {
    if head.is_none() {
        println!("List is empty.");
        return;
    }

    println!("Linked List Content:");
    while let Some(node) = head {
        print!("Data = {}\n", node.borrow().num);
        head = node.borrow().next.clone();
    }
}

fn reverse_and_display(mut head: Option<Rc<RefCell<Node>>>) -> Option<Rc<RefCell<Node>>> {
    let mut prev: Option<Rc<RefCell<Node>>> = None;
    let mut current = head.take();
    let mut next: Option<Rc<RefCell<Node>>>;

    while let Some(curr_node) = current {
        next = curr_node.borrow_mut().next.take(); // Detach the current node from the rest of the list
        curr_node.borrow_mut().next = prev.clone(); // Reverse the link
        prev = Some(curr_node); // Move prev to the current node
        current = next; // Move to the next node
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

    if let Some(ref mut current) = head {
        let mut current = Rc::clone(current); // Clone the Rc to avoid borrowing issues
        loop {
            let next_node = current.borrow().next.clone(); // Clone the next pointer
            match next_node {
                Some(next) => current = next, // Move to the next node
                None => break, // Exit the loop if we reach the end
            }
        }
        current.borrow_mut().next = Some(new_node); // Attach the new node at the end
    } else {
        *head = Some(new_node); // Handle empty list
    }
}



fn main() {
    println!("Enter the number of nodes: ");
    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("Failed to read input");
    let n: usize = input.trim().parse().expect("Please enter a valid number");

    let mut head = create_node_list(n);
    println!("\nOriginal Linked List Content:");
    display_list(head.clone());

    insert_node_beginning(&mut head, 7);
    println!("\nAfter inserting 7 at the beginning:");
    display_list(head.clone());

    insert_node_end(&mut head, 99);
    println!("\nAfter inserting 99 at the end:");
    display_list(head.clone());

    reverse_and_display(head);
}