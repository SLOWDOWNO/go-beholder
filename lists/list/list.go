package list

import (
	"fmt"
	"go-beholder/lists"
	"go-beholder/utils"
	"slices"
	"strings"
)

// _ is an assertion to ensure that *List implements the lists.List interface.
var _ lists.List[int] = (*List[int])(nil)

type List[T comparable] struct {
	elements []T
	size     int
}

// isValidIndex checks if the provided index is within the bounds of the list.
func (list *List[T]) isValidIndex(idx int) bool {
	return idx >= 0 && idx < list.size
}

// New creates a new List and adds the provided values, if any.
func New[T comparable](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Values returns a new slice that contains all elements in the list.
func (list *List[T]) Values() []T {
	newElements := make([]T, list.size)
	copy(newElements, list.elements[:list.size])
	return newElements
}

// Add appends the provided values to the end of the list.
func (list *List[T]) Add(values ...T) {
	for _, value := range values {
		list.elements = append(list.elements, value)
		list.size++
	}
}

// Get retrieves the element at the specified index from the list.
// It returns the element and true if the index is valid, and nil and false otherwise.
func (list *List[T]) Get(index int) (T, bool) {

	if list.isValidIndex(index) {
		var t T
		return t, false
	}
	return list.elements[index], true
}

// Remove deletes the element at the specified index from the list.
// If the index is not valid, the function does nothing.
func (list *List[T]) Remove(index int) {
	if !list.isValidIndex(index) {
		return
	}
	list.elements = append(list.elements[:index], list.elements[index+1:]...)
	list.size--
}

// Contains checks if the list contains all the provided values.
// It returns true if all values are found in the list, and false otherwise.
func (list *List[T]) Contains(values ...T) bool {
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
func (list *List[T]) Swap(i, j int) {
	if list.isValidIndex(i) && list.isValidIndex(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

func (list *List[T]) Insert(index int, values ...T) {

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
func (list *List[T]) Set(index int, value T) {

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
func (list *List[T]) IndexOf(value T) int {
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

func (list *List[T]) Sort(comparator utils.Comparator[T]) {
	if len(list.elements) < 2 {
		return
	}
	slices.SortFunc(list.elements[:list.size], comparator)
}

// Empty returns true if list does not contain any elements.
func (list *List[T]) Empty() bool {
	return list.size == 0
}

// Size returns size of list
func (list *List[T]) Size() int {
	return list.size
}

// Clear removes all elements from the list.
func (list *List[T]) Clear() {
	list.size = 0
	list.elements = []T{}
}

// String returns a string representation of the list
// with elements separated by commas.
func (list *List[T]) String() string {
	str := "[List]\n"
	strs := []string{}
	for _, element := range list.elements[:list.size] {
		strs = append(strs, fmt.Sprintf("%v", element))
	}
	str += strings.Join(strs, ", ")
	return str
}
