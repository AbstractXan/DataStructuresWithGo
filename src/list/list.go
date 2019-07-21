package list

type node struct {
	val  int
	next *node
}

// List structure
type List struct {
	head *node
	tail *node
	len  int
}

// Returns previous node to this position
func (l List) previousNode(pos int) *node {
	ptr := l.head

	// Out of left bounds
	if pos <= 1 {
		return l.head
	}

	// Out of right bounds
	if pos > l.len {
		return l.tail
	}

	// Iterate through pos-1
	for i := 1; i < pos-1; i++ {
		ptr = ptr.next
	}

	return ptr
}

// inserts node after a node
func (l *List) insertAfter(val int, ptr *node) {

	var temp node        // Temporary variable
	temp.val = val       // Initialise value
	temp.next = ptr.next // Initialise pointer
	ptr.next = &temp     // Reset ptr next as temp

	if l.tail == ptr { // If ptr is tail, update list tail
		l.tail = &temp
	}

	// Increase length
	l.len++
}

// GetNode returns node
func (l List) getNode(pos int) *node {
	ptr := l.head

	if pos < 1 || pos > l.len {
		return nil
	}

	for i := 1; i < pos; i++ {
		ptr = ptr.next
	}

	return ptr
}

// Insert a node to the list
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

// Search returns node position with given value
func (l List) Search(val int) int {
	ptr := l.head

	for i := 1; ptr != nil; i++ {
		if ptr.val == val {
			return i
		}
		ptr = ptr.next
	}

	return -1
}

// DeleteAt deletes a node given a position
func (l *List) DeleteAt(pos int) {

	//Out of bounds
	if pos <= 0 || pos > l.len {
		return
	}

	//Remove head
	if pos == 1 {
		l.head = l.head.next
		l.len--
		return
	}

	prev := l.previousNode(pos)
	prev.next = prev.next.next
	l.len--
}

// DeleteVal deletes a node given a value
func (l *List) DeleteVal(val int) {
	pos := l.Search(val)
	l.DeleteAt(pos)
}

// GetAt returns node value at given position
func (l List) GetAt(pos int) int {
	node := l.getNode(pos)
	if node == nil {
		return -1
	}
	return node.val
}

// --- Additional functions --- //

// Length of LL
func (l List) Length() int {
	return l.len
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
		l.len++
		return
	}

	l.insertAfter(newVAl, l.previousNode(pos))
}

// GetArray returns the list as a slice
func (l List) GetArray() []int {
	var p = l.head
	var arr []int

	for p != nil {
		arr = append(arr, p.val)
		p = p.next
	}
	return arr
}
