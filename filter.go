package giter

import (
	"iter"
)

// Filter2 - filters iter.Seq2 by func - it must return bool value. False value on elements is a signal to filter value.
func Filter2[FF ~func(K, V) bool, Seq iter.Seq2[K, V], K, V any](seq2 Seq, f FF) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq2 {
			if !f(k, v) {
				continue
			}

			if !yield(k, v) {
				return
			}
		}
	}
}

// Filter - filters iter.Seq by func - it must return bool value. False value on elements is a signal to filter value.
func Filter[FF ~func(V) bool, Seq iter.Seq[V], V any](seq Seq, f FF) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if !f(v) {
				continue
			}

			if !yield(v) {
				return
			}
		}
	}
}
