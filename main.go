package main

import (
	"fmt"
)

type Stack struct {
	elements []int
}

func (s *Stack) push(value int) {
	s.elements = append(s.elements, value)
}

func (s *Stack) pop() (int, error) {
	if len(s.elements) == 0 {
		return 0, fmt.Errorf("cannot pop() element from an empty stack")
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[0 : len(s.elements)-1]
	return element, nil
}

func (s *Stack) peek() (int, error) {
	if s.empty() {
		return 0, fmt.Errorf("cannot peek at empty stack")
	}

	return s.elements[len(s.elements)-1], nil
}

func (s *Stack) empty() bool {
	return len(s.elements) == 0
}

func (s *Stack) size() int {
	return len(s.elements)
}

func main() {

	stack := Stack{}

	if element, err := stack.peek(); err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Printf("element: %v\n", element)
	}

	fmt.Printf("stack.empty(): %v\n", stack.empty())

	stack.push(1)
	stack.push(2)

	if element, err := stack.peek(); err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Printf("element: %v\n", element)
	}

	stack.push(3)
	stack.push(5)
	stack.push(8)

	if element, err := stack.peek(); err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Printf("element: %v\n", element)
	}

	fmt.Printf("stack: %v\n", stack)
	fmt.Printf("stack.empty(): %v\n", stack.empty())
	fmt.Printf("stack.size(): %v\n", stack.size())

	if element, err := stack.pop(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("element: %v\n", element)
	}

	fmt.Printf("stack.size(): %v\n", stack.size())

	if element, err := stack.pop(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("element: %v\n", element)
	}

	fmt.Printf("stack.size(): %v\n", stack.size())
}
