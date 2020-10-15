package main

type Hashable interface {
	Equal() bool
	Hash() int
}

type HashTable interface {
	Hashable
}
