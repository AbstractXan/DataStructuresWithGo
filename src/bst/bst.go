package bst

import "fmt"

// Node of binary Tree
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// BST Binary Search Tree
type BST struct {
	root *Node
}

//insert a value, doesnot allow duplicate value
func (node *Node) insert(val int) {

	// if same, do nothing
	if val == node.Val {
		return
	}

	// insert on left if less than
	if val < node.Val {

		// if left is nil
		if node.Left == nil {
			var tmp Node
			tmp.Val = val
			node.Left = &tmp
		} else {
			// insert in left subtree
			node.Left.insert(val)
		}
		return
	}

	//insert on right if greater than.
	// if right is nil
	if node.Right == nil {
		var tmp Node
		tmp.Val = val
		node.Right = &tmp
	} else {
		// insert on right subtree
		node.Right.insert(val)
	}
}

// search helper function to iteratively search for value
func (node *Node) search(val int) bool {

	// if empty
	if node == nil {
		return false
	}

	// if equal, return true
	if val == node.Val {
		return true
	}

	// if less, search on the left subtree
	if val < node.Val {
		return node.Left.search(val)
	}

	// if greater, search on the right subtree
	return node.Right.search(val)
}

// Insert value
func (b *BST) Insert(val int) {
	if b.root == nil {
		var tmp Node
		tmp.Val = val
		b.root = &tmp
		return
	}

	b.root.insert(val)
}

// Search value
func (b *BST) Search(val int) bool {
	return b.root.search(val)
}

func (node *Node) find(val int) (*Node, bool) {

	// if empty
	if node == nil {
		var tmp Node
		return &tmp, false
	}

	// // if equal, return true
	if val == node.Val {
		return node, true
	}

	// if less, search on the left subtree
	if val < node.Val {
		if node.Left != nil && val == node.Left.Val {
			return node.Left, true
		}
		return node.Left.find(val)
	}

	// if greater, search on the right subtree
	if node.Right != nil && val == node.Right.Val {
		return node.Right, true
	}
	return node.Right.find(val)
}

// Delete value
func (b *BST) Delete(val int) {

	if b.root != nil {
		return
	}

	// Find iteratively through the tree
	// Obtain pointer to node to delete
	node, flag := b.root.find(val)
	if !flag {
		return
	}

	// if node is leaf
	if node.Left == nil && node.Right == nil {
		fmt.Println("Pointer set to null", node, node.Val)
		node = nil
	}

}
