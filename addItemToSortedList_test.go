package string2strings

import (
	"testing"
)

func TestAddItemToListAddsFirstItem(t *testing.T) {
	var list []string

	list = insertStringToSortedStrings("foo", list)

	if len(list) != 1 {
		t.Errorf("Expected: %#v; Actual: %#v\n", 1, len(list))
	}
	if list[0] != "foo" {
		t.Errorf("Expected: %#v; Actual: %#v\n", "foo", list[0])
	}
}

func TestAddItemToListInsertBeginning(t *testing.T) {
	list := []string{"b", "c"}

	list = insertStringToSortedStrings("a", list)

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

func TestAddItemToListInsertMiddle(t *testing.T) {
	list := []string{"a", "c"}

	list = insertStringToSortedStrings("b", list)

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

func TestAddItemToListInsertEnd(t *testing.T) {
	list := []string{"a", "b"}

	list = insertStringToSortedStrings("c", list)

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

func TestAddItemToListNoRepeatBeginning(t *testing.T) {
	list := []string{"a", "b", "c"}

	list = insertStringToSortedStrings("a", list)

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

func TestAddItemToListNoRepeatMiddle(t *testing.T) {
	list := []string{"a", "b", "c"}

	list = insertStringToSortedStrings("b", list)

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

func TestAddItemToListNoRepeatEnd(t *testing.T) {
	list := []string{"a", "b", "c"}

	list = insertStringToSortedStrings("c", list)

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
