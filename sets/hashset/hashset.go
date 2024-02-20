package hashset

import (
	"fmt"
	"go-beholder/sets"
	"strings"
)

// _ is an assertion to ensure that *Set implements the sets.Set interface.
var _ sets.Set[int] = (*Set[int])(nil)

// Set represents a set data structure.
type Set[T comparable] struct {
	items map[T]struct{}
}

// placeholder is a singleton struct used as a placeholder value.
var placeholder = struct{}{}

// New creates and returns a new set instance initialized with the provided values.
func New[T comparable](values ...T) *Set[T] {
	set := &Set[T]{items: make(map[T]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the specified items to the set.
func (set *Set[T]) Add(items ...T) {
	for _, item := range items {
		set.items[item] = placeholder
	}
}

// Remove removes specified items from set.
func (set *Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains checks whether the set contains all the specified items.
func (set *Set[T]) Contains(items ...T) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty checks if the set is empty.
func (set *Set[T]) Empty() bool {
	return set.Size() == 0
}

// Size returns the number of items in the set.
func (set *Set[T]) Size() int {
	return len(set.items)
}

// Clear removes all items from the set, making it empty.
func (set *Set[T]) Clear() {
	set.items = make(map[T]struct{})
}

// Values returns a slice containing all the items in the set.
func (set *Set[T]) Values() []T {
	values := make([]T, set.Size())
	cnt := 0
	for item := range set.items {
		values[cnt] = item
		cnt++
	}
	return values
}

// String returns a string representation of the set.
func (set *Set[T]) String() string {
	str := "[HashSet]\n"
	items := []string{}
	for key := range set.items {
		items = append(items, fmt.Sprintf("%v", key))
	}
	str += strings.Join(items, ", ")
	return str
}
