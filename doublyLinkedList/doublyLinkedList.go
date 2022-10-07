package main

import "fmt"

// struct Node
type Node struct {
	value interface{}
	next  *Node
	prev  *Node
}

// struct DoublyLinkedList
type DoublyLinkedList struct {
	head *Node
}

// method Display LinkedList
func (l *DoublyLinkedList) Display() {
	node := l.head
	node.Display()
	for node.next != nil {
		node = node.next
		node.Display()
	}
}

// method Display Node
func (n *Node) Display() {
	fmt.Printf("Node : %+v\n", n.value)
}

// method Add
func (l *DoublyLinkedList) Add(elt interface{}) {
	// naviguate to the end of the list
	if l.head != nil {
		node := l.head
		for node.next != nil {
			node = node.next
		}
		node.next = &Node{value: elt, next: nil, prev: node}
	} else {
		l.head = &Node{value: elt, next: nil, prev: nil}
	}
}

// method InsertAt
func (l *DoublyLinkedList) InsertAt(pos int, elt interface{}) {
	// naviguate to the insertion position
	node := l.head
	i := 1
	for node.next != nil && i < pos {
		node = node.next
		i++
	}

	prev := node.prev
	insertedNode := &Node{value: elt, next: node, prev: node.prev}
	prev.next = insertedNode
	node.prev = insertedNode
}

// method RemoveAt
func (l *DoublyLinkedList) RemoveAt(pos int) {
	// naviguate to the insertion position
	trav := l.head
	i := 1
	for trav.next != nil && i < pos {
		trav = trav.next
		i++
	}

	prev := trav.prev
	next := trav.next

	if next == nil {
		prev.next = nil
	} else if prev == nil {
		next.prev = nil
		l.head = next
	} else {
		prev.next = next
		next.prev = prev
	}

}

// method RemoveAt

func main() {
	fmt.Println("Hello")
	list := DoublyLinkedList{}
	list.Add(7)
	list.Add(0)
	list.Add(4)
	list.Add(9)
	list.Add(15)
	list.RemoveAt(1)
	list.Display()
}
