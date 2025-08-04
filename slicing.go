package giter

import (
	"iter"
)

// Take - takes only n first elements from iter.Seq
func Take[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	if n < 1 {
		return seq
	}

	var i int

	return func(yield func(T) bool) {
		for item := range seq {
			if i == n {
				return
			}

			if !yield(item) {
				return
			}
			i++
		}
	}
}

// Take2 - takes only n first elements from iter.Seq2
func Take2[K, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	if n < 1 {
		return seq
	}

	var i int

	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if i == n {
				return
			}

			if !yield(k, v) {
				return
			}
			i++
		}
	}
}
