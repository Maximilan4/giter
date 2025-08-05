package giter

import (
	"maps"
	"reflect"
	"slices"
	"strconv"
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
	orderedIter := func(yield func(k string, v int) bool) {
		for i := 1; i <= 4; i++ {
			if !yield(strconv.Itoa(i), i) {
				return
			}
		}
		return
	}
	itr := Take2(orderedIter, 2)

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

func TestTakeWhile(t *testing.T) {
	sl := []int{2, 4, 6, 8, 9, 10, 11, 12}
	itr := TakeWhile(slices.Values(sl), func(i int) bool {
		return (i % 2) == 0
	})

	got := slices.Collect(itr)
	exp := []int{2, 4, 6, 8}
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("bad take - exp %v, got %v", exp, got)
	}
}

func TestTakeWhileFuncNil(t *testing.T) {
	sl := []int{2, 4, 6, 8, 9, 10, 11, 12}
	itr := TakeWhile(slices.Values(sl), nil)

	got := slices.Collect(itr)
	exp := sl
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("bad take - exp %v, got %v", exp, got)
	}
}

func TestTakeWhile2(t *testing.T) {
	orderedIter := func(yield func(k string, v int) bool) {
		for i := 1; i <= 4; i++ {
			if !yield(strconv.Itoa(i), i) {
				return
			}
		}
		return
	}

	itr := TakeWhile2(orderedIter, func(k string, v int) bool {
		return k == "1" || ((v%2) == 0 && k != "4")
	})

	got := maps.Collect(itr)
	exp := map[string]int{
		"1": 1,
		"2": 2,
	}

	if !reflect.DeepEqual(exp, got) {
		t.Errorf("bad take - exp %v, got %v", exp, got)
	}
}

func TestTakeWhile2FuncNil(t *testing.T) {
	m := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
	}
	itr := TakeWhile2(maps.All(m), nil)

	got := maps.Collect(itr)
	exp := m
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("bad take - exp %v, got %v", exp, got)
	}
}

func TestDrop(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	itr := Drop(slices.Values(sl), 5)

	got := slices.Collect(itr)
	exp := sl[5:]
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("bad take - exp %v, got %v", exp, got)
	}
}

func TestDropLessThanOne(t *testing.T) {
	for name, n := range map[string]int{
		"0":  0,
		"-1": -1,
	} {
		t.Run(name, func(t *testing.T) {
			sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			itr := Drop(slices.Values(sl), n)

			got := slices.Collect(itr)
			exp := sl
			if !reflect.DeepEqual(exp, got) {
				t.Errorf("bad take - exp %v, got %v", exp, got)
			}
		})
	}
}

func TestDrop2(t *testing.T) {
	orderedIter := func(yield func(k string, v int) bool) {
		for i := 1; i <= 4; i++ {
			if !yield(strconv.Itoa(i), i) {
				return
			}
		}
		return
	}
	itr := Drop2(orderedIter, 2)

	got := maps.Collect(itr)
	exp := map[string]int{
		"3": 3,
		"4": 4,
	}

	if !reflect.DeepEqual(exp, got) {
		t.Errorf("bad take - exp %v, got %v", exp, got)
	}
}

func TestDrop2LessThanOne(t *testing.T) {
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
			itr := Drop2(maps.All(m), n)

			got := maps.Collect(itr)
			exp := m
			if !reflect.DeepEqual(exp, got) {
				t.Errorf("bad take - exp %v, got %v", exp, got)
			}
		})
	}
}

func TestDropWhile(t *testing.T) {
	sl := []int{2, 4, 6, 8, 9, 10, 11, 12}
	itr := DropWhile(slices.Values(sl), func(i int) bool {
		return (i % 2) == 0
	})

	got := slices.Collect(itr)
	exp := sl[4:]
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("bad take - exp %v, got %v", exp, got)
	}
}

func TestDropWhileFuncNil(t *testing.T) {
	sl := []int{2, 4, 6, 8, 9, 10, 11, 12}
	itr := DropWhile(slices.Values(sl), nil)

	got := slices.Collect(itr)
	exp := sl
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("bad take - exp %v, got %v", exp, got)
	}
}

func TestDropWhile2(t *testing.T) {
	orderedIter := func(yield func(k string, v int) bool) {
		for i := 1; i <= 4; i++ {
			if !yield(strconv.Itoa(i), i) {
				return
			}
		}
		return
	}

	itr := DropWhile2(orderedIter, func(k string, v int) bool {
		return k != "3"
	})

	got := maps.Collect(itr)
	exp := map[string]int{
		"3": 3,
		"4": 4,
	}

	if !reflect.DeepEqual(exp, got) {
		t.Errorf("bad take - exp %v, got %v", exp, got)
	}
}

func TestDropWhile2FuncNil(t *testing.T) {
	m := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
	}
	itr := DropWhile2(maps.All(m), nil)

	got := maps.Collect(itr)
	exp := m
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("bad take - exp %v, got %v", exp, got)
	}
}
