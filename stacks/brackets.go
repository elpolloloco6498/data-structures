package main

import "fmt"

type Node struct {
	value string // anytype variable
	next  *Node
}

type List struct {
	head *Node
}

type Stack struct {
	List
}

func (s *Stack) Push(elt string) {
	s.AddTop(elt)
}

func (s *Stack) Pop() {
	s.RemoveFirst()
}

func (n *Node) Display() {
	fmt.Printf("%v %v\n", n.value, n.next)
}

func (l *List) Display() {
	node := l.head
	if node != nil {
		node.Display()
		for node.next != nil {
			node = node.next
			node.Display()
		}
		fmt.Printf("\n")
	} else {
		fmt.Printf("Empty list\n")
	}

}

func (l *List) AddTop(elt string) {
	node := &Node{value: elt, next: nil}
	node.next = l.head
	l.head = node
}

func (l *List) Add(elt string) {
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

func (l *List) InsertAt(pos int, elt string) {
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
	if l.head != nil {
		fmt.Println(l.head.next)
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

// isEmpty
func (l *List) isEmpty() bool {
	return l.head == nil
}

// First
func (l *List) First() *Node {
	return l.head
}

func reverse(char1, char2 string) bool {
	if char1 == "{" && char2 == "}" || char1 == "[" && char2 == "]" || char1 == "(" && char2 == ")" {
		return true
	}
	return false
}

func main() {
	str := "[[{}()]]"

	leftBrackets := Stack{}

	for _, char := range str {
		charStr := string(char)
		if charStr == "[" || charStr == "{" || charStr == "(" {
			leftBrackets.Push(charStr)
		} else {
			if reverse(leftBrackets.First().value, charStr) {
				leftBrackets.Pop()
			}
		}
		leftBrackets.Display()
	}
	if leftBrackets.isEmpty() {
		fmt.Println("Chaine correcte")
	} else {
		fmt.Println("Chaine incorrecte")
	}

}
