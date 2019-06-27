package tree

import (
	"github.com/golang-collections/collections/queue"
	"strconv"
)

// Node of binary Tree
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// Tree A simple binary tree
type Tree struct {
	Root *Node
}

// LeftVal gives Left Node's value
func (n Node) LeftVal() int {
	if n.Left != nil {
		return n.Left.Val
	}
	return -1
}

// RightVal gives Right Node's value
func (n Node) RightVal() int {
	if n.Right != nil {
		return n.Right.Val
	}
	return -1
}

// Inorder of tree
func (t *Tree) Inorder() string {
	return t.Root.inorder()
}

func (n *Node) inorder() (ret string) {
	ret = ""
	if n.Left != nil {
		ret += n.Left.inorder()
	}
	ret += strconv.Itoa(n.Val)
	if n.Right != nil {
		ret += n.Right.inorder()
	}
	return
}

// BFS Breadth First Search
func (t *Tree) BFS() string {
	q := queue.New()
	ret := ""
	q.Enqueue(t.Root)

	for q.Len() != 0 {
		n := q.Dequeue()
		t := n.(*Node)
		ret += strconv.Itoa(t.Val)
		if t.Left != nil {
			q.Enqueue(t.Left)
		}
		if t.Right != nil {
			q.Enqueue(t.Right)
		}
	}
	return ret
}
