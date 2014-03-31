package string2strings

import (
	"encoding/json"
	"testing"
)

func TestSortedStringsInsertAddsFirstItem(t *testing.T) {
	var list SortedStrings

	list = list.Insert("foo")

	if len(list) != 1 {
		t.Errorf("Expected: %#v; Actual: %#v\n", 1, len(list))
	}
	if list[0] != "foo" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "foo", list[0])
	}
}

func TestSortedStringsNewThenInsertAddsFirstItem(t *testing.T) {
	list := NewSortedStrings()

	list = list.Insert("foo")

	if len(list) != 1 {
		t.Errorf("Expected: %#v; Actual: %#v\n", 1, len(list))
	}
	if list[0] != "foo" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "foo", list[0])
	}
}

func TestSortedStringsUninitializedString(t *testing.T) {
	list := NewSortedStrings()
	actual := list.String()
	expected := "[]"
	if expected != actual {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestSortedStringsString(t *testing.T) {
	list := NewSortedStringsFromStrings([]string{"def", "abc"})
	actual := list.String()
	expected := "[abc,def]"
	if expected != actual {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestSortedStringsUninitializedJSON(t *testing.T) {
	list := NewSortedStrings()
	blob, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Expected: %#v; Actual: %#v\n", nil, err)
	}
	actual := string(blob)
	expected := `[]`
	if expected != actual {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}
}

func TestSortedStringsJSON(t *testing.T) {
	list := NewSortedStringsFromStrings([]string{"def", "abc"})

	blob, err := json.Marshal(list)
	if err != nil {
		t.Errorf("Expected: %#v; Actual: %#v\n", nil, err)
	}
	actual := string(blob)
	expected := `["abc","def"]`
	if expected != actual {
		t.Errorf("Expected: %#v; Actual: %#v\n", expected, actual)
	}

}

func TestSortedStringsInsertDeleteFromZeroValue(t *testing.T) {
	var list SortedStrings

	list = list.Delete("foo")

	if len(list) != 0 {
		t.Errorf("Expected: %#v; Actual: %#v\n", 0, len(list))
	}
}

func TestSortedStringsInsertInsertBeginning(t *testing.T) {
	list := NewSortedStringsFromStrings([]string{"b", "c"})

	list = list.Insert("a")

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
	list := NewSortedStringsFromStrings([]string{"a", "c"})

	list = list.Insert("b")

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
	list := NewSortedStringsFromStrings([]string{"a", "b"})

	list = list.Insert("c")

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
	list := NewSortedStringsFromStrings([]string{"a", "b", "c"})

	list = list.Insert("a")

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
	list := NewSortedStringsFromStrings([]string{"a", "b", "c"})

	list = list.Insert("b")

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
	list := NewSortedStringsFromStrings([]string{"a", "b", "c"})

	list = list.Insert("c")

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
