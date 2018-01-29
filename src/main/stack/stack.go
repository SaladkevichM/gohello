package main

import (
	"errors"
	"fmt"
)

var ErrEmptyList = errors.New("List is empty")
var ErrElementIsNil = errors.New("Element is nil")

type Item struct {
	value interface{}
}

type LinkedList struct {
	Length int
	Tail   *Item
	Head   *Item
	Items  *[]Item
}

func (i Item) String() string {
	return fmt.Sprintf("value - %v", i.value.(string))
}

func (l LinkedList) String() string {
	return fmt.Sprintf("head: %v, tail: %v, length: %v", l.Head, l.Tail, l.Length)
}

func NewList() *LinkedList {
	l := new(LinkedList)
	slice := make([]Item, 0, 0)
	l.Items = &slice
	return l
}

func (list *LinkedList) pop() (*Item, error) {
	if list.Length == 0 {
		return nil, ErrEmptyList
	}

	// get last element reference
	popped := list.Tail

	// list size
	list.Length = list.Length - 1

	// remove last element
	slice := *list.Items
	update := slice[0 : len(slice)-1]
	list.Items = &update

	// update references
	if len(update) == 1 {
		list.Head = &update[len(update)]
		list.Tail = &update[len(update)]
	} else {
		list.Tail = &update[len(update)]
	}

	return popped, nil
}

func (list *LinkedList) push(element Item) error {

	if &element == nil {
		return ErrElementIsNil
	}

	// list size
	list.Length = list.Length + 1

	// push element in back
	slice := *list.Items
	update := make([]Item, len(slice)+1, cap(slice)+1)
	copy(slice, update)
	update[len(update)] = element
	list.Items = &update

	// update references
	if len(update) == 1 {
		list.Head = &element
		list.Tail = &element
	} else {
		list.Tail = &element
	}

	return nil
}

func main() {

	stack := NewList()
	fmt.Println(stack)

	e1 := Item{value: "first"}
	e2 := Item{value: "second"}
	e3 := Item{value: "third"}

	stack.push(e1)
	fmt.Println(stack)

	stack.push(e2)
	stack.push(e3)
	fmt.Println(stack)

}
