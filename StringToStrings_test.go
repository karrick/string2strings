package main

import (
	"encoding/json"
	"sort"
	"testing"
)

func TestGetEmptyDb(t *testing.T) {
	db := NewStringToStrings()

	actual, ok := db.Get("")
	if ok != false {
		t.Errorf("Expected: %v; Actual: %v\n", false, ok)
	}
	if len(actual) != 0 {
		t.Errorf("Expected: %v; Actual: %v\n", 0, len(actual))
	}
}

func TestAppendOnMissingKey(t *testing.T) {
	db := NewStringToStrings()

	db.Append("key", "value")

	actual, ok := db.Get("this key is not there")
	if ok != false {
		t.Errorf("Expected: %v; Actual: %v\n", false, ok)
	}
	if len(actual) != 0 {
		t.Errorf("Expected: %v; Actual: %v\n", 0, len(actual))
	}

	actual, ok = db.Get("key")
	if ok != true {
		t.Errorf("Expected: %v; Actual: %v\n", true, ok)
	}
	if len(actual) != 1 {
		t.Errorf("Expected: %v; Actual: %v\n", 1, len(actual))
	}
	if actual[0] != "value" {
		t.Errorf("Expected: %v; Actual: %v\n", "value", actual[0])
	}
}

func TestAppendOnExistingKey(t *testing.T) {
	db := NewStringToStrings()

	db.Append("key", "value1")
	db.Append("key", "value2")

	actual, ok := db.Get("key")
	if ok != true {
		t.Errorf("Expected: %v; Actual: %v\n", true, ok)
	}
	if len(actual) != 2 {
		t.Errorf("Expected: %v; Actual: %v\n", 2, len(actual))
	}
	if actual[0] != "value1" {
		t.Errorf("Expected: %v; Actual: %v\n", "value", actual[0])
	}
	if actual[1] != "value2" {
		t.Errorf("Expected: %v; Actual: %v\n", "value", actual[1])
	}
}

func TestAppendKeepsStringsSortedInsertionOrder(t *testing.T) {
	db := NewStringToStrings()

	db.Append("key", "value3")
	db.Append("key", "value1")
	db.Append("key", "value2")

	actual, ok := db.Get("key")
	if ok != true {
		t.Errorf("Expected: %v; Actual: %v\n", true, ok)
	}
	if len(actual) != 3 {
		t.Errorf("Expected: %v; Actual: %v\n", 3, len(actual))
	}
	if actual[0] != "value3" {
		t.Errorf("Expected: %v; Actual: %v\n", "value3", actual[0])
	}
	if actual[1] != "value1" {
		t.Errorf("Expected: %v; Actual: %v\n", "value1", actual[1])
	}
	if actual[2] != "value2" {
		t.Errorf("Expected: %v; Actual: %v\n", "value2", actual[2])
	}
}

func TestAppendKeepsStringsSorted(t *testing.T) {
	db := NewStringToSortedStrings()

	db.Append("key", "value3")
	db.Append("key", "value1")
	db.Append("key", "value2")

	actual, ok := db.Get("key")
	if ok != true {
		t.Errorf("Expected: %v; Actual: %v\n", true, ok)
	}
	if len(actual) != 3 {
		t.Errorf("Expected: %v; Actual: %v\n", 3, len(actual))
	}
	if actual[0] != "value1" {
		t.Errorf("Expected: %v; Actual: %v\n", "value1", actual[0])
	}
	if actual[1] != "value2" {
		t.Errorf("Expected: %v; Actual: %v\n", "value2", actual[1])
	}
	if actual[2] != "value3" {
		t.Errorf("Expected: %v; Actual: %v\n", "value3", actual[2])
	}
}

func TestKeysEmpty(t *testing.T) {
	db := NewStringToStrings()

	actual := db.Keys()
	if len(actual) != 0 {
		t.Errorf("Expected: %v; Actual: %v\n", 0, len(actual))
	}
}

func TestKeysSingleItem(t *testing.T) {
	db := NewStringToStrings()

	db.Append("key1", "value1")
	actual := db.Keys()
	if len(actual) != 1 {
		t.Errorf("Expected: %v; Actual: %v\n", 1, len(actual))
	}
	if actual[0] != "key1" {
		t.Errorf("Expected: %v; Actual: %v\n", "key1", actual[0])
	}

	// single key with multiple values should also return only one key
	db.Append("key1", "value2")
	actual = db.Keys()
	if len(actual) != 1 {
		t.Errorf("Expected: %v; Actual: %v\n", 1, len(actual))
	}
	if actual[0] != "key1" {
		t.Errorf("Expected: %v; Actual: %v\n", "key1", actual[0])
	}
}

func TestKeysMultipleItems(t *testing.T) {
	db := NewStringToStrings()

	db.Append("key1", "value1")
	db.Append("key1", "value2")
	db.Append("key2", "value1")
	db.Append("key2", "value2")

	actual := db.Keys()
	if len(actual) != 2 {
		t.Errorf("Expected: %v; Actual: %v\n", 2, len(actual))
	}
	sort.Strings(actual)
	if actual[0] != "key1" {
		t.Errorf("Expected: %v; Actual: %v\n", "key1", actual[0])
	}
	if actual[1] != "key2" {
		t.Errorf("Expected: %v; Actual: %v\n", "key2", actual[0])
	}
}

func TestMarshalJSON(t *testing.T) {
	db := NewStringToStrings()
	db.Append("foo", "bar")

	bytes, err := json.Marshal(db)
	if err != nil {
		t.Errorf("Expected: %#v; Actual: %#v\n", nil, err)
	}

	actual := string(bytes)
	expected := `{"foo":["bar"]}`
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestScrubKeyMissing(t *testing.T) {
	db := NewStringToStrings()

	db.ScrubKey("foo")

	bytes, err := json.Marshal(db)
	if err != nil {
		t.Errorf("Expected: %#v; Actual: %#v\n", nil, err)
	}

	actual := string(bytes)
	expected := `{}`
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestScrubKey(t *testing.T) {
	db := NewStringToStrings()
	db.Append("foo", "bar")

	db.ScrubKey("foo")

	bytes, err := json.Marshal(db)
	if err != nil {
		t.Errorf("Expected: %#v; Actual: %#v\n", nil, err)
	}

	actual := string(bytes)
	expected := `{}`
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestScrubValueMissing(t *testing.T) {
	db := NewStringToSortedStrings()
	db.Append("foo", "bar")

	db.ScrubValue("foo")

	bytes, err := json.Marshal(db)
	if err != nil {
		t.Errorf("Expected: %#v; Actual: %#v\n", nil, err)
	}

	actual := string(bytes)
	expected := `{"foo":["bar"]}`
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestScrubValueSingleFromSingle(t *testing.T) {
	db := NewStringToSortedStrings()
	db.Append("foo", "bar")

	db.ScrubValue("bar")

	bytes, err := json.Marshal(db)
	if err != nil {
		t.Errorf("Expected: %#v; Actual: %#v\n", nil, err)
	}

	actual := string(bytes)
	expected := `{}`
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestScrubValueSingleFromMultiple(t *testing.T) {
	db := NewStringToSortedStrings()
	db.Append("foo", "bar")
	db.Append("foo", "baz")

	db.ScrubValue("baz")

	bytes, err := json.Marshal(db)
	if err != nil {
		t.Errorf("Expected: %#v; Actual: %#v\n", nil, err)
	}

	actual := string(bytes)
	expected := `{"foo":["bar"]}`
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestScrubValueSingleFromMultipleUnsorted(t *testing.T) {
	db := NewStringToStrings()
	db.Append("foo", "bar")
	db.Append("foo", "baz")
	db.Append("quz", "baz")

	db.ScrubValue("baz")

	bytes, err := json.Marshal(db)
	if err != nil {
		t.Errorf("Expected: %#v; Actual: %#v\n", nil, err)
	}

	actual := string(bytes)
	expected := `{"foo":["bar"]}`
	if actual != expected {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}
