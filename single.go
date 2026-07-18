package seq2linq

import (
	"errors"
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// Single returns the only element (a pair of values) of a [sequence]
// and returns an error if there is not exactly one element (a pair of values) in the [sequence].
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func Single[K, V any](seq2 iter.Seq2[K, V]) (K, V, error) {
	if seq2 == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilInput)
	}
	next, stop := iter.Pull2(seq2)
	defer stop()
	k, v, ok := next()
	if !ok {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrEmptyInput)
	}
	_, _, ok = next()
	if ok {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrMultiElements)
	}
	return k, v, nil
}

// SinglePred returns the only element (a pair of values) of a [sequence] that satisfies a specified predicate.
// It returns an error if no such element exists or if more than one such element exists.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func SinglePred[K, V any](seq2 iter.Seq2[K, V], pred func(K, V) bool) (K, V, error) {
	if seq2 == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilInput)
	}
	if pred == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilPredicate)
	}
	empty := true
	found := false
	var rK K
	var rV V
	for k, v := range seq2 {
		empty = false
		if pred(k, v) {
			if found {
				return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrMultiMatch)
			}
			found = true
			rK = k
			rV = v
		}
	}
	if empty {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrEmptyInput)
	}
	if !found {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNoMatch)
	}
	return rK, rV, nil
}

// SingleOrDefault returns the only element (a pair of values) of a [sequence]
// or a specified pair of default values if the [sequence] is empty.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func SingleOrDefault[K, V any](seq2 iter.Seq2[K, V], defK K, defV V) (K, V, error) {
	if seq2 == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilInput)
	}
	k, v, err := Single(seq2)
	if err != nil {
		if errors.Is(err, ErrMultiElements) {
			return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrMultiElements)
		}
		// sequence is empty - return default values
		return defK, defV, nil
	}
	return k, v, nil
}

// SingleOrZero returns the only element (a pair of values) of a [sequence]
// or a pair of corresponding [zero values] if the [sequence] is empty.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
// [zero values]: https://go.dev/ref/spec#The_zero_value
func SingleOrZero[K, V any](seq2 iter.Seq2[K, V]) (K, V, error) {
	if seq2 == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilInput)
	}
	return SingleOrDefault(seq2, generichelper.ZeroValue[K](), generichelper.ZeroValue[V]())
}

// SingleOrDefaultPred returns the only element (a pair of values) of a [sequence]
// that satisfies a specified predicate or a pair of corresponding [zero values]
// if no such element exists or if more than one such element exists.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
// [zero values]: https://go.dev/ref/spec#The_zero_value
func SingleOrDefaultPred[K, V any](seq2 iter.Seq2[K, V], pred func(K, V) bool) (K, V, error) {
	if seq2 == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilInput)
	}
	if pred == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilPredicate)
	}
	k, v, err := SinglePred(seq2, pred)
	if err != nil {
		if errors.Is(err, ErrMultiMatch) {
			return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrMultiMatch)
		}
		// sequence is empty or no match - return zero values
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), nil
	}
	return k, v, nil
}
