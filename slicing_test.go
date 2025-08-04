package giter

import (
	"maps"
	"reflect"
	"slices"
	"testing"
)

func TestTake(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	itr := Take(slices.Values(sl), 5)

	got := slices.Collect(itr)
	exp := sl[:5]
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("bad take - exp %v, got %v", exp, got)
	}
}

func TestTakeLessThanOne(t *testing.T) {
	for name, n := range map[string]int{
		"0":  0,
		"-1": -1,
	} {
		t.Run(name, func(t *testing.T) {
			sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			itr := Take(slices.Values(sl), n)

			got := slices.Collect(itr)
			exp := sl
			if !reflect.DeepEqual(exp, got) {
				t.Errorf("bad take - exp %v, got %v", exp, got)
			}
		})
	}
}

func TestTake2(t *testing.T) {
	m := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
	}
	itr := Take2(maps.All(m), 2)

	got := maps.Collect(itr)
	exp := map[string]int{
		"1": 1,
		"2": 2,
	}
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("bad take - exp %v, got %v", exp, got)
	}
}

func TestTake2LessThanOne(t *testing.T) {
	for name, n := range map[string]int{
		"0":  0,
		"-1": -1,
	} {
		t.Run(name, func(t *testing.T) {
			m := map[string]int{
				"1": 1,
				"2": 2,
				"3": 3,
				"4": 4,
			}
			itr := Take2(maps.All(m), n)

			got := maps.Collect(itr)
			exp := m
			if !reflect.DeepEqual(exp, got) {
				t.Errorf("bad take - exp %v, got %v", exp, got)
			}
		})
	}
}
