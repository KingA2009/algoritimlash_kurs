package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *LinkedList) Display() {
	current := list.head

	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}

	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}
func (list *LinkedList) Remove(value int) *Node {
	head := list.head
	if head != nil && head.data == value {
		return head.next
	}

	currentNode := head
	var prevNode *Node

	for currentNode != nil && currentNode.data != value {
		prevNode = currentNode
		currentNode = currentNode.next
	}

	if currentNode != nil {
		prevNode.next = currentNode.next
	}

	return head
}

func main() {
	list := LinkedList{}

	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(40)

	list.Remove(20)

	list.Display()
}
