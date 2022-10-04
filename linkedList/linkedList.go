package main

import "fmt"

// implement Add
// implement InsertAt
// implement RemoveAt TODO
// implement Display

type Node struct {
	value interface{} // anytype variable
	next *Node
}

type List struct {
	head *Node
}

func (n *Node) Display() {
	fmt.Println(n.value, n.next)
	fmt.Println("\n")
}

func (l *List) Display() {
	node := l.head
	node.Display()
	for node.next != nil {
		node = node.next
		node.Display()
	}
}

func (l *List) Add(elt interface{}) {
	if l.head == nil {
		//fmt.Println(l.head)
		l.head = &Node{value: elt, next: nil}
	} else {
		// get head
		node := l.head
		// get the last node
		for node.next != nil {
			node = node.next
		}
		node.next = &Node{value: elt, next: nil}
	}
}

func (l *List) InsertAt(pos int, elt interface{}) {
	// navigate to insert position
	i := 1
	node := l.head
	for node.next != nil && i < pos-1 {
		node = node.next
		i++
	}
	// insert node
	fmt.Println(node.value)
	insertedNode := &Node{value: elt, next: node.next}
	node.next = insertedNode

}

func main() {
	fmt.Println("Hello")

	// create linked list
	list := List{}
	list.Add(5)
	list.Add(44)
	list.Add(3)
	list.Add(21)
	list.InsertAt(3, "test")
	list.Display()
	// expected 5,44,"test",3,21
}