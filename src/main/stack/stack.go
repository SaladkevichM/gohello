package main

import (
	"errors"
	"fmt"
	"math"
)

// ErrEmptyStack describes empty stack error
var ErrEmptyStack = errors.New("Stack is empty")

// ErrStackIsFull describes full stack error
var ErrStackIsFull = errors.New("Stack is full")

// ErrElementIsNil describes that element is nil
var ErrElementIsNil = errors.New("Element is nil")

// Item holds element of the stack
type Item struct {
	value string
}

// Stack is a data structure
type Stack struct {
	Length int
	Tail   *Item
	Head   *Item
	Items  []Item
}

func (i Item) String() string {
	return fmt.Sprintf("value - %v", i.value)
}

func (s Stack) String() string {
	print := fmt.Sprintf("Head: %v; Tail: %v; Length: %v", s.Head, s.Tail, s.Length)
	print += "\nElements:\n"
	for i, v := range s.Items {
		print += fmt.Sprintf("%d - %v\n", i, v.value)
	}
	return print
}

// NewStack is like a constructor
func NewStack() *Stack {
	s := new(Stack)
	s.Items = make([]Item, 0, 0)
	return s
}

// Pop element from stack
func (s *Stack) Pop() (*Item, error) {
	if s.Length == 0 {
		return nil, ErrEmptyStack
	}

	// get last element reference
	popped := s.Tail

	// stack size
	s.Length = s.Length - 1

	// remove last element
	update := make([]Item, s.Length, s.Length)
	copy(update, s.Items)
	s.Items = update

	// update references
	if s.Length == 1 {
		s.Head = &update[s.Length-1]
		s.Tail = &update[s.Length-1]
	}
	if s.Length > 1 {
		s.Tail = &update[s.Length-1]
	}
	if s.Length == 0 {
		s.Head = nil
		s.Tail = nil
	}

	return popped, nil
}

// Push element to stack
func (s *Stack) Push(element Item) error {

	if &element == nil {
		return ErrElementIsNil
	}

	if s.Length == math.MaxInt32 {
		return ErrStackIsFull
	}

	// stack size
	s.Length = s.Length + 1

	// push element in back
	slice := s.Items

	update := make([]Item, s.Length, s.Length)
	if s.Length > 0 {
		copy(update, slice)
	}

	update[s.Length-1] = element
	s.Items = update

	// update references
	if s.Length == 1 {
		s.Head = &element
		s.Tail = &element
	} else {
		s.Tail = &element
	}

	return nil
}

func main() {

	e1 := Item{value: "first"}
	e2 := Item{value: "second"}
	e3 := Item{value: "third"}

	stack := NewStack()
	fmt.Println(stack)

	fmt.Println("")
	fmt.Println("start pushing")

	stack.Push(e1)
	fmt.Println(stack.Length)

	stack.Push(e2)
	fmt.Println(stack.Length)

	stack.Push(e3)
	fmt.Println(stack.Length)

	fmt.Println("")
	fmt.Println("start poping")

	p1, _ := stack.Pop()
	fmt.Println(p1.value)

	p2, _ := stack.Pop()
	fmt.Println(p2.value)

	p3, _ := stack.Pop()
	fmt.Println(p3.value)

	fmt.Println("")
	fmt.Println("start pushing")

	stack.Push(e1)
	stack.Push(e2)
	stack.Push(e3)

	fmt.Println(stack)

}
