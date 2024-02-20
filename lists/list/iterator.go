package list

type Iterator[T comparable] struct {
	list  *List[T]
	index int
}

// Iterator represents an iterator for the List data structure.
func (list *List[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{list: list, index: -1}
}

// Next advances the iterator to the next element in the list and returns true if there is a next element.
// If there is no next element, it returns false.
func (iterator *Iterator[T]) Next() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	return iterator.list.isValidIndex(iterator.index)
}

// Value returns the value of the current element at the iterator's position.
func (iterator *Iterator[T]) Value() interface{} {
	return iterator.list.elements[iterator.index]
}

// Index returns the current index position of the iterator.
func (iterator *Iterator[T]) Index() int {
	return iterator.index
}

// Begin resets the iterator to the beginning of the list.
func (iterator *Iterator[T]) Begin() {
	iterator.index = -1
}

// First moves the iterator to the first element in the list and returns true if successful.
// If the list is empty, it returns false.
func (iterator *Iterator[T]) First() bool {
	iterator.Begin()
	return iterator.Next()
}

// Prev moves the iterator to the previous element in the list and returns true if successful.
// If the iterator is already at the beginning of the list, it remains unchanged and returns false.
func (iterator *Iterator[T]) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	return iterator.list.isValidIndex(iterator.index)
}

// End moves the iterator past the last element
func (iterator *Iterator[T]) End() {
	iterator.index = iterator.list.size
}

// Last returns true if the iterator is positioned at the last element in the list.
// It moves the iterator to the end of the list and then moves it to the previous element.
func (iterator *Iterator[T]) Last() bool {
	iterator.End()
	return iterator.Prev()
}
