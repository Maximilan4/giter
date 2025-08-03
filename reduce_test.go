package giter

import (
	"maps"
	"reflect"
	"slices"
	"testing"
)

func TestReduce(t *testing.T) {
	sl := []int{1, 2, 3}
	sum := func(acc, v int) int {
		return acc + v
	}
	exp := 6
	got := Reduce(slices.Values(sl), 0, sum)
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("bad reduce - exp %v, got %v", exp, got)
	}

}

func TestReduce2(t *testing.T) {
	sl := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	sum2 := func(acc, k, v int) int {
		return acc + (k * v)
	}
	exp := 14
	got := Reduce2(maps.All(sl), 0, sum2)
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("bad reduce - exp %v, got %v", exp, got)
	}

}
