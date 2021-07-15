package stack

// Stack that stores integers
type Stack struct {
	elements []int
}

// Push adds an elements to the stack
func (s *Stack) Push(i int) {
	s.elements = append(s.elements, i)
}

// Pop removes the first element from the stack and returns it
func (s *Stack) Pop() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	popped := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return popped, true
}

// Read returns the element on the top of the stack
func (s *Stack) Read() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	return s.elements[len(s.elements)-1], true
}

// Stack that stores integers
type StackString struct {
	elements []string
}

// Push adds an elements to the stack
func (s *StackString) Push(i string) {
	s.elements = append(s.elements, i)
}

// Pop removes the first element from the stack and returns it
func (s *StackString) Pop() (string, bool) {
	if len(s.elements) == 0 {
		return "", false
	}
	popped := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return popped, true
}

// Read returns the element on the top of the stack
func (s *StackString) Read() (string, bool) {
	if len(s.elements) == 0 {
		return "", false
	}
	return s.elements[len(s.elements)-1], true
}
