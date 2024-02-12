package containers

// JSONSerializer defines the interface for serializing an object to JSON format.
type JSONSerializer interface {
	ToJSON() ([]byte, error)
	MarshalJSON() ([]byte, error)
}

// JSONDeSerializer defines the interface for deserializing an object from JSON format.
type JSONDeSerializer interface {
	FromJSON([]byte) error
	UnMarshalJSON([]byte) error
}
