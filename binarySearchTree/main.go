package main

import "fmt"

var count int

// Node represents the components of a binary search tree
type Node struct {
	Key int
	Left *Node
	Right *Node
}

// Insert will add a node to the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		// Move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// Move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search will take in a key value
// and RETURN true if there is a node with that value
func (n *Node) Search (k int) bool {

	count++

	if n == nil {
		return false
	}

	if n.Key < k {
		// Move Right
		return n.Right.Search(k)
	} else if n.Key > k {
		// Move Left
		return n.Left.Search(k)
	}

	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(50)
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(150)
	fmt.Println(tree)
	fmt.Println(tree.Search(200))
	fmt.Println(count)
}