package circularqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	q := New[int](3)
	assert.True(t, q.Empty())

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	vals := q.Values()
	assert.Equal(t, 1, vals[0])
	assert.Equal(t, 2, vals[1])
	assert.Equal(t, 3, vals[2])
	assert.False(t, q.Empty())
	assert.True(t, true, q.Full())
	assert.Equal(t, 3, q.Size())

	peek_val, ok := q.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1, peek_val)

	val, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, val)
	assert.Equal(t, 2, q.Size())

	vals = q.Values()
	assert.Equal(t, 2, vals[0])
	assert.Equal(t, 3, vals[1])

	q.Dequeue()
	q.Dequeue()
	assert.True(t, q.Empty())
}
