package stack

import (
	"fmt"
	"go-beholder/lists/linkedlist"
	"go-beholder/stacks"
	"strings"
)

var _ stacks.Stack[int] = (*Stack[int])(nil)

type Stack[T comparable] struct {
	list *linkedlist.List[T]
}

func New[T comparable]() *Stack[T] {
	return &Stack[T]{list: &linkedlist.List[T]{}}
}

func (stack *Stack[T]) Push(value T) {
	stack.list.AddToHead(value)
}

func (stack *Stack[T]) Pop() (value T, ok bool) {
	value, ok = stack.list.Get(0)
	stack.list.Remove(0)
	return value, ok
}

func (stack *Stack[T]) Peek() (value T, ok bool) {
	return stack.list.Get(0)
}

func (stack *Stack[T]) Empty() bool {
	return stack.list.Empty()
}

func (stack *Stack[T]) Size() int {
	return stack.list.Size()
}

func (stack *Stack[T]) Clear() {
	stack.list.Clear()
}

func (stack *Stack[T]) Values() []T {
	return stack.list.Values()
}

func (stack *Stack[T]) String() string {
	str := "[LinkedListStack]\n"
	values := []string{}
	for _, value := range stack.list.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}
