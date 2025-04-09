package giter

import (
	"iter"
	"maps"
	"slices"
)

// Iter - main container for chaining seq funcs
type Iter[V any] struct {
	seq iter.Seq[V]
}

// IterateBySlice - creates a sequence over slice
func IterateBySlice[Slice ~[]V, V any](s Slice) *Iter[V] {
	return Iterate(slices.Values(s))
}

// IterateByMapKeys - create a sequence over map keys
func IterateByMapKeys[Map ~map[K]V, K comparable, V any](s Map) *Iter[K] {
	return Iterate(maps.Keys(s))
}

// IterateByMapValues - create a sequence over map values
func IterateByMapValues[Map ~map[K]V, K comparable, V any](s Map) *Iter[V] {
	return Iterate(maps.Values(s))
}

// Iterate - wraps existing sequence
func Iterate[V any](seq iter.Seq[V]) *Iter[V] {
	return &Iter[V]{seq: seq}
}

// Filter - wrap sequence by a filter func
func (i *Iter[V]) Filter(f func(v V) bool) *Iter[V] {
	i.seq = Filter(i.seq, f)
	return i
}

// Map - wrap sequence by a map func
func (i *Iter[V]) Map(f func(v V) V) *Iter[V] {
	i.seq = Map(i.seq, f)
	return i
}

// Reduce - reduces seq by a given func to v
func (i *Iter[V]) Reduce(start V, f func(acc V, v V) V) V {
	return Reduce(i.seq, start, f)
}

// Pull - return a pull iterator from existing seq
func (i *Iter[V]) Pull() (next func() (V, bool), stop func()) {
	return iter.Pull(i.seq)
}

// Seq - returns a existing seq
func (i *Iter[V]) Seq() iter.Seq[V] {
	return i.seq
}

// Slice - collects a slice from seq
func (i *Iter[V]) Slice() []V {
	return slices.Collect(i.seq)
}

func (i *Iter[V]) Length() int64 {
	return Length(i.seq)
}
