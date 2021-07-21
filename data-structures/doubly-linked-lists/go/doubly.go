package doubly

import "errors"

var (
	doublyInitError = errors.New("value error: cannot initialize doubly linked list with nil")
)

//DoublyLinkedList currently only has functionality for use in a Queue as it
//only allows adding to the end and removing from the front.
//
//It gives access to the first and last elements of the linked list only!
type DoublyLinkedList struct {
	FirstNode *Node
	LastNode  *Node
}

//New creates a new DoublyLinkedList
func (dll *DoublyLinkedList) New(initialNode *Node) error {
	//cannot accept nil as initialNode
	if initialNode == nil {
		return doublyInitError
	}

	dll.FirstNode = initialNode
	dll.LastNode = initialNode

	return nil
}

//InsertAtEnd adds a new node with value `val` at the end of the list.
func (dll *DoublyLinkedList) InsertAtEnd(val rune) {
	newNode := Node{val, nil, nil}

	//if the list is empty: the new node becomes both the first and last node
	if dll.FirstNode == nil {
		dll.FirstNode = &newNode
		dll.LastNode = &newNode
	} else {
		newNode.PrevNode = dll.LastNode
		dll.LastNode.NextNode = &newNode
		dll.LastNode = &newNode
	}
}

//PopAtFront removes the first node from the linked list and returns its value.
//Returns `rune`.
func (dll *DoublyLinkedList) PopAtFront() rune {
	if dll.FirstNode == nil {
		return ' '
	}

	removedNode := dll.FirstNode
	dll.FirstNode = dll.FirstNode.NextNode

	return removedNode.Val
}

//ReadFirst returns the value of first node in the linked list, without removing it.
func (dll *DoublyLinkedList) ReadFirst() rune {
	if dll.FirstNode == nil {
		return ' '
	}

	return dll.FirstNode.Val
}

//ReadLast returns the value of last node in the linked list, without removing it.
func (dll *DoublyLinkedList) ReadLast() rune {
	if dll.LastNode == nil {
		return ' '
	}

	return dll.LastNode.Val
}
