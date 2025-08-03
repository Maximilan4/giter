package giter

import (
	"iter"
)

// Map2 - apply a modifying func to each element of iter.Seq2. Return a new iter.Seq2 of return type (RV)
func Map2[MF ~func(K, V) (RK, RV), Seq iter.Seq2[K, V], K, V, RK, RV any](seq Seq, f MF) iter.Seq2[RK, RV] {
	return func(yield func(RK, RV) bool) {
		for k, v := range seq {
			if !yield(f(k, v)) {
				return
			}
		}
	}
}

// Map21 - apply a modifying func to each element of iter.Seq2. Return a new iter.Seq of return type (RV)
func Map21[MF ~func(K, V) RV, Seq iter.Seq2[K, V], K, V, RV any](seq Seq, f MF) iter.Seq[RV] {
	return func(yield func(RV) bool) {
		for k, v := range seq {
			if !yield(f(k, v)) {
				return
			}
		}
	}
}

// Map12 - apply a modifying func to each element of iter.Seq. Return a new iter.Seq2 of return type (RK, RV)
func Map12[MF ~func(V) (RK, RV), Seq iter.Seq[V], V, RK, RV any](seq Seq, f MF) iter.Seq2[RK, RV] {
	return func(yield func(RK, RV) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

// Map - apply a modifying func to each element. Return a new iter.Seq of return type (RV)
func Map[MF ~func(V) RV, Seq iter.Seq[V], V, RV any](seq Seq, f MF) iter.Seq[RV] {
	return func(yield func(RV) bool) {
		for v := range seq {
			rv := f(v)
			if !yield(rv) {
				return
			}
		}
	}
}
