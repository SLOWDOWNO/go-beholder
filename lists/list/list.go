package list

import (
	"fmt"
	"go-beholder/lists"
	"go-beholder/utils"
	"strings"
)

// _ is an assertion to ensure that *List implements the lists.List interface.
var _ lists.List = (*List)(nil)

type List struct {
	elements []interface{}
	size     int
}

// isValidIndex checks if the provided index is within the bounds of the list.
func (list *List) isValidIndex(idx int) bool {
	return idx >= 0 && idx < list.size
}

// New creates a new List and adds the provided values, if any.
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Values returns a new slice that contains all elements in the list.
func (list *List) Values() []interface{} {
	newElements := make([]interface{}, list.size)
	copy(newElements, list.elements[:list.size])
	return newElements
}

// Add appends the provided values to the end of the list.
func (list *List) Add(values ...interface{}) {
	for _, value := range values {
		list.elements = append(list.elements, value)
		list.size++
	}
}

// Get retrieves the element at the specified index from the list.
// It returns the element and true if the index is valid, and nil and false otherwise.
func (list *List) Get(index int) (interface{}, bool) {

	if list.isValidIndex(index) {
		return list.elements[index], true
	}
	return nil, false
}

// Remove deletes the element at the specified index from the list.
// If the index is not valid, the function does nothing.
func (list *List) Remove(index int) {
	if !list.isValidIndex(index) {
		return
	}
	list.elements = append(list.elements[:index], list.elements[index+1:]...)
	list.size--
}

// Contains checks if the list contains all the provided values.
// It returns true if all values are found in the list, and false otherwise.
func (list *List) Contains(values ...interface{}) bool {
	isFound := false
	for _, target := range values {
		for idx := 0; idx < list.size; idx++ {
			if list.elements[idx] == target {
				isFound = true
				break
			}
		}
		if !isFound {
			return false
		}
	}
	return true
}

// Swap exchanges the elements at the indices i and j.
// If either i or j is not a valid index, the function does nothing.
func (list *List) Swap(i, j int) {
	if list.isValidIndex(i) && list.isValidIndex(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

func (list *List) Insert(index int, values ...interface{}) {

	if !list.isValidIndex(index) {
		// Append
		if index == list.size {
			list.Add(values...)
		}
		return
	}
	n := len(values)
	list.size += n
	// Move elements starting from the index n positions backward
	// Insert the new values
	copy(list.elements[index+n:], list.elements[index:list.size-n])
	copy(list.elements[index:], values)
}

// Set replaces the element at the specified index with the provided value.
// If the index is invalid but equals the list size, it appends the value.
// Otherwise, it does nothing.
func (list *List) Set(index int, value interface{}) {

	if !list.isValidIndex(index) {
		// Append
		if index == list.size {
			list.Add(value)
		}
		return
	}
	list.elements[index] = value
}

// IndexOf returns index of provided element
func (list *List) IndexOf(value interface{}) int {
	if list.size == 0 {
		return -1
	}
	for index, element := range list.elements {
		if value == element {
			return index
		}
	}
	return -1
}

func (list *List) Sort(comparator utils.Comparator) {
	if len(list.elements) < 2 {
		return
	}
	utils.Sort(list.elements[:list.size], comparator)
}

// Empty returns true if list does not contain any elements.
func (list *List) Empty() bool {
	return list.size == 0
}

// Size returns size of list
func (list *List) Size() int {
	return list.size
}

// Clear removes all elements from the list.
func (list *List) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

// String returns a string representation of the list
// with elements separated by commas.
func (list *List) String() string {
	str := "[List]\n"
	strs := []string{}
	for _, element := range list.elements[:list.size] {
		strs = append(strs, fmt.Sprintf("%v", element))
	}
	str += strings.Join(strs, ", ")
	return str
}
