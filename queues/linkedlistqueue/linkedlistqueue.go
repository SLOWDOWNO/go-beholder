package linkedlistqueue

import (
	"fmt"
	"go-beholder/lists/linkedlist"
	"go-beholder/queues"
	"strings"
)

// _ is an assertion to ensure that *Queue[T] implements the queues.Queue interface.
var _ queues.Queue[int] = (*Queue[int])(nil)

// Queue represents a queue data structure implemented using a linked list.
type Queue[T comparable] struct {
	list *linkedlist.List[T] // Underlying linked list to store queue elements.
}

// New creates and returns a new Queue instance.
func New[T comparable]() *Queue[T] {
	return &Queue[T]{list: &linkedlist.List[T]{}}
}

// Enqueue adds a new element to the end of the queue.
func (queue *Queue[T]) Enqueue(value T) {
	queue.list.Add(value)
}

// Dequeue removes and returns the element at the front of the queue.
func (queue *Queue[T]) Dequeue() (value T, ok bool) {
	value, ok = queue.list.Get(0)
	if ok {
		queue.list.Remove(0)
	}
	return
}

// Peek returns the element at the front of the queue without removing it.
func (queue *Queue[T]) Peek() (value T, ok bool) {
	return queue.list.Get(0)
}

// Empty checks if the queue is empty.
func (queue *Queue[T]) Empty() bool {
	return queue.list.Empty()
}

// Size returns the number of elements in the queue.
func (queue *Queue[T]) Size() int {
	return queue.list.Size()
}

// Clear removes all elements from the queue.
func (queue *Queue[T]) Clear() {
	queue.list.Clear()
}

// Values returns a slice containing all the elements in the queue.
func (queue *Queue[T]) Values() []T {
	return queue.list.Values()
}

// String returns a string representation of the queue.
func (queue *Queue[T]) String() string {
	str := "[LinkedListQueue]\n"
	values := []string{}
	for _, value := range queue.list.Values() {
		values = append(values, fmt.Sprintf("%v ", value))
	}
	str += strings.Join(values, ",")
	return str
}
