package string2strings

import (
	"bytes"
	"encoding/json"
	"sort"
	"strings"
)

type SortedStrings struct {
	list []string
}

func NewSortedStrings() *SortedStrings {
	return new(SortedStrings)
}

func NewSortedStringsFromStrings(values []string) *SortedStrings {
	result := make([]string, len(values))
	copy(result, values)
	sort.Strings(result)
	return &SortedStrings{list: result}
}

func (self *SortedStrings) Store(value string) {
	index := sort.SearchStrings(self.list, value)
	if index == len(self.list) || self.list[index] != value {
		// Grow list by one element. We'll use string's zero
		// value for now because it will be overwritten.
		self.list = append(self.list, "")

		// Shift elements down one slot.
		copy(self.list[index+1:], self.list[index:])

		// Insert value into proper position.
		self.list[index] = value
	}
}

func (self *SortedStrings) Delete(value string) {
	index := sort.SearchStrings(self.list, value)
	if index < len(self.list) && self.list[index] == value {
		self.list = append(self.list[:index], self.list[index+1:]...)
	}
}

func (self *SortedStrings) Strings() []string {
	return self.list
}

func (self *SortedStrings) String() string {
	var blob bytes.Buffer
	blob.WriteRune('[')
	if len(self.list) > 0 {
		blob.WriteString(strings.Join(self.list, " "))
	}
	blob.WriteRune(']')
	return blob.String()
}

func (self *SortedStrings) MarshalJSON() ([]byte, error) {
	if len(self.list) > 0 {
		return json.Marshal(self.list)
	}
	return []byte{'[', ']'}, nil
}
