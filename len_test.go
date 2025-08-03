package giter

import (
	"slices"
	"testing"
)

func TestLength(t *testing.T) {
	slice := []string{"test", "slice", "len"}
	if int64(len(slice)) != Length(slices.Values(slice)) {
		t.Errorf("slice and seq lens are not equals")
	}
}

func TestLength2(t *testing.T) {
	slice := []string{"test", "slice", "len"}
	if int64(len(slice)) != Length2(slices.All(slice)) {
		t.Errorf("slice and seq lens are not equals")
	}
}
