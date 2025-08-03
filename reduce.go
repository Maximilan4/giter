package giter

import (
	"iter"
)

// Reduce2 - like a Reduce - reduces kv iterators into single value
func Reduce2[AccFun ~func(Acc, K, V) Acc, Seq iter.Seq2[K, V], Acc, K, V any](seq Seq, start Acc, af AccFun) Acc {
	var (
		k   K
		v   V
		acc Acc = start
	)

	for k, v = range seq {
		acc = af(acc, k, v)
	}

	return acc
}

// Reduce - reduce iterator into single value
func Reduce[AccFun ~func(Acc, V) Acc, Seq iter.Seq[V], Acc, V any](seq Seq, start Acc, af AccFun) Acc {
	var (
		v   V
		acc Acc = start
	)

	for v = range seq {
		acc = af(acc, v)
	}

	return acc
}
