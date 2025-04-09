package giter

import (
	"iter"
)

const lookupStructTag = "giter"

type (
	Option[T any]         func(opts T)
	FilterFunc[K, V any]  func(K, V) bool
	MapFunc[K, V, RV any] func(K, V) RV

	ReduceFunc[Acc, K, V any] func(Acc, K, V) Acc
)

func Filter2[FF ~func(K, V) bool, Seq iter.Seq2[K, V], K, V any](seq2 Seq, f FF) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		var (
			k K
			v V
		)

		for k, v = range seq2 {
			if !f(k, v) {
				continue
			}

			if !yield(k, v) {
				return
			}
		}
	}
}

func Filter[FF ~func(V) bool, Seq iter.Seq[V], V any](seq Seq, f FF) iter.Seq[V] {
	return func(yield func(V) bool) {
		var (
			v V
		)

		for v = range seq {
			if !f(v) {
				continue
			}

			if !yield(v) {
				return
			}
		}
	}
}

func Map2[MF ~func(K, V) (RK, RV), Seq iter.Seq2[K, V], K, V, RK, RV any](seq Seq, f MF) iter.Seq2[RK, RV] {
	return func(yield func(RK, RV) bool) {
		var (
			k  K
			v  V
			rk RK
			rv RV
		)

		for k, v = range seq {
			rk, rv = f(k, v)
			if !yield(rk, rv) {
				return
			}
		}
	}
}

func Map21[MF ~func(K, V) RV, Seq iter.Seq2[K, V], K, V, RV any](seq Seq, f MF) iter.Seq[RV] {
	return func(yield func(RV) bool) {
		var (
			k  K
			v  V
			rv RV
		)

		for k, v = range seq {
			rv = f(k, v)
			if !yield(rv) {
				return
			}
		}
	}
}

func Map12[MF ~func(V) (RK, RV), Seq iter.Seq[V], V, RK, RV any](seq Seq, f MF) iter.Seq2[RK, RV] {
	return func(yield func(RK, RV) bool) {
		var (
			v  V
			rk RK
			rv RV
		)

		for v = range seq {
			rk, rv = f(v)
			if !yield(rk, rv) {
				return
			}
		}
	}
}

func Map[MF ~func(V) RV, Seq iter.Seq[V], V, RV any](seq Seq, f MF) iter.Seq[RV] {
	return func(yield func(RV) bool) {
		var (
			v  V
			rv RV
		)

		for v = range seq {
			rv = f(v)
			if !yield(rv) {
				return
			}
		}
	}
}

func Each[EF ~func(V), Seq iter.Seq[V], V any](seq Seq, f EF) {
	var v V
	for v = range seq {
		f(v)
	}
}

func Each2[EF ~func(K, V), Seq iter.Seq2[K, V], K, V any](seq Seq, f EF) {
	var (
		k K
		v V
	)
	for v = range seq {
		f(k, v)
	}
}

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

func Length[V any](seq iter.Seq[V]) (count int64) {
	for range seq {
		count++
	}
	return
}

func Length2[K, V any](seq iter.Seq2[K, V]) (count int64) {
	for range seq {
		count++
	}
	return
}
