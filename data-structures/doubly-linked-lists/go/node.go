package doubly

// Node in a double Linked List that stores runes
type Node struct {
	Val      rune
	NextNode *Node
	PrevNode *Node
}
