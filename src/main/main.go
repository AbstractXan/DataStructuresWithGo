package main

import "list"

import "fmt"

func main() {
	var l list.List
	l.InsertAt(2, 1)
	l.InsertAt(3, 1)
	fmt.Println(l.GetArray())
}
