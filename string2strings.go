// Package string2strings provides concurrency safe implementation of
// a map of strings to slices of strings.
package string2strings

import (
	"encoding/json"
	"fmt"
	"sync"
)

type StringToStrings struct {
	db   map[string]*SortedStrings
	lock sync.RWMutex
}

// NewStringToStrings returns an initialized instance that maintains
// value strings in lexicographical order.
func NewStringToStrings() *StringToStrings {
	return &StringToStrings{db: make(map[string]*SortedStrings)}
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
func (self *StringToStrings) Get(key string) (*SortedStrings, bool) {
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

	ss, ok := self.db[key]
	if !ok {
		ss = NewSortedStrings()
		self.db[key] = ss
	}
	ss.Store(value)
}

// Store stores the entire list of SortedStrings at the specified key.
func (self *StringToStrings) Store(key string, ss *SortedStrings) {
	self.lock.Lock()
	defer self.lock.Unlock()

	self.db[key] = ss
}

// Keys returns a slice of strings representing the keys held in a
// StringToStrings instance. Note that the order of the keys returns is
// indeterminant because of Go's conscience decision to randomize map
// key values.
func (self *StringToStrings) Keys() (keys []string) {
	self.lock.RLock()
	defer self.lock.RUnlock()

	keys = make([]string, 0, len(self.db))
	for k := range self.db {
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
	// NOTE: Alternate, less efficient implementation:
	// NOTE: Subroutine does the locking for this one.
	// for key := range self.db {
	// 	self.ScrubValueFromKey(value, key)
	// }

	self.lock.Lock()
	defer self.lock.Unlock()

	for key, list := range self.db {
		list.Delete(value)
		if len(list.Strings()) == 0 {
			delete(self.db, key)
		}
	}
}

func (self *StringToStrings) ScrubValueFromKey(value, key string) {
	self.lock.Lock()
	defer self.lock.Unlock()

	list := self.db[key]
	list.Delete(value)
	if len(list.Strings()) == 0 {
		delete(self.db, key)
	}
}
