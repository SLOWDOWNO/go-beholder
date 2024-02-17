package queues

import "go-beholder/containers"

type Queue interface {
	Enqueue(value interface{})
	Dequeue() (value interface{}, ok bool)
	Peek() (value interface{}, ok bool)

	containers.Container
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
	// String() string
}
