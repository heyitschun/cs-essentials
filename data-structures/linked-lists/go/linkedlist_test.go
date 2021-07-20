package linkedlist

import (
	"log"
	"testing"
)

type newLinkedListTest struct {
	firstNode *Node
	ans       rune
}

var (
	node1 = Node{'a', &node2}
	node2 = Node{'b', &node3}
	node3 = Node{'c', nil}
	node4 = Node{'d', nil}
)

var nllTests = []newLinkedListTest{
	{&node1, 'a'},
}

func TestNewLinkedList(t *testing.T) {
	for _, nll := range nllTests {
		ll := LinkedList{nll.firstNode}
		ans := ll.FirstNode.val
		if ans != nll.ans {
			t.Error(
				"For", nll.firstNode,
				"expected", nll.ans,
				"got", ans,
			)
		}
	}
}

// Read
type readTest struct {
	targetNode *Node
}

var readTests = []readTest{
	{&node3},
}

func TestRead(t *testing.T) {
	for _, rt := range readTests {
		ll := LinkedList{&node1}
		// very ugly, but for there is no insertion yet
		ll.FirstNode.nextNode = &node2
		ll.FirstNode.nextNode.nextNode = &node3
		// TODO: update with more test cases to handle end of list
		ans := ll.Read(2)
		if ans != rt.targetNode.val {
			t.Error(
				"For", rt.targetNode,
				"expected Read() ->", rt.targetNode.val,
				"got", ans,
			)
		}
	}
}

// Search
type searchTest struct {
	val rune
	ans *Node
	err error
}

var searchTests = []searchTest{
	{'b', &node2, nil},
	{'c', &node3, nil},
	{'d', &node3, valueError},
}

func TestSearch(t *testing.T) {
	for _, st := range searchTests {
		ll := LinkedList{&node1}
		// very ugly, but for there is no insertion
		ll.FirstNode.nextNode = &node2
		ll.FirstNode.nextNode.nextNode = &node3
		ans, err := ll.Search(st.val)
		if ans != st.ans && err != st.err {
			t.Error(
				"For", string(st.ans.val),
				"expected Search() ->", st.val,
				"got", ans,
			)
		}
	}
}

// Insert
type insertTest struct {
	val   rune
	index int
	err   error
}

var insertTests = []insertTest{
	{'d', 1, nil},
	{'d', 0, nil},
	{'d', 10, nil},
}

func TestInsert(t *testing.T) {
	for _, it := range insertTests {
		ll := LinkedList{&node1}
		ll.FirstNode.nextNode = &node2
		ll.FirstNode.nextNode.nextNode = &node3
		// NOTE: add error handling
		err := ll.Insert(it.index, it.val)
		if err != nil {
			log.Fatal(err)
		}
		ans := ll.Read(it.index)
		if ans != it.val {
			t.Errorf(
				"For ll.Insert(%d, %s) expected %s got %s",
				it.index,
				string(it.val),
				string(it.val),
				string(ans),
			)
		}
	}
}

// Insert
type deleteTest struct {
	index        int   // index to delete
	err          error // do not allow auto deletion of last index(?)
	newIndexNode *Node
}

var deleteTests = []deleteTest{
	{1, nil, &node3},     // delete second node
	{4, indexError, nil}, // delete non existing node
}

func TestDeleteAt(t *testing.T) {
	for _, dt := range deleteTests {
		ll := LinkedList{&node1}
		ll.FirstNode.nextNode = &node2
		ll.FirstNode.nextNode.nextNode = &node3
		err := ll.DeleteAt(dt.index)
		if err != nil {
			if err != dt.err {
				t.Error(
					"For", dt.index,
					"expected DeleteAt() ->", dt.err,
					"got", err,
				)
			}
		} else {
			newValAtIndex := ll.Read(dt.index)
			if newValAtIndex != dt.newIndexNode.val {
				t.Error(
					"For", dt.index,
					"expected DeleteAt() ->", string(dt.newIndexNode.val),
					"got", string(newValAtIndex),
				)
			}
		}
	}
}
