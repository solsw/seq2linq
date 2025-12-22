package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// Any determines whether a sequence of pairs of values yielded by the [iterator] contains any element.
//
// [iterator]: https://pkg.go.dev/iter#Seq2
func Any[K, V any](in iter.Seq2[K, V]) (bool, error) {
	if in == nil {
		return false, errorhelper.CallerError(ErrNilInput)
	}
	for range in {
		return true, nil
	}
	return false, nil
}

// AnyPred determines whether any pair of values yielded by the [iterator] satisfies a predicate.
//
// [iterator]: https://pkg.go.dev/iter#Seq2
func AnyPred[K, V any](in iter.Seq2[K, V], pred func(K, V) bool) (bool, error) {
	if in == nil {
		return false, errorhelper.CallerError(ErrNilInput)
	}
	if pred == nil {
		return false, errorhelper.CallerError(ErrNilPredicate)
	}
	for k, v := range in {
		if pred(k, v) {
			return true, nil
		}
	}
	return false, nil
}
