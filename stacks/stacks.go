package main

import "fmt"

type Node struct {
	value interface{} // anytype variable
	next  *Node
}

type List struct {
	head *Node
}

type Stack struct {
	List
}

func (s *Stack) Push(elt interface{}) {
	s.AddTop(elt)
}

func (s *Stack) Pop() {
	s.RemoveFirst()
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

func (l *List) AddTop(elt interface{}) {
	node := &Node{value: elt, next: nil}
	node.next = l.head
	l.head = node
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

func (l *List) RemoveFirst() {
	if l.head.next != nil {
		l.head = l.head.next
	} else {
		fmt.Println("Empty list")
	}
}

func (l *List) RemoveAt(pos int) {
	// 2 pointers traverse linkedlist
	if pos != 1 {
		i := 1
		trav1 := l.head
		trav2 := trav1.next
		for trav1.next != nil && trav2.next != nil && i < pos-1 {
			trav1 = trav1.next
			trav2 = trav2.next
			i++
		}
		// remove node
		trav1.next = trav2.next
	}
	//else
	l.head = l.head.next
}

func main() {
	fmt.Println("Hello")

	// create linked list
	stack := Stack{}
	stack.Add(7)
	stack.Add(0)
	stack.Add(4)
	stack.Add(9)
	stack.Add(15)
	stack.AddTop(33)
	stack.Pop()
	stack.Display()
}
