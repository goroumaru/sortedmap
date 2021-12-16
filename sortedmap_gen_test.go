//go:build go1.18

package sortedmap

import (
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
