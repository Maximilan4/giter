package giter

import (
	"iter"
)

// Each - calls a given func with current iter value, no return value
func Each[EF ~func(V), Seq iter.Seq[V], V any](seq Seq, f EF) {
	for v := range seq {
		f(v)
	}
}

// Each2 - like Each, but for kv iterators
func Each2[EF ~func(K, V), Seq iter.Seq2[K, V], K, V any](seq Seq, f EF) {
	for k, v := range seq {
		f(k, v)
	}
}
