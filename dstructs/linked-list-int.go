package main

import "fmt"
// This file implements a singly linked list in Go, for int types

type Node struct {
 	next *Node
	id	  int
}

type LinkedList struct {
	head *Node
}

// Main routine tests llist functions
func main() {
	l := Initialize(0)
	//Display(l)	
	fmt.Println("Inserting on list...")

	for i := 1; i < 5; i++ {
		Insert(&l, i)
	}

	Display(l)

	//p := Search(&l, 2)
	fmt.Println("Deleting from list...")
	Delete(&l, 4)
	Display(l)
	
}

func Initialize(val int) (l LinkedList) {
	firstNode := Node{next: nil, id: val}
	l = LinkedList{head: &firstNode}
	return 
}

func Display(l LinkedList) {
	for l.head.next != nil {
		fmt.Printf("%d\n", l.head.id)
		l.head = l.head.next
	}
	fmt.Printf("%d\n", l.head.id)
}

func Insert(l *LinkedList, val int) {
	newNode := Node{next: l.head, id: val}
	l.head = &newNode
}

func Search(l *LinkedList, val int) (found *Node) {
	fmt.Println("Inside search")
	found = nil
	if (l.head.id == val) {
		return 
	}

	for l.head.next != nil {
		if(l.head.id == val) {			
			found = l.head
			return 
		}
		fmt.Printf("id: %d\n", l.head.id)
		l.head = l.head.next
	}
	return 
}

func Delete(l *LinkedList, val int) {
	// If node to be deleted is the current head, delete it 
	if l.head.id == val {
		auxNode := l.head.next
		l.head = nil
		l.head = auxNode
	} else {
		// Middle node routine
		auxNode := l.head
		for l.head.next.id != val {
			l.head = l.head.next	
			fmt.Printf("current idx %v next %v\n", l.head.id, l.head.next.id)
		} 
		l.head.next = l.head.next.next
		l.head = auxNode
	}


}