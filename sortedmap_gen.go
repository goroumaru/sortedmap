//go:build go1.18

package sortedmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"

	"golang.org/x/exp/constraints"
)

//
// KeyValuePair describes an entry in SortedMap
//
type KeyValuePair[T any] struct {
	Key   string
	Value T
}

//
// String implements the Stringer interface for KeyValuePair
//
func (e KeyValuePair[T]) String() string {
	return fmt.Sprintf("%q: %v", e.Key, e.Value)
}

//
// SortedMap is a structure that can sort a map[string]type by key
//
type SortedMap[T any] []KeyValuePair[T]

func (s SortedMap[T]) Len() int           { return len(s) }
func (s SortedMap[T]) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortedMap[T]) Less(i, j int) bool { return s[i].Key < s[j].Key }

//
// Sort sorts a SortedMap (that should have probably be called SortableMap
//
func (s SortedMap[T]) Sort() { sort.Sort(s) }

//
// Add adds an entry to a SortedMap (this require re-sorting the SortedMap when ready to display).
// Note that a SortedMap is internally a slice so you need to do something like:
//
// 	s := NewSortedMap()
// 	s = s.Add(key1, value1)
// 	s = s.Add(key2, value2)
//
func (s SortedMap[T]) Add(key string, value T) SortedMap[T] {
	return append(s, KeyValuePair[T]{key, value})
}

//
// Keys returns the list of keys for the entries in this SortedMap
//
func (s SortedMap[T]) Keys() (keys []string) {
	for _, kv := range s {
		keys = append(keys, kv.Key)
	}

	return
}

//
// Values returns the list of values for the entries in this SortedMap
//
func (s SortedMap[T]) Values() (values []T) {
	for _, kv := range s {
		values = append(values, kv.Value)
	}

	return
}

//
// MarshalJSON implements the json.Marshaler interface
//
func (s SortedMap[T]) MarshalJSON() ([]byte, error) {
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
func AsSortedMap[T any](m map[string]T) (s SortedMap[T]) {
	s = make(SortedMap[T], 0, len(m))
	for k, v := range m {
		s = append(s, KeyValuePair[T]{k, v})
	}

	s.Sort()
	return
}

//
// NewSortedMapOf returns a SortedMap of string/T
// Use the Add method to add elements and the Sort method to sort.
//
func NewSortedMapOf[T any]() (s SortedMap[T]) {
	return make(SortedMap[T], 0)
}

//
// NewSortedMap returns a SortedMap.
// Use the Add method to add elements and the Sort method to sort.
//
func NewSortedMap() (s SortedMap[any]) {
	return make(SortedMap[any], 0)
}

//
// ValueKeyPair describes an entry in SortedByValue
//
type ValueKeyPair[T constraints.Ordered] struct {
	Key   string
	Value T
}

//
// SortedByValue is a structure that can sort a map[string]int by value
//
type SortedByValue[T constraints.Ordered] []ValueKeyPair[T]

func (s SortedByValue[T]) Len() int           { return len(s) }
func (s SortedByValue[T]) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortedByValue[T]) Less(i, j int) bool { return s[i].Value < s[j].Value }

//
// Sort sorts a SortedByValue in ascending or descending order
//
func (s SortedByValue[T]) Sort(asc bool) {
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
func AsSortedByValue[T constraints.Ordered](m map[string]T, asc bool) (s SortedByValue[T]) {
	for k, v := range m {
		s = append(s, ValueKeyPair[T]{k, v})
	}

	s.Sort(asc)
	return
}

//
// Keys returns the list of keys for the entries in this SortedByValue
//
func (s SortedByValue[T]) Keys() (keys []string) {
	for _, kv := range s {
		keys = append(keys, kv.Key)
	}

	return
}

//
// Values returns the list of values for the entries in this SortedByValue
//
func (s SortedByValue[T]) Values() (values []T) {
	for _, kv := range s {
		values = append(values, kv.Value)
	}

	return
}
