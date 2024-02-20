package sets

import "go-beholder/containers"

// Set defines the interface for a set data structure.
type Set[T comparable] interface {
	Add(elements ...T)
	Remove(elements ...T)
	Contains(elements ...T) bool

	containers.Container[T]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []T
	// String() string
}
