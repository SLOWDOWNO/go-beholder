package hashset

import (
	"encoding/json"
	"go-beholder/containers"
)

// _ is an assertion to ensure that *Set implements the containers.JSONSerializer interface.
var _ containers.JSONSerializer = (*Set[int])(nil)

// _ is an assertion to ensure that *Set implements the containers.JSONDeSerializer interface.
var _ containers.JSONDeSerializer = (*Set[int])(nil)

// ToJSON serializes the set to JSON format.
func (set *Set[T]) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON deserializes the set from JSON format.
func (set *Set[T]) FromJSON(data []byte) error {
	var elements []T
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}

// MarshalJSON marshals the set to JSON format.
func (set *Set[T]) MarshalJSON() ([]byte, error) {
	return set.ToJSON()
}

// UnMarshalJSON unmarshals the set from JSON format.
func (set *Set[T]) UnMarshalJSON(bytes []byte) error {
	return set.FromJSON(bytes)
}
