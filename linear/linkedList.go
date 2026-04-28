package linear

import (
	"fmt"
)

type Node[T any] struct {
	data T
	next *Node[T]
	prev *Node[T]
}

type LinkedList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length uint8
}

// InsertAtHead() -> inserts data and sets it as head
func (ll *LinkedList[T]) InsertAtHead(data T) {
	newNode := &Node[T]{data: data}

	if ll.head == nil {

		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.next = ll.head
		ll.head.prev = newNode
		ll.head = newNode
	}

	ll.length += 1
}

// InsertAtTail() -> inserts node at tail
func (ll *LinkedList[T]) InsertAtTail(data T) {
	newNode := &Node[T]{data: data}

	if ll.tail == nil {

		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.prev = ll.tail
		ll.tail.next = newNode
		ll.tail = newNode
	}

	ll.length += 1
}

// PrintList() -> prints linked list head to tail
func (ll *LinkedList[T]) PrintList() {
	current := ll.head

	for current != nil {

		if current.prev != nil {
			fmt.Print(" <--> ")
		}

		fmt.Print(current.data)
		current = current.next
	}
	fmt.Println()
}

// Data() -> returns a pointer to head.data
func (ll LinkedList[T]) Data() *T {
	if ll.head == nil {
		return nil
	}

	return &ll.head.data
}

// DeleteHead() -> deletes front node if there is a node to delete
func (ll *LinkedList[T]) DeleteHead() {
	if ll.head != nil {
		if ll.head == ll.tail {
			ll.head = nil
			ll.tail = nil
		} else {
			ll.head = ll.head.next
			ll.head.prev = nil

		}

		ll.length -= 1
	}
}

// DeleteTail() -> removes end node
func (ll *LinkedList[T]) DeleteTail() {
	if ll.tail != nil {
		if ll.head == ll.tail {
			ll.head = nil
			ll.tail = nil
		} else {
			ll.tail = ll.tail.prev
			ll.tail.next = nil

		}

		ll.length -= 1
	}
}
