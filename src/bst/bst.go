package bst

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

func (node *Node) insert(val int) {
	// if empty
	if node == nil {
		node.Val = val
		return
	}

	// if not empty
	if val < node.Val {
		node.Left.insert(val)
		return
	}

	//insert on right if equal to or greater than.
	node.Right.insert(val)
}

// Insert value
func (b *BST) Insert(val int) {
	b.root.insert(val)
}
