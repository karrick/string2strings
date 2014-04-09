package string2strings

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"
)

func TestStringToStringsStringUninitialized(t *testing.T) {
	db := NewStringToStrings()

	actual := fmt.Sprint(db)
	expected := "map[]"
	if expected != actual {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestStringToStringsString(t *testing.T) {
	sample := make(map[string][]string)
	sample["foo"] = []string{"bar", "flux"}
	expected := fmt.Sprintf("%v", sample)

	db := NewStringToStrings()
	db.Store("foo", "flux")
	db.Store("foo", "bar")

	actual := fmt.Sprint(db)
	if expected != actual {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestMarshalJSON(t *testing.T) {
	db := NewStringToStrings()
	db.Store("foo", "bar")

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

func TestGetEmptyDb(t *testing.T) {
	db := NewStringToStrings()

	list, ok := db.Get("")
	if ok != false {
		t.Errorf("Expected: %v; Actual: %v\n", false, ok)
	}
	if list != nil {
		t.Errorf("Expected: %v; Actual: %v\n", nil, list)
	}
}

func TestStoreOnMissingKey(t *testing.T) {
	db := NewStringToStrings()

	db.Store("key", "value")

	list, ok := db.Get("this key is not there")
	if ok != false {
		t.Errorf("Expected: %v; Actual: %v\n", false, ok)
	}
	if list != nil {
		t.Errorf("Expected: %v; Actual: %v\n", nil, list)
	}

	list, ok = db.Get("key")
	if ok != true {
		t.Errorf("Expected: %v; Actual: %v\n", true, ok)
	}
	actual := list.Strings()
	if len(actual) != 1 {
		t.Errorf("Expected: %v; Actual: %v\n", 1, len(actual))
	}
	if actual[0] != "value" {
		t.Errorf("Expected: %v; Actual: %v\n", "value", actual[0])
	}
}

func TestStoreOnExistingKey(t *testing.T) {
	db := NewStringToStrings()

	db.Store("key", "value1")
	db.Store("key", "value2")

	list, ok := db.Get("key")
	if ok != true {
		t.Errorf("Expected: %v; Actual: %v\n", true, ok)
	}
	actual := list.Strings()
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

func TestStoreKeepsStringsSorted(t *testing.T) {
	db := NewStringToStrings()

	db.Store("key", "value3")
	db.Store("key", "value1")
	db.Store("key", "value2")

	list, ok := db.Get("key")
	if ok != true {
		t.Errorf("Expected: %v; Actual: %v\n", true, ok)
	}
	actual := list.Strings()
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

func TestStoreOverwritesValue(t *testing.T) {
	db := NewStringToStrings()

	db.Store("key", "value3")
	db.Store("key", "value1")
	db.Store("key", "value2")

	ss := NewSortedStringsFromStrings([]string{"abc", "def"})
	db.StoreStrings("key", ss)

	list, ok := db.Get("key")
	if ok != true {
		t.Errorf("Expected: %v; Actual: %v\n", true, ok)
	}
	actual := list.Strings()
	if len(actual) != 2 {
		t.Errorf("Expected: %v; Actual: %v\n", 2, len(actual))
	}
	if actual[0] != "abc" {
		t.Errorf("Expected: %v; Actual: %v\n", "abc", actual[0])
	}
	if actual[1] != "def" {
		t.Errorf("Expected: %v; Actual: %v\n", "def", actual[1])
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

	db.Store("key1", "value1")
	actual := db.Keys()
	if len(actual) != 1 {
		t.Errorf("Expected: %v; Actual: %v\n", 1, len(actual))
	}
	if actual[0] != "key1" {
		t.Errorf("Expected: %v; Actual: %v\n", "key1", actual[0])
	}

	// single key with multiple values should also return only one key
	db.Store("key1", "value2")
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

	db.Store("key1", "value1")
	db.Store("key1", "value2")
	db.Store("key2", "value1")
	db.Store("key2", "value2")

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
	db.Store("foo", "bar")

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
	db := NewStringToStrings()
	db.Store("foo", "bar")

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

func TestScrubValueSingleFromSingle(t *testing.T) {
	db := NewStringToStrings()
	db.Store("foo", "bar")

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
	db := NewStringToStrings()
	db.Store("foo", "bar")
	db.Store("foo", "baz")

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
	db.Store("foo", "bar")
	db.Store("foo", "baz")
	db.Store("quz", "baz")

	db.ScrubValue("baz")

	bytes, err := json.Marshal(db)
	if err != nil {
		t.Errorf("Expected: %#v; Actual: %#v\n", nil, err)
	}

	actual := string(bytes)
	expected := `{"foo":["bar"]}`
	if expected != actual {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestScrubValueFromKey(t *testing.T) {
	db := NewStringToStrings()
	db.Store("foo", "bar")
	db.Store("baz", "bar")
	db.ScrubValueFromKey("bar", "foo")

	actual := db.String()
	expected := "map[baz:[bar]]"
	if expected != actual {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}
