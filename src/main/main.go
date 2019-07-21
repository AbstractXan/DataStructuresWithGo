package main

import "list"
import "stack"
import "queue"
import "tree"
import "bst"
import "fmt"

func main() {

	var l list.List
	fmt.Println(l.GetAt(1))
	l.Insert(2)
	l.InsertAt(1, 1) //Additional functions
	fmt.Println(l.GetArray())
	l.DeleteVal(2)
	fmt.Println(l.GetArray())

	var s stack.Stack
	s.Push(1)
	fmt.Println("Top: ", s.Top())
	s.Push(2)
	fmt.Println("Top: ", s.Top())
	fmt.Println("Popped: ", s.Pop())
	fmt.Println("Top: ", s.Top())
	fmt.Println()

	var q queue.Queue
	q.Enqueue(1)
	fmt.Println("Front: ", q.Front())
	q.Enqueue(2)
	fmt.Println("Front: ", q.Front())
	fmt.Println("Back: ", q.Back())
	q.Dequeue()
	fmt.Println("Back: ", q.Back())
	q.Dequeue()
	q.Dequeue()
	fmt.Println("Back: ", q.Back())

	var t tree.Tree
	var n2 tree.Node
	n2.Val = 2
	var n3 tree.Node
	n3.Val = 3
	var n1 tree.Node
	n1.Val = 1
	n2.Left = &n1
	n2.Right = &n3
	t.Root = &n2
	fmt.Println(t.Inorder())
	fmt.Println(t.BFS())

	var st bst.BST
	st.Insert(1)
	st.Insert(2)
	st.Insert(-1)
	fmt.Println(st.Search(-1))

}
