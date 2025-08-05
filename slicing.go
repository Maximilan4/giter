package giter

import (
	"iter"
)

func Drop2[K, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	if n < 1 {
		return seq
	}

	var i int

	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if i < n {
				i++
				continue
			}

			if !yield(k, v) {
				return
			}
			i++
		}
	}
}

// Drop first n items from iter.Seq and continue iteration
func Drop[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	if n < 1 {
		return seq
	}

	var i int

	return func(yield func(T) bool) {
		for item := range seq {
			if i < n {
				i++
				continue
			}

			if !yield(item) {
				return
			}
			i++
		}
	}
}

// TakeWhile - take elements from iter.Seq, until given func returns true
func TakeWhile[WF func(T) bool, T any](seq iter.Seq[T], f WF) iter.Seq[T] {
	if f == nil {
		return seq
	}

	return func(yield func(T) bool) {
		for item := range seq {
			if !f(item) {
				return
			}

			if !yield(item) {
				return
			}
		}
	}
}

// TakeWhile2 - take elements from iter.Seq2, until given func returns true
func TakeWhile2[WF func(K, V) bool, K, V any](seq iter.Seq2[K, V], f WF) iter.Seq2[K, V] {
	if f == nil {
		return seq
	}

	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !f(k, v) {
				return
			}

			if !yield(k, v) {
				return
			}
		}
	}
}

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
