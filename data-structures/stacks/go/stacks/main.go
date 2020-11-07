package main

import (
	"errors"
	"fmt"
	"log"
)

const arrSize = 10

type Stack struct {
	Top  int
	Data [arrSize]int
}

func (s *Stack) Push(i int) error {
	if s.Top == len(s.Data) {
		return errors.New("Stack overflow")
	}

	s.Data[s.Top] = i
	s.Top += 1

	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.Top == 0 {
		return 0, errors.New("Empty stack")
	}

	i := s.Data[s.Top-1]
	s.Top -= 1

	return i, nil
}

func (s *Stack) Empty() bool {
	return s.Top == 0
}

func main() {
	s := Stack{}
	err := s.Push(2)
	if err != nil {
		log.Fatal(err)
	}

	var i int = 0
	i, err = s.Pop()
	if err != nil {
		log.Fatal(err)
	}

	empty := s.Empty()
	fmt.Println(i, s)
	fmt.Println(empty)
}
