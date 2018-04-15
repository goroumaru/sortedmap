package sortedmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
)

//
// KeyValuePair describes an entry in SortedMap
//
type KeyValuePair struct {
	Key   string
	Value interface{}
}

//
// String implements the Stringer interface for KeyValuePair
//
func (e KeyValuePair) String() string {
	return fmt.Sprintf("%q: %v", e.Key, e.Value)
}

//
// SortedMap is a structure that can sort a map[string]type by key
//
type SortedMap []KeyValuePair

func (s SortedMap) Len() int           { return len(s) }
func (s SortedMap) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortedMap) Less(i, j int) bool { return s[i].Key < s[j].Key }

//
// Sort sorts a SortedMap (that should have probably be called SortableMap
//
func (s SortedMap) Sort() { sort.Sort(s) }

//
// Add adds an entry to a SortedMap (this require re-sorting the SortedMap when ready to display).
// Note that a SortedMap is internally a slice so you need to do something like:
//
// 	s := NewSortedMap()
// 	s = s.Add(key1, value1)
// 	s = s.Add(key2, value2)
//
func (s SortedMap) Add(key string, value interface{}) SortedMap {
	return append(s, KeyValuePair{key, value})
}

//
// Keys returns the list of keys for the entries in this SortdMap
//
func (s SortedMap) Keys() (keys []string) {
	for _, kv := range s {
		keys = append(keys, kv.Key)
	}

	return
}

//
// MarshalJSON implements the json.Marshaler interface
//
func (s SortedMap) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	var l = len(s)

	b.WriteString("{")

	for i, kv := range s {
		if bk, err := json.Marshal(kv.Key); err != nil {
			return nil, err
		} else {
			b.Write(bk)
		}

		b.WriteString(":")

		if bv, err := json.Marshal(kv.Value); err != nil {
			return nil, err
		} else {
			b.Write(bv)
		}

		if i < l-1 {
			b.WriteString(",")
		}
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

//
// AsSortedMap return a SortedMap from a map[string]type.
// Note that this will panic if the input object is not a map
//
func AsSortedMap(m interface{}) (s SortedMap) {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		panic("input object should be a map")
	}

	s = make(SortedMap, 0, val.Len())
	for _, k := range val.MapKeys() {
		v := val.MapIndex(k).Interface()
		s = append(s, KeyValuePair{k.String(), v})
	}

	s.Sort()
	return
}

//
// NewSortedMap returns a SortedMap.
// Use the Add method to add elements and the Sort method to sort.
//
func NewSortedMap() (s SortedMap) {
	return
}

//
// KeyIntValue describes an entry in SortedByValue
//
type KeyIntValue struct {
	Key   string
	Value int64
}

//
// SortedByValue is a structure that can sort a map[string]int by value
//
type SortedByValue []KeyIntValue

func (s SortedByValue) Len() int           { return len(s) }
func (s SortedByValue) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortedByValue) Less(i, j int) bool { return s[i].Value < s[j].Value }

//
// Sort sorts a SortedByValue in ascending or descending order
//
func (s SortedByValue) Sort(asc bool) {
	if asc {
		sort.Sort(s)
	} else {
		sort.Sort(sort.Reverse(s))
	}
}

//
// AsSortedByValue return a SortedByValue from a map[string]int
// Note that this will panic if the input object is not a map string/int
//
func AsSortedByValue(m interface{}, asc bool) (s SortedByValue) {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		panic("input object should be a map")
	}

	for _, k := range val.MapKeys() {
		v := val.MapIndex(k).Int()
		s = append(s, KeyIntValue{k.String(), v})
	}

	s.Sort(asc)
	return
}
