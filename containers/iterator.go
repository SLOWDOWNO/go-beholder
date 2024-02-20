package containers

// IteratorWithIndex is an interface that represents an iterator with index.
type IteratorWithIndex[T any] interface {
	// Next advances the iterator to the next element and returns true if there is a next element, false otherwise.
	Next() bool

	// Values returns the current element of the iterator.
	Values() T

	// Index returns the index of the current element.
	Index() int

	// Begin resets the iterator to the beginning.
	Begin()

	// First returns true if the iterator is at the first element, false otherwise.
	First() bool
}

// RevIteratorWithIndex is an interface that represents a reverse iterator with index.
type RevIteratorWithIndex[T any] interface {
	// Prev moves the iterator to the previous element and returns true if there is a previous element, false otherwise.
	Prev() bool

	// End moves the iterator to the end of the collection.
	End()

	// Last moves the iterator to the last element and returns true if there is a last element, false otherwise.
	Last() bool

	IteratorWithIndex[T]
}
