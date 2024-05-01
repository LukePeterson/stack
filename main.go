package main

import "fmt"

type Stack struct {
	elements []int
}

func (s *Stack) push(value int) {
	s.elements = append(s.elements, value)
}

func (s *Stack) pop() int {
	poppedValue := s.elements[len(s.elements)-1]
	s.elements = s.elements[0 : len(s.elements)-1]
	return poppedValue
}

func main() {

	stack := Stack{}

	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(5)
	stack.push(8)

	fmt.Printf("stack: %v\n", stack)

	fmt.Printf("stack.pop(): %v\n", stack.pop())
	fmt.Printf("stack: %v\n", stack)

	fmt.Printf("stack.pop(): %v\n", stack.pop())
	fmt.Printf("stack: %v\n", stack)

	fmt.Printf("stack.pop(): %v\n", stack.pop())
	fmt.Printf("stack: %v\n", stack)

	fmt.Printf("stack.pop(): %v\n", stack.pop())
	fmt.Printf("stack: %v\n", stack)

	fmt.Printf("stack.pop(): %v\n", stack.pop())
	fmt.Printf("stack: %v\n", stack)
}
