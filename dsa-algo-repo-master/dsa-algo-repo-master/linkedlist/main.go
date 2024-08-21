package main

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next *node
}

// Two string methods for the node struct
func (n node) String() string {
	return fmt.Sprintf("%d", n.value)
}

type linkedList struct {
	head *node
	len int
}

// Defining the 4 methods of the linkedList struct

func (l *linkedList) add(value int) {
	newNode := new(node)
	newNode.value = value
	
	if l.head == nil {
		l.head = newNode
	} else {
		iterator := l.head
		// Continue iterating until the last node is reached
		for ; iterator.next != nil; iterator = iterator.next {
		}
		iterator.next = newNode
	}
}

func (l *linkedList) remove(value int) {
	var previous *node

	for current := l.head; current != nil; current = current.next {
		if current.value == value {
			if l.head == current {
				l.head = current.next
			} else {
				previous.next = current.next
				return
			}
		}

		previous = current
	}
}

func (l linkedList) get(value int) *node {
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		if iterator.value == value {
			return iterator
		}
	}

	 return nil
}

func (l linkedList) String() string {
	// Define a string builder to capture the string value as we iterate through the linked list
	sb := strings.Builder{}

	for iterator := l.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(fmt.Sprintf("%d", iterator.value))
	}

	return sb.String()
}

func main() {
	
	l := linkedList{}
	l.add(1)
	l.add(2)
	l.add(3)
	l.add(4)
	l.add(5)
	fmt.Println(l)
	fmt.Println(l.get(5))
	l.remove(3)
	l.remove(1)
	fmt.Println(l)
}

