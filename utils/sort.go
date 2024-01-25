package utils

import "sort"

func Sort(values []interface{}, comparator Comparator) {
	sort.Sort(sortable{values, comparator})
}

type sortable struct {
	values     []interface{}
	comparator Comparator
}

// Implement interface
func (s sortable) Len() int {
	return len(s.values)
}

func (s sortable) Less(i, j int) bool {
	return s.comparator(s.values[i], s.values[j]) < 0
}

func (s sortable) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}
