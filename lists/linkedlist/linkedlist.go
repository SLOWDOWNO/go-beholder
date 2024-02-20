package linkedlist

import (
	"fmt"
	"go-beholder/lists"
	"go-beholder/utils"
	"slices"
	"strings"
)

var _ lists.List[int] = (*List[int])(nil)

// List represents a singly linked list data structure.
type List[T comparable] struct {
	first *node[T] // Pointer to the first node in the list.
	tail  *node[T] // Pointer to the last node in the list.
	size  int      // Number of elements in the list.
}

// node represents a node in the linked list.。
type node[T comparable] struct {
	value T        // Value stored in the node.
	next  *node[T] // Pointer to the next node in the list.
}

// isValidIndex checks if the given index is valid for the list.
func (list *List[T]) isValidIndex(index int) bool {
	return index >= 0 && index < list.size
}

// New creates and returns a new List instance with the provided values.
func New[T comparable](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Add adds the specified values to the end of the list.
func (list *List[T]) Add(values ...T) {
	for _, value := range values {
		newNode := &node[T]{value: value}
		if list.size == 0 {
			list.first = newNode
			list.tail = newNode
		} else {
			list.tail.next = newNode
			list.tail = newNode
		}
		list.size++
	}
}

// Add adds the specified values to the head of the list.
func (list *List[T]) AddToHead(values ...T) {
	for v := len(values) - 1; v >= 0; v-- {
		newNode := &node[T]{value: values[v], next: list.first}
		list.first = newNode
		if list.size == 0 {
			list.tail = newNode
		}
		list.size++
	}
}

// Get retrieves the value at the specified index in the list.
func (list *List[T]) Get(index int) (T, bool) {
	if !list.isValidIndex(index) {
		var t T
		return t, false
	}
	node := list.first
	for e := 0; e != index; e, node = e+1, node.next {
	}
	return node.value, true
}

// Remove removes the value at the specified index from the list.
func (list *List[T]) Remove(index int) {
	if !list.isValidIndex(index) {
		return
	}

	if list.size == 1 {
		list.Clear()
		return
	}

	var beforeNode *node[T]
	node := list.first
	for e := 0; e != index; e, node = e+1, node.next {
		beforeNode = node
	}

	if node == list.first {
		list.first = node.next
	}
	if node == list.tail {
		list.tail = beforeNode
	}
	if beforeNode != nil {
		beforeNode.next = node.next
	}

	node = nil

	list.size--
}

// Contains checks if the list contains all of the specified values.
func (list *List[T]) Contains(values ...T) bool {
	if len(values) == 0 {
		return true
	}
	if list.size == 0 {
		return false
	}

	for _, value := range values {
		found := false
		for cur := list.first; cur != nil; cur = cur.next {
			if cur.value == value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Sort sorts the elements of the list using the provided comparator function.
func (list *List[T]) Sort(comparator utils.Comparator[T]) {
	if list.size < 2 {
		return
	}
	values := list.Values()

	slices.SortFunc(values, comparator)
	list.Clear()
	list.Add(values...)
}

// Swap swaps the values at the specified indices in the list.
func (list *List[T]) Swap(i, j int) {
	if list.isValidIndex(i) && list.isValidIndex(j) && i != j {
		var node1, node2 *node[T]
		for e, cur := 0, list.first; node1 == nil || node2 == nil; e, cur = e+1, cur.next {
			switch e {
			case i:
				node1 = cur
			case j:
				node2 = cur
			}
		}
		node1.value, node2.value = node2.value, node1.value
	}
}

// Insert inserts the specified values at the specified index in the list.
func (list *List[T]) Insert(index int, values ...T) {
	if !list.isValidIndex(index) {
		if index == list.size {
			list.Add(values...)
		}
		return
	}

	list.size += len(values)

	var beforeNode *node[T]
	cur := list.first
	for e := 0; e != index; e, cur = e+1, cur.next {
		beforeNode = cur
	}

	if cur == list.first {
		oldNextElement := list.first
		for i, value := range values {
			newNode := &node[T]{value: value}
			if i == 0 {
				list.first = newNode
			} else {
				beforeNode.next = newNode
			}
			beforeNode = newNode
		}
		beforeNode.next = oldNextElement
	} else {
		oldNextElement := beforeNode.next
		for _, value := range values {
			newNode := &node[T]{value: value}
			beforeNode.next = newNode
			beforeNode = newNode
		}
		beforeNode.next = oldNextElement
	}
}

// Set sets the value at the specified index in the list.
func (list *List[T]) Set(index int, value T) {
	if !list.isValidIndex(index) {
		if list.size == index {
			list.Add(value)
		}
		return
	}
	cur := list.first
	for i := 0; i != index; i++ {
		cur = cur.next
	}
	cur.value = value
}

// Empty checks if the list is empty.
func (list *List[T]) Empty() bool {
	return list.size == 0
}

// Size returns the number of elements in the list.
func (list *List[T]) Size() int {
	return list.size
}

// Clear removes all elements from the list.
func (list *List[T]) Clear() {
	list.size = 0
	list.first = nil
	list.tail = nil
}

// Values returns a slice containing all the values in the list.
func (list *List[T]) Values() []T {
	values := make([]T, list.size)
	for i, cur := 0, list.first; cur != nil; i, cur = i+1, cur.next {
		values[i] = cur.value
	}
	return values
}

// IndexOf returns the index of the first occurrence of the specified value in the list, or -1 if not found.
func (list *List[T]) IndexOf(value T) int {
	if list.size == 0 {
		return -1
	}
	for index, node := range list.Values() {
		if node == value {
			return index
		}
	}
	return -1
}

// String returns a string representation of the list.
func (list *List[T]) String() string {
	str := "[LinkedList]\n"
	values := []string{}
	for cur := list.first; cur != nil; cur = cur.next {
		values = append(values, fmt.Sprintf("%v", cur.value))
	}
	str += strings.Join(values, ",")
	return str
}
