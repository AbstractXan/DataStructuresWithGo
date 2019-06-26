package list

import (
	"testing"
)

func initList() *List {
	var l List
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	return &l
}

func TestInsertOne(t *testing.T) {
	l := initList()
	input := 4
	expected := input

	//test step
	l.Insert(input)

	output := l.GetVal(l.Length())
	if output != expected {
		t.Error("\nTest Failed: {} inserted, {} expected, recieved: {}", input, expected, output)
	}
}

func TestDeleteOne(t *testing.T) {
	l := initList()
	input := l.Length()
	expected := input - 1

	//test step
	l.Delete(l.Length())

	output := l.Length()
	if output != input-1 {
		t.Error("\nTest Failed: {} deleted, {} expected, recieved: {}", input, expected, output)
	}
}
func TestOutOfBoundsDelete(t *testing.T) {
	var l List
	l.Insert(1)

	input := l.len
	expected := input
	//test step
	l.Delete(-1)

	if output := l.Length(); output != expected {
		t.Error("Test Failed: {} inputted, {} expected, recieved: {}", input, expected, output)
	}
}

func TestDeleteListWithOne(t *testing.T) {
	var l List
	l.Insert(1)

	input := l.Length()
	expected := l.Length() - 1
	//test step
	l.Delete(l.Length())
	if output := l.Length(); output != expected {
		t.Error("Test Failed: {} inputted, {} expected, recieved: {}", input, expected, output)
	}
}

func TestGetArray(t *testing.T) {
	var l List
	input := []int{1, 2, 3}
	for i := 0; i < len(input); i++ {
		l.Insert(input[i])
	}
	expected := input
	output := l.GetArray()

	if len(output) != len(expected) {
		t.Error("Test Failed: {} inputted, {} expected, recieved: {}", input, expected, output)
	}
}

func TestInsertAtEmpty(t *testing.T) {
	var l List
	input := 4
	expected := input

	l.InsertAt(input, 1)

	output := l.GetVal(1)

	if output != expected {
		t.Error("Test Failed: {} inputted, {} expected, recieved: {}", input, expected, output)
	}
}

func TestInsertAtStarting(t *testing.T) {
	l := initList()
	input := 4
	expected := input

	l.InsertAt(input, 1)

	output := l.GetVal(1)

	if output != expected {
		t.Error("Test Failed: {} inputted, {} expected, recieved: {}", input, expected, output)
	}
}

func TestInsertAtMiddle(t *testing.T) {
	l := initList()
	input := 4
	expected := input

	l.InsertAt(input, 2)

	output := l.GetVal(2)

	if output != expected {
		t.Error("Test Failed: {} inputted, {} expected, recieved: {}", input, expected, output)
	}
}

func TestGetNodeLessThanScope(t *testing.T) {
	l := initList()
	input := 4
	expected := l.head

	output := l.getNode(-1)

	if output != expected {
		t.Error("Test Failed: {} inputted, {} expected, recieved: {}", input, expected, output)
	}
}

func TestGetNodeGreaterThanScope(t *testing.T) {
	l := initList()
	expected := l.tail

	output := l.getNode(l.len + 1)

	if output != expected {
		t.Error("Test Failed: {} expected, recieved: {}", expected, output)
	}
}

func TestPreviousNodeLessThanScope(t *testing.T) {
	l := initList()
	input := l.GetArray()
	expected := l.head
	output := l.previousNode(-1)
	if output != expected {
		t.Error("Test Failed: {} inputted, {} expected, recieved: {}", input, expected, output)
	}
}

func TestPreviousNodeGreaterThanScope(t *testing.T) {
	l := initList()
	input := l.GetArray()
	expected := l.tail
	output := l.previousNode(l.len + 1)
	if output != expected {
		t.Error("Test Failed: {} inputted, {} expected, recieved: {}", input, expected, output)
	}
}

func TestDeleteLast(t *testing.T) {
	l := initList()
	input := l.Length()
	expected := input - 1
	l.DeleteLast()
	if output := l.Length(); output != expected {
		t.Error("Test Failed: {} inputted, {} expected, recieved: {}", input, expected, output)
	}

}
