package giter

import (
	"maps"
	"reflect"
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
	sl := []int{1, 2, 3, 4}
	itr := Filter(slices.Values(sl), func(v int) bool {
		return v%2 == 0
	})

	exp := []int{2, 4}
	got := slices.Collect(itr)
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("bad filter - exp %v, got %v", exp, got)
	}
}

func TestFilter2(t *testing.T) {
	m := map[string]int{
		"skip":          1,
		"skip by value": 2,
		"no skip":       3,
	}
	itr := Filter2(maps.All(m), func(k string, v int) bool {
		return !(k == "skip" || v == 2)
	})

	exp := map[string]int{
		"no skip": 3,
	}
	got := maps.Collect(itr)
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("bad filter - exp %v, got %v", exp, got)
	}
}
