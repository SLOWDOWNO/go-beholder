package linkedlist

import (
	"fmt"
	"go-beholder/utils"
	"strings"
)

// List represents a singly linked list data structure.
type List struct {
	first *node // Pointer to the first node in the list.
	tail  *node // Pointer to the last node in the list.
	size  int   // Number of elements in the list.
}

// node represents a node in the linked list.。
type node struct {
	value interface{} // Value stored in the node.
	next  *node       // Pointer to the next node in the list.
}

// isValidIndex checks if the given index is valid for the list.
func (list *List) isValidIndex(index int) bool {
	return index >= 0 && index < list.size
}

// New creates and returns a new List instance with the provided values.
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Add adds the specified values to the end of the list.
func (list *List) Add(values ...interface{}) {
	for _, value := range values {
		newNode := &node{value: value}
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
func (list *List) AddToHead(values ...interface{}) {
	for v := len(values) - 1; v >= 0; v-- {
		newNode := &node{value: values[v], next: list.first}
		list.first = newNode
		if list.size == 0 {
			list.tail = newNode
		}
		list.size++
	}
}

// Get retrieves the value at the specified index in the list.
func (list *List) Get(index int) (interface{}, bool) {
	if !list.isValidIndex(index) {
		return nil, false
	}
	node := list.first
	for e := 0; e != index; e, node = e+1, node.next {
	}
	return node.value, true
}

// Remove removes the value at the specified index from the list.
func (list *List) Remove(index int) {
	if !list.isValidIndex(index) {
		return
	}

	if list.size == 1 {
		list.Clear()
		return
	}

	var beforeNode *node
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
func (list *List) Contains(values ...interface{}) bool {
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
func (list *List) Sort(comparator utils.Comparator) {
	if list.size < 2 {
		return
	}
	values := list.Values()
	utils.Sort(values, comparator)
	list.Clear()
	list.Add(values...)
}

// Swap swaps the values at the specified indices in the list.
func (list *List) Swap(i, j int) {
	if list.isValidIndex(i) && list.isValidIndex(j) && i != j {
		var node1, node2 *node
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
func (list *List) Insert(index int, values ...interface{}) {
	if !list.isValidIndex(index) {
		if index == list.size {
			list.Add(values...)
		}
		return
	}

	list.size += len(values)

	var beforeNode *node
	cur := list.first
	for e := 0; e != index; e, cur = e+1, cur.next {
		beforeNode = cur
	}

	if cur == list.first {
		oldNextElement := list.first
		for i, value := range values {
			newNode := &node{value: value}
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
			newNode := &node{value: value}
			beforeNode.next = newNode
			beforeNode = newNode
		}
		beforeNode.next = oldNextElement
	}
}

// Set sets the value at the specified index in the list.
func (list *List) Set(index int, value interface{}) {
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
func (list *List) Empty() bool {
	return list.size == 0
}

// Size returns the number of elements in the list.
func (list *List) Size() int {
	return list.size
}

// Clear removes all elements from the list.
func (list *List) Clear() {
	list.size = 0
	list.first = nil
	list.tail = nil
}

// Values returns a slice containing all the values in the list.
func (list *List) Values() []interface{} {
	values := make([]interface{}, list.size)
	for i, cur :=

		0, list.first; cur != nil; i, cur = i+1, cur.next {
		values[i] = cur.value
	}
	return values
}

// IndexOf returns the index of the first occurrence of the specified value in the list, or -1 if not found.
func (list *List) IndexOf(value interface{}) int {
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
func (list *List) String() string {
	str := "[LinkedList]\n"
	values := []string{}
	for cur := list.first; cur != nil; cur = cur.next {
		values = append(values, fmt.Sprintf("%v", cur.value))
	}
	str += strings.Join(values, ",")
	return str
}
