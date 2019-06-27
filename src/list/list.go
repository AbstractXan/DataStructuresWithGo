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

// DeleteLast deletes Last Element
func (l *List) DeleteLast() {
	l.Delete(l.len)
}

// Delete node at position
func (l *List) Delete(pos int) {

	//Out of bounds
	if pos <= 0 || pos > l.len {
		return
	}
	prev := l.previousNode(pos)

	//Remove head
	if pos == 1 {
		prev = l.head
		l.head = l.head.next
		prev = new(node)
		l.len--
		return
	}
	toRemove := prev.next
	prev.next = toRemove.next
	toRemove = new(node)
	l.len--
}

// Length of LL
func (l List) Length() int {
	return l.len
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

	l.insertAfter(value, l.tail)
}

// inserts node after a node
func (l *List) insertAfter(val int, ptr *node) {
	var temp node
	temp.val = val
	temp.next = ptr.next
	ptr.next = &temp
	if l.tail == ptr {
		l.tail = &temp
	}
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
		var temp node
		temp.val = newVAl
		temp.next = l.head
		l.head = &temp
		return
	}

	l.insertAfter(newVAl, l.previousNode(pos))
}

// returns previous node to this position
func (l List) previousNode(pos int) *node {
	ptr := l.head

	if pos <= 1 {
		return l.head
	}

	if pos > l.len {
		return l.tail
	}

	for i := 1; i < pos-1; i++ {
		ptr = ptr.next
	}

	return ptr
}

// GetArray returns the list as a slice
func (l List) GetArray() []int {
	var p = l.head
	var arr []int

	for p != nil {
		arr = append(arr, p.val)
		p = p.next
	}
	fmt.Println(" ")
	return arr
}

// GetVal returns value for given position
func (l List) GetVal(pos int) int {
	return l.getNode(pos).val
}

// GetNode returns node
func (l List) getNode(pos int) *node {
	ptr := l.head

	if pos < 1 {
		return l.head
	}

	if pos > l.len {
		return l.tail
	}

	for i := 1; i < pos; i++ {
		ptr = ptr.next
	}

	return ptr
}
