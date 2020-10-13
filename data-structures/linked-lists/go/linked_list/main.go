package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type SinglyLinkedList struct {
	Length int
	Head   *Node
}

func initSLL() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (sll *SinglyLinkedList) Append(data int) {
	if sll.Head == nil {
		n := Node{Data: data, Next: nil}
		sll.Head = &n
		return
	}

	curr := sll.Head
	for curr.Next != nil {
		curr = curr.Next
	}

	n := Node{Data: data, Next: nil}
	curr.Next = &n
}

func (sll *SinglyLinkedList) Prepend(data int) {
	n := Node{Data: data, Next: nil}
	new_head := &n
	new_head.Next = sll.Head
	sll.Head = new_head
}

func (sll *SinglyLinkedList) Delete(data int) {
	if sll.Head == nil {
		return
	}

	if sll.Head.Data == data {
		sll.Head = sll.Head.Next
	}

	curr := sll.Head
	for curr.Next != nil {
		if curr.Next.Data == data {
			curr.Next = curr.Next.Next
			return
		}
		curr = curr.Next
	}
}

func (sll *SinglyLinkedList) PrintData() {
	curr := sll.Head

	fmt.Println("###")
	for curr.Next != nil {
		fmt.Printf("%d\n", curr.Data)
		curr = curr.Next
	}
}

func main() {
	sll := initSLL()
	sll.Append(4)
	sll.Append(5)
	sll.Append(6)
	sll.Append(7)

	sll.PrintData()

	sll.Prepend(8)

	sll.PrintData()

	sll.Delete(6)

	sll.PrintData()
}
