package string2strings

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSortedStringsEmptyList(t *testing.T) {
	list := NewSortedStrings().Strings()
	if len(list) != 0 {
		t.Errorf("Expected: %#v; Actual: %#v\n", 0, len(list))
	}
}

func TestSortedStringsNewThenInsertAddsFirstItem(t *testing.T) {
	ss := NewSortedStrings()

	ss.Store("foo")

	list := ss.Strings()
	if len(list) != 1 {
		t.Errorf("Expected: %#v; Actual: %#v\n", 1, len(list))
	}
	if list[0] != "foo" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "foo", list[0])
	}
}

func TestSortedStringsUninitializedString(t *testing.T) {
	ss := NewSortedStrings()
	actual := ss.String()
	expected := "[]"
	if expected != actual {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestSortedStringsStringReturnsSameAsFmtSprintf(t *testing.T) {
	list := []string{"abc", "def"}
	expected := fmt.Sprintf("%v", list)
	ss := NewSortedStringsFromStrings(list)

	actual := ss.String()

	if expected != actual {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestSortedStringsUninitializedJSON(t *testing.T) {
	list := []string{}
	blob, err := json.Marshal(list)
	expected := string(blob)

	ss := NewSortedStrings()
	blob, err = json.Marshal(ss)
	if err != nil {
		t.Errorf("Expected: %#v; Actual: %#v\n", nil, err)
	}
	actual := string(blob)
	if expected != actual {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestSortedStringsJSON(t *testing.T) {
	list := []string{"abc", "def"}
	blob, err := json.Marshal(list)
	expected := string(blob)

	ss := NewSortedStringsFromStrings(list)
	blob, err = json.Marshal(ss)
	if err != nil {
		t.Errorf("Expected: %#v; Actual: %#v\n", nil, err)
	}
	actual := string(blob)
	if expected != actual {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestSortedStringsInsertDeleteFromZeroValue(t *testing.T) {
	ss := NewSortedStrings()

	ss.Delete("foo")

	list := ss.Strings()
	if len(list) != 0 {
		t.Errorf("Expected: %#v; Actual: %#v\n", 0, len(list))
	}
}

func TestSortedStringsInsertInsertBeginning(t *testing.T) {
	ss := NewSortedStringsFromStrings([]string{"b", "c"})

	ss.Store("a")

	list := ss.Strings()
	if len(list) != 3 {
		t.Errorf("Expected: %#v; Actual: %#v\n", 3, len(list))
	}
	if list[0] != "a" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "a", list[0])
	}
	if list[1] != "b" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "b", list[1])
	}
	if list[2] != "c" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "c", list[2])
	}
}

func TestSortedStringsInsertInsertMiddle(t *testing.T) {
	ss := NewSortedStringsFromStrings([]string{"a", "c"})

	ss.Store("b")

	list := ss.Strings()
	if len(list) != 3 {
		t.Errorf("Expected: %#v; Actual: %#v\n", 3, len(list))
	}
	if list[0] != "a" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "a", list[0])
	}
	if list[1] != "b" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "b", list[1])
	}
	if list[2] != "c" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "c", list[2])
	}
}

func TestSortedStringsInsertInsertEnd(t *testing.T) {
	ss := NewSortedStringsFromStrings([]string{"a", "b"})

	ss.Store("c")

	list := ss.Strings()
	if len(list) != 3 {
		t.Errorf("Expected: %#v; Actual: %#v\n", 3, len(list))
	}
	if list[0] != "a" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "a", list[0])
	}
	if list[1] != "b" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "b", list[1])
	}
	if list[2] != "c" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "c", list[2])
	}
}

func TestSortedStringsInsertNoRepeatBeginning(t *testing.T) {
	ss := NewSortedStringsFromStrings([]string{"a", "b", "c"})

	ss.Store("a")

	list := ss.Strings()
	if len(list) != 3 {
		t.Errorf("Expected: %#v; Actual: %#v\n", 3, len(list))
	}
	if list[0] != "a" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "a", list[0])
	}
	if list[1] != "b" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "b", list[1])
	}
	if list[2] != "c" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "c", list[2])
	}
}

func TestSortedStringsInsertNoRepeatMiddle(t *testing.T) {
	ss := NewSortedStringsFromStrings([]string{"a", "b", "c"})

	ss.Store("b")

	list := ss.Strings()
	if len(list) != 3 {
		t.Errorf("Expected: %#v; Actual: %#v\n", 3, len(list))
	}
	if list[0] != "a" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "a", list[0])
	}
	if list[1] != "b" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "b", list[1])
	}
	if list[2] != "c" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "c", list[2])
	}
}

func TestSortedStringsInsertNoRepeatEnd(t *testing.T) {
	ss := NewSortedStringsFromStrings([]string{"a", "b", "c"})

	ss.Store("c")

	list := ss.Strings()
	if len(list) != 3 {
		t.Errorf("Expected: %#v; Actual: %#v\n", 3, len(list))
	}
	if list[0] != "a" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "a", list[0])
	}
	if list[1] != "b" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "b", list[1])
	}
	if list[2] != "c" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "c", list[2])
	}
}
