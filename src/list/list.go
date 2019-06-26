package list

import "fmt"

type node struct {
	val  int
	next *node
}

// List structure
// Insert(), Print()
type List struct {
	head *node
	tail *node
	len  int
}

// Insert a value at the end
func (l *List) Insert(value int) {
	if l.head == nil {
		var temp node
		temp.val = value
		l.head = &temp
		l.tail = &temp
		l.len++
		return
	}

	l.insertAfter(value, l.head)
}

// inserts node after a node
func (l *List) insertAfter(val int, ptr *node) {
	var temp node
	temp.val = val
	temp.next = ptr.next
	ptr.next = &temp
	l.len++
}

// InsertAt inserts a value at some middle position
func (l *List) InsertAt(newVAl int, pos int) {
	//if pos>len, insert at end
	if pos > l.len || l.len == 0 {
		l.Insert(newVAl)
		return
	}
	//if insertion at starting
	if pos == 1 {
		l.insertAfter(newVAl, l.head)
		return
	}

	l.insertAfter(newVAl, l.previousnode(pos))
}

// returns previous node to this position
func (l List) previousnode(pos int) *node {
	ptr := l.head

	if pos < 1 {
		fmt.Println("Error. Previous node parameter cannot be < 1 \n Returning list head.")
		return l.head
	}

	for i := 1; i < pos-1; i++ {
		ptr = ptr.next
	}

	return ptr
}

// Print the linked list
func (l List) Print() {
	var p = l.head
	for p != nil {
		fmt.Print(p.val, " ")
		p = p.next
	}
	fmt.Println(" ")
	return
}
