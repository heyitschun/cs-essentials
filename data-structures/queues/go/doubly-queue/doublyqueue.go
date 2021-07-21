package doublyqueue

import (
	db "github.com/heyitschun/cs-essentials/data-structures/doubly-linked-lists/go"
)

type Queue struct {
	Data db.DoublyLinkedList
}

func (q *Queue) Enqueue(val rune) {
	//if first node is nil it means the queue is empty and a new queue should be
	//created
	if q.Data.FirstNode == nil {
		q.Data.New(&db.Node{
			Val:      val,
			NextNode: nil,
			PrevNode: nil,
		})
	} else {
		q.Data.InsertAtEnd(val)
	}
}

func (q *Queue) Dequeue() rune {
	return q.Data.PopAtFront()
}

func (q *Queue) Head() rune {
	return q.Data.ReadFirst()
}

func (q *Queue) Tail() rune {
	return q.Data.ReadLast()
}
