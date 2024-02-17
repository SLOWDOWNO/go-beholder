package linkedlistqueue_test

import (
	"go-beholder/queues/linkedlistqueue"
	"testing"
)

func TestQueueEnqueue(t *testing.T) {
	queue := linkedlistqueue.New()
	if actualValue := queue.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	if actualValue := queue.Values(); actualValue[0].(int) != 1 || actualValue[1].(int) != 2 || actualValue[2].(int) != 3 {
		t.Errorf("Got %v expected %v", actualValue, "[1,2,3]")
	}
	if actualValue := queue.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := queue.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := queue.Peek(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}

func TestQueuePeek(t *testing.T) {
	queue := linkedlistqueue.New()
	if actualValue, ok := queue.Peek(); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	if actualValue, ok := queue.Peek(); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}

func TestQueueDequeue(t *testing.T) {
	queue := linkedlistqueue.New()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Dequeue()
	if actualValue, ok := queue.Peek(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := queue.Dequeue(); actualValue != 2 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue, ok := queue.Dequeue(); actualValue != 3 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := queue.Dequeue(); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	if actualValue := queue.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := queue.Values(); len(actualValue) != 0 {
		t.Errorf("Got %v expected %v", actualValue, "[]")
	}
}
