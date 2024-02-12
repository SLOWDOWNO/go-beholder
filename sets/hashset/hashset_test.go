package hashset_test

import (
	"go-beholder/sets/hashset"
	"testing"
)

func TestNew(t *testing.T) {
	set := hashset.New(1, 2)

	if value := set.Size(); value != 2 {
		t.Errorf("Got %v but expected %v", value, 2)
	}
	if value := set.Contains(1); value != true {
		t.Errorf("Got %v but expected %v", value, true)
	}
	if value := set.Contains(2); value != true {
		t.Errorf("Got %v but expected %v", value, true)
	}
	if value := set.Contains(3); value != false {
		t.Errorf("Got %v but expected %v", value, false)
	}
}

func TestSetAdd(t *testing.T) {
	set := hashset.New()
	set.Add()
	set.Add(1)
	set.Add(2)
	set.Add(2, 3)
	set.Add()
	if actualValue := set.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := set.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
}

func TestSetRemove(t *testing.T) {
	set := hashset.New()
	set.Add(3, 1, 2)
	set.Remove()
	if actualValue := set.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	set.Remove(1)
	if actualValue := set.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	set.Remove(3)
	set.Remove(3)
	set.Remove()
	set.Remove(2)
	if actualValue := set.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestSetContains(t *testing.T) {
	set := hashset.New()
	set.Add(3, 1, 2)
	set.Add(2, 3)
	set.Add()
	if actualValue := set.Contains(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(1); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(1, 2, 3); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := set.Contains(1, 2, 3, 4); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}
