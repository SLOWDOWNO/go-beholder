package utils

import "time"

type Comparator[T any] func(a, b T) int

func TimeComparator(a, b time.Time) int {
	switch {
	case a.After(b):
		return 1
	case a.Before(b):
		return -1
	default:
		return 0
	}
}
