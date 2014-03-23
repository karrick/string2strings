// Package string2strings provides concurrency safe implementation of
// a map of strings to slices of strings.
package string2strings

import (
	"encoding/json"
	"fmt"
	"sort"
	"sync"
)

type StringToStrings struct {
	db     map[string][]string
	lock   sync.RWMutex
	sorted bool
}

// NewStringToStrings returns an initialized instance that maintains
// value strings in insertion order.
func NewStringToStrings() *StringToStrings {
	return &StringToStrings{db: make(map[string][]string)}
}

// NewStringToStrings returns an initialized instance that maintains
// value strings in lexicographical order.
func NewStringToSortedStrings() *StringToStrings {
	return &StringToStrings{db: make(map[string][]string), sorted: true}
}

// MarshallJSON implements Marshaler interface for converting instance
// to JSON. This method is called by json.Marshal().
//
//     db := NewStringToStrings()
//     bytes, err := json.Marshal(db)
//
func (self *StringToStrings) MarshalJSON() ([]byte, error) {
	self.lock.RLock()
	defer self.lock.RUnlock()
	bytes, err := json.Marshal(self.db)
	return bytes, err
}

func (self *StringToStrings) String() string {
	self.lock.RLock()
	defer self.lock.RUnlock()
	return fmt.Sprintf("%v", self.db)
}

// Get returns the list of strings associated with the specified key
// string.
func (self *StringToStrings) Get(key string) ([]string, bool) {
	self.lock.RLock()
	defer self.lock.RUnlock()
	v, ok := self.db[key]
	return v, ok
}

// Append either appends, when unsorted, or inserts, when sorted, the
// value to the slice of strings associated with the specified key
// string.
func (self *StringToStrings) Append(key, value string) {
	self.lock.Lock()
	defer self.lock.Unlock()
	if self.sorted {
		self.db[key] = insertStringToSortedStrings(value, self.db[key])
	} else {
		self.db[key] = append(self.db[key], value)
	}
}

func insertStringToSortedStrings(item string, list []string) []string {
	index := sort.SearchStrings(list, item)
	if index == len(list) || list[index] != item {
		// Grow list by one element. We'll use item but it
		// could be the empty string because it will be
		// overwritten.
		list = append(list, item)

		// Shift elements down one slot.
		copy(list[index+1:], list[index:])

		// Insert item into proper position.
		list[index] = item
	}
	return list
}

// Keys returns a slice of strings representing the keys held in a
// StringToStrings instance. Note that the order of the keys returns is
// indeterminant because of Go's conscience decision to randomize map
// key values.
func (self *StringToStrings) Keys() (keys []string) {
	self.lock.RLock()
	defer self.lock.RUnlock()
	keys = make([]string, 0, len(self.db))
	for k, _ := range self.db {
		keys = append(keys, k)
	}
	return
}

// ScrubKey removes the specified key from the instance, also removing
// the slice of strings associated with that key.
func (self *StringToStrings) ScrubKey(key string) {
	self.lock.Lock()
	defer self.lock.Unlock()
	delete(self.db, key)
}

// ScrubValue removes the specified value from all slices of strings
// in the instance. This handles both sorted and unsorted
// instances. Whereas the removal of a value from a sorted instance
// uses a binary tree to quickly remove the item, the removal of a
// value from an unsorted instance requires walking each slice of
// strings for each key in the instance.
func (self *StringToStrings) ScrubValue(value string) {
	self.lock.Lock()
	defer self.lock.Unlock()

	for key, list := range self.db {
		if self.sorted {
			index := sort.SearchStrings(list, value)
			if index < len(list) && list[index] == value {
				list = append(list[:index], list[index+1:]...)
				if len(list) == 0 {
					delete(self.db, key)
				} else {
					self.db[key] = list
				}
			}
		} else {
			for index, val := range list {
				if val == value {
					list = append(list[:index], list[index+1:]...)
					if len(list) == 0 {
						delete(self.db, key)
					} else {
						self.db[key] = list
					}
					// NOTE: if values are unique in a list, then could break
				}
			}
		}
	}
}
