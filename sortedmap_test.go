package sortedmap

import (
	"encoding/json"
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
	s1.Sort()

	m2 := map[string]interface{}{
		"Z": s1,
		"A": nil,
	}

	s2 := AsSortedMap(m2)
	s2.Sort()

	b, err := json.MarshalIndent(s2, "", "  ")
	t.Log(string(b))
	if err != nil {
		t.Fail()
	}
}
