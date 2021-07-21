package doublyqueue

import (
	"testing"

	db "github.com/heyitschun/cs-essentials/data-structures/doubly-linked-lists/go"
)

var readNodes = []db.Node{
	{Val: 'a', NextNode: nil, PrevNode: nil},
	{Val: 'b', NextNode: nil, PrevNode: nil},
	{Val: 'c', NextNode: nil, PrevNode: nil},
}

type returnCases struct {
	dll    []db.Node
	answer rune
}

var readFirstTests = []returnCases{
	{[]db.Node{}, ' '},    // if list is empty, expect empty rune
	{readNodes[0:1], 'a'}, // if there is one element, expect the first
	{readNodes, 'a'},      // if there are multiple elements, expect the first
}

func TestReadFirst(t *testing.T) {
	for _, rf := range readFirstTests {
		q := Queue{db.DoublyLinkedList{}}
		for _, n := range rf.dll {
			q.Enqueue(n.Val)
		}
		ans := q.Head()
		if ans != rf.answer {
			t.Error(
				"ReadFirst for", rf.dll,
				"expected", rf.answer,
				"got", ans,
			)
		}
	}
}

var readLastTests = []returnCases{
	{[]db.Node{}, ' '},    // if list is empty, expect empty rune
	{readNodes[0:1], 'a'}, // if there is one element, expect the first
	{readNodes, 'c'},      // if there are multiple elements, expect the last
}

func TestReadLast(t *testing.T) {
	for _, rf := range readLastTests {
		q := Queue{db.DoublyLinkedList{}}
		for _, n := range rf.dll {
			q.Enqueue(n.Val)
		}
		ans := q.Tail()
		if ans != rf.answer {
			t.Error(
				"ReadLast for", rf.dll,
				"expected", rf.answer,
				"got", ans,
			)
		}
	}
}

var dequeueTests = []returnCases{
	{[]db.Node{}, ' '},    // if list is empty, expect empty rune
	{readNodes[0:1], 'a'}, // if there is one element, expect the first
	{readNodes, 'a'},      // if there are multiple elements, expect the last
}

func TestDequeue(t *testing.T) {
	for _, rf := range dequeueTests {
		q := Queue{db.DoublyLinkedList{}}
		for _, n := range rf.dll {
			q.Enqueue(n.Val)
		}
		ans := q.Dequeue()
		if ans != rf.answer {
			t.Error(
				"ReadLast for", rf.dll,
				"expected", rf.answer,
				"got", ans,
			)
		}
	}
}
