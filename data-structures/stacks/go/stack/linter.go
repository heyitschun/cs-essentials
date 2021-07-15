package stack

import (
	"errors"
)

// helper function
func contains(s []string, e string) bool {
	for _, char := range s {
		if char == e {
			return true
		}
	}

	return false
}

var (
	SyntaxError1 error = errors.New(
		"Syntax Error (#1): missing corresponding closing brace",
	)
	SyntaxError2 error = errors.New(
		"Syntax Error (#2): missing corresponding opening brace",
	)
	SyntaxError3 error = errors.New(
		"Syntax Error (#3): mismatched braces",
	)
)

func BraceChecker(s string) error {
	opening := []string{"(", "[", "{"}
	closing := []string{")", "]", "}"}
	matches := make(map[string]string)
	for i, b := range opening {
		matches[closing[i]] = b
	}
	var split []string
	for _, x := range s {
		split = append(split, string(x))
	}
	stack := StackString{}
	for _, char := range split {
		// opening brace
		if contains(opening, char) {
			stack.Push(char)
		} else if contains(closing, char) {
			lastBrace, ok := stack.Pop()
			if !ok {
				return SyntaxError2
			}
			if matches[char] != lastBrace {
				return SyntaxError3
			}
		}
	}
	char, _ := stack.Pop()
	if char != "" {
		return SyntaxError1
	}

	return nil
}
