package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// Last returns the last element (a pair of values) of a [sequence].
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func Last[K, V any](seq2 iter.Seq2[K, V]) (K, V, error) {
	if seq2 == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilInput)
	}
	empty := true
	var resK K
	var resV V
	for k, v := range seq2 {
		empty = false
		resK = k
		resV = v
	}
	if empty {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrEmptyInput)
	}
	return resK, resV, nil
}

// LastPred returns the last element (a pair of values) of a [sequence] that satisfies a specified predicate.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func LastPred[K, V any](seq2 iter.Seq2[K, V], pred func(K, V) bool) (K, V, error) {
	if seq2 == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilInput)
	}
	if pred == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilPredicate)
	}
	empty := true
	found := false
	var resK K
	var resV V
	for k, v := range seq2 {
		empty = false
		if pred(k, v) {
			found = true
			resK = k
			resV = v
		}
	}
	if empty {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrEmptyInput)
	}
	if !found {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNoMatch)
	}
	return resK, resV, nil
}

// LastOrDefault returns the last element (a pair of values) of a [sequence],
// or a pair of corresponding [zero values] if the sequence contains no elements.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
// [zero values]: https://go.dev/ref/spec#The_zero_value
func LastOrDefault[K, V any](seq2 iter.Seq2[K, V]) (K, V, error) {
	if seq2 == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilInput)
	}
	k, v, err := Last(seq2)
	if err != nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), nil
	}
	return k, v, nil
}

// LastOrDefaultPred returns the last element (a pair of values) of a [sequence]
// that satisfies a predicate or a pair of corresponding [zero values] if no such element is found.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
// [zero values]: https://go.dev/ref/spec#The_zero_value
func LastOrDefaultPred[K, V any](seq2 iter.Seq2[K, V], pred func(K, V) bool) (K, V, error) {
	if seq2 == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilInput)
	}
	if pred == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilPredicate)
	}
	k, v, err := LastPred(seq2, pred)
	if err != nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), nil
	}
	return k, v, nil
}
