package circularqueue

import (
	"fmt"
	"strings"
)

type Queue[T comparable] struct {
	vals    []T
	start   int
	end     int
	isFull  bool
	maxSize int
	size    int
}

func New[T comparable](maxSize int) *Queue[T] {
	if maxSize < 1 {
		panic("Invalid maxSize!")
	}
	queue := &Queue[T]{maxSize: maxSize}
	queue.Clear()
	return queue
}

func (q *Queue[T]) Enqueue(value T) {
	if q.Full() {
		q.Dequeue()
	}
	q.vals[q.end] = value
	q.end += 1
	if q.end >= q.maxSize {
		q.end = 0
	}
	if q.end == q.start {
		q.isFull = true
	}

	q.size = q.calculSize()
}

func (q *Queue[T]) Dequeue() (value T, ok bool) {
	if q.Empty() {
		return value, false
	}

	value, ok = q.vals[q.start], true
	q.start += 1
	if q.start >= q.maxSize {
		q.start = 0
	}
	q.isFull = false
	q.size--

	return
}

func (q *Queue[T]) Peek() (value T, ok bool) {
	if q.Empty() {
		return value, false
	}
	return q.vals[q.start], true
}

func (q *Queue[T]) Clear() {
	q.vals = make([]T, q.maxSize)
	q.start = 0
	q.end = 0
	q.isFull = false
	q.size = 0
}

func (q *Queue[T]) Full() bool {
	return q.Size() == q.maxSize
}

func (q *Queue[T]) Empty() bool {
	return q.Size() == 0
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) calculSize() int {
	if q.end < q.start {
		return q.maxSize - q.start + q.end
	} else if q.end == q.start {
		if q.isFull {
			return q.maxSize
		}
		return 0
	}
	return q.end - q.start
}

func (q *Queue[T]) Values() []T {
	vals := make([]T, q.size)
	for i := 0; i < q.size; i++ {
		vals[i] = q.vals[(q.start+i)%q.maxSize]
	}
	return vals
}

func (q *Queue[T]) String() string {
	str := "[CircularQueue]\n"
	var vals []string
	for _, val := range q.Values() {
		vals = append(vals, fmt.Sprintf("%v", val))
	}
	str += strings.Join(vals, ", ")
	return str
}
