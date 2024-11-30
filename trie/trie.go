package trie

import "strings"

// the number of possible characters in the trie
const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	isEnd bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	trie := &Trie{&Node{}}
	return trie
}

func (trie *Trie) Insert(_w string) {
	w := strings.ToLower(_w)
	current := trie.root
	wordLength := len(w)
	for i :=0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		// insert the node if it doesn't exist
		if current.children[charIndex] == nil {
			current.children[charIndex] = &Node{}
		}
		current = current.children[charIndex]
	}
	current.isEnd = true
}

func (trie *Trie) Search(_w string) bool {
	w := strings.ToLower(_w)
	current := trie.root
	wordLength := len(w)
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if current.children[charIndex] == nil {
			return false
		}
		current = current.children[charIndex]
	}
	return current.isEnd
}