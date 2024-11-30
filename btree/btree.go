package btree

// A component in the binary tree
type Node struct {
	Key int
	Left *Node
	Right *Node
}

// Insert a node into the tree
func (node *Node) Insert(key int) {
	if key > node.Key {
		if (node.Right == nil) {
			node.Right = &Node{key, nil, nil}
		} else {
			node.Right.Insert(key)
		}
	} else {
		if (node.Left == nil) {
			node.Left = &Node{key, nil, nil}
		} else {
			node.Left.Insert(key)
		}
	}
}


// Find a node in the tree
// It return true if the node is found and false if it isn't
func (node *Node) Search(key int) bool {
	if key == node.Key {
		return true
	} else if key > node.Key {
		if (node.Right == nil) {
			return false
		} else {
			return node.Right.Search(key)
		}
	} else {
		if (node.Left == nil) {
			return false
		} else {
			return node.Left.Search(key)
		}
	}
}
