package hashset

import (
	"encoding/json"
	"go-beholder/containers"
)

// _ is an assertion to ensure that *Set implements the containers.JSONSerializer interface.
var _ containers.JSONSerializer = (*Set)(nil)

// _ is an assertion to ensure that *Set implements the containers.JSONDeSerializer interface.
var _ containers.JSONDeSerializer = (*Set)(nil)

// ToJSON serializes the set to JSON format.
func (set *Set) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON deserializes the set from JSON format.
func (set *Set) FromJSON(data []byte) error {
	elements := []interface{}{}
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}

// MarshalJSON marshals the set to JSON format.
func (set *Set) MarshalJSON() ([]byte, error) {
	return set.ToJSON()
}

// UnMarshalJSON unmarshals the set from JSON format.
func (set *Set) UnMarshalJSON(bytes []byte) error {
	return set.FromJSON(bytes)
}
