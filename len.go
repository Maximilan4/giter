package giter

import (
	"iter"
)

// Length - evals iter.Seq to int64 count of elems
func Length[V any](seq iter.Seq[V]) (count int64) {
	for range seq {
		count++
	}
	return
}

// Length2 - evals iter.Seq2 to int64 count of elems
func Length2[K, V any](seq iter.Seq2[K, V]) (count int64) {
	for range seq {
		count++
	}
	return
}
