package main

import (
	"fmt"
	"strings"
)

// Doubly linked list node
type node struct {
	value int
	next *node
	prev *node
}

// String method for the node to print its value
func (n node) String() string {
	return fmt.Sprintf("%d", n.value)
}

// Defining the doubly linked list struct
type doublyLinkedList struct {
	head *node
	tail *node
	len int
}

// Inserting a new node at the end of the list
func (dll *doublyLinkedList) add(value int) {
	newNode := &node{value: value}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	}
	dll.len++
}

// Removing a node from the list
func (dll *doublyLinkedList) remove(value int) {
	if dll.head == nil {
		return
	}	

	current := dll.head
	for current != nil {
		if current.value == value {
			// 1 check if current is the head
			// 2 check if current is the tail
			// 3 check if current is in the middle
			if current == dll.head {
				dll.head = current.next
				if dll.head != nil {
					dll.head.prev = nil
				} else {
					dll.tail = nil
				}
			} else if current == dll.tail {
				dll.tail = current.prev
				dll.tail.next = nil
			} else {
				current.prev.next = current.next
				current.next.prev = current.prev
			}
			dll.len--
			return
		}
		current = current.next
	}
}

// Retrieve a node from the list
func (dll doublyLinkedList) get(value int) *node {
	for iterator := dll.head; iterator != nil; iterator = iterator.next {
		if iterator.value == value {
			return iterator
		}
	}
	return nil
}

// String method for the doubly linked list to print its nodes
func (dll doublyLinkedList) String() string {

	// Define a string builder to capture the string value as we iterate through the linked list
	sb := strings.Builder{}

	for iterator := dll.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(fmt.Sprintf("%d", iterator.value))
	}

	return sb.String()

}

func main() {

	dll := doublyLinkedList{}

	// Add nodes
	dll.add(1)
	dll.add(2)
	dll.add(3)
	dll.add(4)
	dll.add(5)

	// Print initial list
	fmt.Println("Initial List:", dll)

	// Get node with data 5
	fmt.Println("Get Node with value 5:", dll.get(5))

	// Remove nodes
	dll.remove(3)
	dll.remove(1)

	// Print updated list
	fmt.Println("Updated List:", dll)
}