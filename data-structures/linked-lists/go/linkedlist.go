package linkedlist

import (
	"errors"
)

var (
	valueError = errors.New("value error: could not be found")
	indexError = errors.New("index error: index out of range")
)

type LinkedList struct {
	FirstNode *Node
}

func (ll *LinkedList) New(n *Node) {
	ll.FirstNode = n
}

// Read takes an index and returns the value at that index Node.
func (ll *LinkedList) Read(i int) rune {
	currNode := ll.FirstNode
	currIndex := 0

	for currIndex < i {
		// if nextNode is nil then the end of the LinkedList is reached and
		// the value of the current Node should be returned.
		if currNode.nextNode == nil {
			return currNode.val
		}
		currNode = currNode.nextNode
		currIndex++
	}

	return currNode.val
}

// Search expects a value of type `rune` and returns the memory address of the
// first Node it was found in.
//
//Returns reference to Node if the value is found, an error if it was not.
func (ll *LinkedList) Search(val rune) (*Node, error) {
	currNode := ll.FirstNode

	for {
		if currNode.val == val {
			return currNode, nil
		}
		if currNode.nextNode == nil {
			break
		}
		currNode = currNode.nextNode
	}

	return nil, valueError
}

// Insert expects an index and a value and creates a new Node at the given index.
// If the given index is out of range, the Node is create at the end of the
// LinkedList.
// Returns an error.
func (ll *LinkedList) Insert(i int, val rune) error {
	newNode := Node{val, nil}

	if i == 0 {
		newNode.nextNode = ll.FirstNode
		ll.FirstNode = &newNode

		return nil
	} else {
		currNode := ll.FirstNode
		currIndex := 0
		for currIndex < i-1 {
			// if Node at currIndex has no nextNode, then we have reached the
			// end of the LinkedList, so newNode can be appended.
			if currNode.nextNode == nil {
				currNode.nextNode = &newNode

				return nil
			}
			currNode = currNode.nextNode
			currIndex++
		}
		newNode.nextNode = currNode.nextNode
		currNode.nextNode = &newNode

	}

	return nil
}

// Delete takes an index and deletes the node at the given index.
// To delete the last element, use DeleteLast.
// Returns an error.
func (ll *LinkedList) DeleteAt(i int) error {
	if i == 0 {
		ll.FirstNode = ll.FirstNode.nextNode

		return nil
	}

	currNode := ll.FirstNode
	currIndex := 0

	for currIndex < i-1 {
		if currNode.nextNode.nextNode == nil {
			// reached end of list
			return indexError
		}
		currNode = currNode.nextNode
		currIndex++
	}

	succeedingNode := currNode.nextNode.nextNode
	currNode.nextNode = succeedingNode

	return nil
}
