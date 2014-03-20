package main

import (
	"encoding/json"
	"sort"
	"sync"
)

type StringToStrings struct {
	db     map[string][]string
	lock   sync.RWMutex
	sorted bool
}

func NewStringToStrings() *StringToStrings {
	return &StringToStrings{db: make(map[string][]string)}
}

func NewStringToSortedStrings() *StringToStrings {
	return &StringToStrings{db: make(map[string][]string), sorted: true}
}

func (self *StringToStrings) Get(key string) ([]string, bool) {
	self.lock.RLock()
	defer self.lock.RUnlock()
	v, ok := self.db[key]
	return v, ok
}

func (self *StringToStrings) Append(key, value string) {
	self.lock.Lock()
	defer self.lock.Unlock()
	if self.sorted {
		self.db[key] = addItemToSortedList(value, self.db[key])
	} else {
		self.db[key] = append(self.db[key], value)
	}
}

func addItemToSortedList(item string, list []string) []string {
	index := sort.SearchStrings(list, item)
	if index == len(list) || list[index] != item {
		newList := append(make([]string, 0, 1+len(list)), list[:index]...)
		newList = append(newList, item)
		newList = append(newList, list[index:]...)
		return newList
	} else {
		return list
	}
}

func (self *StringToStrings) Keys() (keys []string) {
	self.lock.RLock()
	defer self.lock.RUnlock()
	keys = make([]string, 0, len(self.db))
	for k, _ := range self.db {
		keys = append(keys, k)
	}
	return
}

func (self *StringToStrings) MarshalJSON() ([]byte, error) {
	self.lock.RLock()
	defer self.lock.RUnlock()
	bytes, err := json.Marshal(self.db)
	return bytes, err
}

func (self *StringToStrings) ScrubKey(datum string) {
	self.lock.Lock()
	defer self.lock.Unlock()
	delete(self.db, datum)
}

func (self *StringToStrings) ScrubValue(datum string) {
	self.lock.Lock()
	defer self.lock.Unlock()

	for key, list := range self.db {
		if self.sorted {
			index := sort.SearchStrings(list, datum)
			if index < len(list) && list[index] == datum {
				list = append(list[:index], list[index+1:]...)
				if len(list) == 0 {
					delete(self.db, key)
				} else {
					self.db[key] = list
				}
			}
		} else {
			for index, value := range list {
				if value == datum {
					list = append(list[:index], list[index+1:]...)
					if len(list) == 0 {
						delete(self.db, key)
					} else {
						self.db[key] = list
					}
					// NOTE: if values are unique in a list, then can break
				}
			}
		}
	}
}
