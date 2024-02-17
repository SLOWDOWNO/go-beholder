package linkedlistqueue

import (
	"fmt"
	"go-beholder/lists/linkedlist"
	"go-beholder/queues"
	"strings"
)

// _ is an assertion to ensure that *Queue implements the queues.Queue interface.
var _ queues.Queue = (*Queue)(nil)

// Queue represents a queue data structure implemented using a linked list.
type Queue struct {
	list *linkedlist.List // Underlying linked list to store queue elements.
}

// New creates and returns a new Queue instance.
func New() *Queue {
	return &Queue{list: &linkedlist.List{}}
}

// Enqueue adds a new element to the end of the queue.
func (queue *Queue) Enqueue(value interface{}) {
	queue.list.Add(value)
}

// Dequeue removes and returns the element at the front of the queue.
func (queue *Queue) Dequeue() (value interface{}, ok bool) {
	value, ok = queue.list.Get(0)
	if ok {
		queue.list.Remove(0)
	}
	return
}

// Peek returns the element at the front of the queue without removing it.
func (queue *Queue) Peek() (value interface{}, ok bool) {
	return queue.list.Get(0)
}

// Empty checks if the queue is empty.
func (queue *Queue) Empty() bool {
	return queue.list.Empty()
}

// Size returns the number of elements in the queue.
func (queue *Queue) Size() int {
	return queue.list.Size()
}

// Clear removes all elements from the queue.
func (queue *Queue) Clear() {
	queue.list.Clear()
}

// Values returns a slice containing all the elements in the queue.
func (queue *Queue) Values() []interface{} {
	return queue.list.Values()
}

// String returns a string representation of the queue.
func (queue *Queue) String() string {
	str := "[LinkedListQueue\n]"
	values := []string{}
	for _, value := range queue.list.Values() {
		values = append(values, fmt.Sprintf("%v ", value))
	}
	str += strings.Join(values, ",")
	return str
}
