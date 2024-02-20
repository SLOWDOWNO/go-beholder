package queues

import "go-beholder/containers"

type Queue[T comparable] interface {
	Enqueue(value T)
	Dequeue() (value T, ok bool)
	Peek() (value T, ok bool)

	containers.Container[T]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() T
	// String() string
}
