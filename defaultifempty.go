package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// DefaultIfEmpty returns the elements of a specified [sequence] or a pair of values
// of the type parameters' [zero value]s in a singleton collection if the [sequence] is empty.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
// [zero value]: https://go.dev/ref/spec#The_zero_value
func DefaultIfEmpty[K, V any](seq2 iter.Seq2[K, V]) (iter.Seq2[K, V], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	r, err := DefaultIfEmptyDef(seq2, generichelper.ZeroValue[K](), generichelper.ZeroValue[V]())
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// DefaultIfEmptyDef returns the elements of a specified [sequence]
// or a pair of specified default values in a singleton collection if the [sequence] is empty.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func DefaultIfEmptyDef[K, V any](seq2 iter.Seq2[K, V], defK K, defV V) (iter.Seq2[K, V], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	return func(yield func(K, V) bool) {
			empty := true
			for k, v := range seq2 {
				empty = false
				if !yield(k, v) {
					return
				}
			}
			if empty {
				if !yield(defK, defV) {
					return
				}
			}
		},
		nil
}
