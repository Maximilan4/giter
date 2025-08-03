package giter

import (
	"slices"
	"testing"
)

func TestEach(t *testing.T) {
	slice := []string{"test", "slice", "values"}
	var (
		called = 0
		expMap = map[string]int{
			"test":   1,
			"slice":  2,
			"values": 3,
		}
		f = func(s string) {
			called++
			if expMap[s] != called {
				t.Errorf("iteration order mismatch for s %s - exp %d, got %d", s, expMap[s], called)
			}
		}
	)

	Each(slices.Values(slice), f)
	if called != len(slice) {
		t.Errorf("iteration mismatch, exp %d, got %d", called, len(slice))
	}
}

func TestEach2(t *testing.T) {
	slice := []string{"test", "slice", "values"}
	var (
		called = 0
		expMap = map[string]int{
			"test":   0,
			"slice":  1,
			"values": 2,
		}
		f = func(i int, s string) {
			called++
			if expMap[s] != i {
				t.Errorf("iteration order mismatch for s %s - exp %d, got %d", s, expMap[s], i)
			}
		}
	)

	Each2(slices.All(slice), f)
	if called != len(slice) {
		t.Errorf("iteration mismatch, exp %d, got %d", called, len(slice))
	}
}
