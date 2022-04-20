package sortedmap

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestAsSortedMap(t *testing.T) {
	m := map[string]interface{}{
		"a": 1,
		"b": 2.0,
		"c": true,
		"d": "four",
		"e": nil,
	}

	s := AsSortedMap(m)
	keys := s.Keys()
	expected := []string{"a", "b", "c", "d", "e"}

	if !reflect.DeepEqual(keys, expected) {
		t.Log("expected", expected, "got", keys)
		t.Fail()
	}
}

func TestSortdMap(t *testing.T) {
	s := NewSortedMap()

	s = s.Add("e", nil)
	s = s.Add("d", "four")
	s = s.Add("c", true)
	s = s.Add("b", 2.0)
	s = s.Add("a", 1)

	s.Sort()

	keys := s.Keys()
	expected := []string{"a", "b", "c", "d", "e"}

	if !reflect.DeepEqual(keys, expected) {
		t.Log("expected", expected, "got", keys)
		t.Fail()
	}
}

func TestMapOfMaps(t *testing.T) {
	m1 := map[string]interface{}{
		"a": 1,
		"b": 2.0,
		"c": true,
		"d": "four",
		"e": nil,
	}

	s1 := AsSortedMap(m1)

	m2 := map[string]interface{}{
		"Z": s1,
		"A": nil,
	}

	s2 := AsSortedMap(m2)

	b, err := json.MarshalIndent(s2, "", "  ")
	t.Log(string(b))
	if err != nil {
		t.Fail()
	}
}

func ExampleAsSortedMap() {
	unsorted := map[string]interface{}{
		"b": 2.0,
		"a": 1,
		"c": true,
		"e": nil,
		"d": "four",
	}

	fmt.Println(AsSortedMap(unsorted))
	// Output:
	// ["a": 1 "b": 2 "c": true "d": four "e": <nil>]
}
