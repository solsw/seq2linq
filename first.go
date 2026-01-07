package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// First returns the first element (a pair of values) of a [sequence].
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func First[K, V any](seq2 iter.Seq2[K, V]) (K, V, error) {
	if seq2 == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilInput)
	}
	for k, v := range seq2 {
		return k, v, nil
	}
	return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrEmptyInput)
}

// FirstPred returns the first element (a pair of values) in a [sequence] that satisfies a specified predicate.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func FirstPred[K, V any](seq2 iter.Seq2[K, V], pred func(K, V) bool) (K, V, error) {
	if seq2 == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilInput)
	}
	if pred == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilPredicate)
	}
	empty := true
	for k, v := range seq2 {
		empty = false
		if pred(k, v) {
			return k, v, nil
		}
	}
	if empty {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrEmptyInput)
	}
	return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNoMatch)
}

// FirstOrDefault returns the first element (a pair of values) of a [sequence],
// or a pair of corresponding [zero values] if the sequence contains no elements.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
// [zero values]: https://go.dev/ref/spec#The_zero_value
func FirstOrDefault[K, V any](seq2 iter.Seq2[K, V]) (K, V, error) {
	if seq2 == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilInput)
	}
	k, v, err := First(seq2)
	if err != nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), nil
	}
	return k, v, nil
}

// FirstOrDefaultPred returns the first element (a pair of values) of the [sequence]
// that satisfies a predicate or a pair of corresponding [zero values] if no such element is found.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
// [zero values]: https://go.dev/ref/spec#The_zero_value
func FirstOrDefaultPred[K, V any](seq2 iter.Seq2[K, V], pred func(K, V) bool) (K, V, error) {
	if seq2 == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilInput)
	}
	if pred == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilPredicate)
	}
	k, v, err := FirstPred(seq2, pred)
	if err != nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), nil
	}
	return k, v, nil
}
