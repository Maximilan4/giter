package giter

import (
	"maps"
	"reflect"
	"slices"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	m := []int{1, 2, 3}
	seq := Map(slices.Values(m), func(v int) string {
		return strconv.Itoa(v)
	})

	exp := []string{"1", "2", "3"}
	got := slices.Collect(seq)
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("wrong seq elems, exp %v, got %v", exp, got)
	}
}

func TestMap12(t *testing.T) {
	m := []int{1, 2, 3}
	seq := Map12(slices.Values(m), func(v int) (string, int) {
		return strconv.Itoa(v), v
	})

	exp := map[string]int{"1": 1, "2": 2, "3": 3}
	got := maps.Collect(seq)
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("wrong seq elems, exp %v, got %v", exp, got)
	}
}

func TestMap2(t *testing.T) {
	m := map[string]int{
		"a": 97,
		"b": 98,
		"c": 99,
	}
	seq := Map2(maps.All(m), func(k string, v int) (string, rune) {
		return k, rune(v + 13)
	})

	exp := map[string]rune{
		"a": 'n',
		"b": 'o',
		"c": 'p',
	}
	got := maps.Collect(seq)
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("wrong seq elems, exp %v, got %v", exp, got)
	}
}

func TestMap21(t *testing.T) {
	m := map[string]int{
		"a": 97,
		"b": 98,
		"c": 99,
	}
	seq := Map21(maps.All(m), func(_ string, v int) rune {
		return rune(v + 13)
	})

	exp := []rune{'n', 'o', 'p'}
	got := slices.Collect(seq)
	slices.Sort(got)
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("wrong seq elems, exp %v, got %v", exp, got)
	}
}
