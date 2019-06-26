package main

import "fmt"
import "list"

func main() {
	var l list.List
	l.Insert(1)
	l.Insert(2)
	l.InsertAt(3, 2)
	l.Print()
}
