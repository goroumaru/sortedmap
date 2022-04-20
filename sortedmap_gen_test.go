//go:build go1.18

package sortedmap

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAsSortedMapOfInts(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	s := AsSortedMap(m)
	keys := s.Keys()
	expected := []string{"a", "b", "c", "d", "e"}

	if !reflect.DeepEqual(keys, expected) {
		t.Log("expected", expected, "got", keys)
		t.Fail()
	}
}

func TestAsSortedMapIntKeys(t *testing.T) {
	m := map[int]string{
		5: "a",
		4: "b",
		3: "c",
		2: "d",
		1: "e",
	}

	s := AsSortedMap(m)
	keys := s.Keys()
	expected := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(keys, expected) {
		t.Log("expected", expected, "got", keys)
		t.Fail()
	}
}

func TestAsSortedByIntValue(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	s := AsSortedByValue(m, true)
	keys := s.Keys()
	expected := []string{"a", "b", "c", "d", "e"}

	if !reflect.DeepEqual(keys, expected) {
		t.Log("expected", expected, "got", keys)
		t.Fail()
	}
}

func TestAsSortedByStringValue(t *testing.T) {
	m := map[string]string{
		"a": "e",
		"b": "d",
		"c": "c",
		"d": "b",
		"e": "a",
	}

	s := AsSortedByValue(m, true)
	values := s.Values()
	expected := []string{"a", "b", "c", "d", "e"}

	if !reflect.DeepEqual(values, expected) {
		t.Log("expected", expected, "got", values)
		t.Fail()
	}
}

func ExampleAsSortedMapIntKey() {
	unsorted := map[int]any{
		2: 2.0,
		1: 1,
		3: true,
		5: nil,
		4: "four",
	}

	fmt.Println(AsSortedMap(unsorted))
	// Output:
	// [1: 1 2: 2 3: true 4: four 5: <nil>]
}
