package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []int
}

// Push the way we have to add items to the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop the way we have to retrive items from the stack
func (s *Stack) Pop() (int, error) {
	if len(s.items) > 0 {
		indexToRemove := len(s.items) - 1
		itemToRemove := s.items[indexToRemove]
		s.items = s.items[0:indexToRemove]

		return itemToRemove, nil
	}

	return -1, errors.New("the stack does not have items")
}

func main(){
	s := Stack{}
	PrintWithCoolFormat(s, "Initial state")
	s.Push(100)
	PrintWithCoolFormat(s, "100 pushed")
	s.Push(200)
	PrintWithCoolFormat(s, "200 pushed")
	s.Push(300)
	PrintWithCoolFormat(s, "300 pushed")

	item, _ := s.Pop()
	PrintWithCoolFormat(s, "Item Pop: " + fmt.Sprint(item))
	item, _ = s.Pop()
	PrintWithCoolFormat(s, "Item Pop: " + fmt.Sprint(item))
	item, _ = s.Pop()
	PrintWithCoolFormat(s, "Item Pop: " + fmt.Sprint(item))
}
