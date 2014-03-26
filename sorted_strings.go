package string2strings

import (
	"sort"
)

type SortedStrings []string

func NewSortedStrings(values []string) SortedStrings {
	list := make([]string, len(values))
	copy(list, values)
	return SortedStrings(list)
}

func (list SortedStrings) Insert(value string) SortedStrings {
	if list == nil {
		list = make(SortedStrings, 0)
	}
	index := sort.SearchStrings(list, value)
	if index == len(list) || list[index] != value {
		// Grow list by one element. We'll use string's zero
		// value for now because it will be overwritten.
		list = append(list, "")

		// Shift elements down one slot.
		copy(list[index+1:], list[index:])

		// Insert value into proper position.
		list[index] = value
	}
	return list
}

func (list SortedStrings) Delete(value string) SortedStrings {
	// if list == nil {
	// 	list = make(SortedStrings, 0)
	// }
	index := sort.SearchStrings(list, value)
	if index < len(list) && list[index] == value {
		list = append(list[:index], list[index+1:]...)
	}
	return list
}
