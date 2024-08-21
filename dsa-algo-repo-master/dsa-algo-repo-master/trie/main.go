package main

import "fmt"

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

// Node struct
type Node struct {
	children [AlphabetSize]*Node
	isEnd bool
}

// Trie represent a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie will create new Trie
func InitTrie() *Trie {
	result := &Trie{root:&Node{}}
	return result
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search will take in a word and RETURN true if that word is included in the trie
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode.isEnd
}

func main() {
	myTrie := InitTrie()

	toAdd := []string{
		"oluwatobi",
		"argon",
		"soap",
		"lifo",
	}
	for _, value := range toAdd {
		myTrie.Insert(value)
	}
	myTrie.Insert("oluwatobi")
	fmt.Println(myTrie.Search("wizard"))
	fmt.Println(myTrie.root)
}