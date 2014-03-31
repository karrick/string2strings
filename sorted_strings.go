package string2strings

import (
	"bytes"
	"sort"
	"strings"
)

type SortedStrings []string

func NewSortedStrings() SortedStrings {
	var sorted SortedStrings
	return sorted
}

func NewSortedStringsFromStrings(values []string) SortedStrings {
	var sorted SortedStrings
	for _, value := range values {
		sorted = sorted.Insert(value)
	}
	return sorted
}

func (list SortedStrings) Insert(value string) SortedStrings {
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
	index := sort.SearchStrings(list, value)
	if index < len(list) && list[index] == value {
		list = append(list[:index], list[index+1:]...)
	}
	return list
}

func (list SortedStrings) String() string {
	var blob bytes.Buffer
	blob.WriteRune('[')
	blob.WriteString(strings.Join(list, ","))
	blob.WriteRune(']')
	return blob.String()
}
