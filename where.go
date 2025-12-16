package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// Where uses a predicate to filter a sequence of pairs of values yielded by the [iterator].
//
// [iterator]: https://pkg.go.dev/iter#Seq2
func Where[K, V any](in iter.Seq2[K, V], pred func(K, V) bool) (iter.Seq2[K, V], error) {
	if in == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if pred == nil {
		return nil, errorhelper.CallerError(ErrNilPredicate)
	}
	return func(yield func(K, V) bool) {
			for k, v := range in {
				if !pred(k, v) {
					continue
				}
				if !yield(k, v) {
					return
				}
			}
		},
		nil
}

// WhereIdx uses a predicate to filter a sequence of pairs of values yielded by the [iterator].
// Each value's index is used in the logic of the predicate function.
//
// [iterator]: https://pkg.go.dev/iter#Seq2
func WhereIdx[K, V any](in iter.Seq2[K, V], pred func(K, V, int) bool) (iter.Seq2[K, V], error) {
	if in == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if pred == nil {
		return nil, errorhelper.CallerError(ErrNilPredicate)
	}
	return func(yield func(K, V) bool) {
			i := -1
			for k, v := range in {
				i++
				if !pred(k, v, i) {
					continue
				}
				if !yield(k, v) {
					return
				}
			}
		},
		nil
}
