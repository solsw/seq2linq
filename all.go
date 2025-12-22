package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// All determines whether all elements of a sequence of pairs of values yielded by the [iterator] satisfy a predicate.
//
// [iterator]: https://pkg.go.dev/iter#Seq2
func All[K, V any](in iter.Seq2[K, V], pred func(K, V) bool) (bool, error) {
	if in == nil {
		return false, errorhelper.CallerError(ErrNilInput)
	}
	if pred == nil {
		return false, errorhelper.CallerError(ErrNilPredicate)
	}
	for k, v := range in {
		if !pred(k, v) {
			return false, nil
		}
	}
	return true, nil
}
