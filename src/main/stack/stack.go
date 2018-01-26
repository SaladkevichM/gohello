package main

import (
	//"fmt"
	"errors"
)

var ErrStackIsEmpty = errors.New("stack is empty")

type LinkedList struct {
	Tail  *SItem
	Head  *SItem
	Items []SItem
}

type SItem struct {
	Value string
}

type Stack struct {
	Length int
	Data   *LinkedList
}

func (s *Stack) pop() (*SItem, error) {
	if s.Length == 0 {
		return nil, ErrStackIsEmpty
	}
	return &SItem{}, nil
}

func (s *Stack) push(item SItem) error {

	return nil
}

func main() {

}
