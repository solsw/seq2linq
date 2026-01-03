package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// Concat concatenates [sequences].
//
// [sequences]: https://pkg.go.dev/iter#Seq2
func Concat[K, V any](seq2s ...iter.Seq2[K, V]) (iter.Seq2[K, V], error) {
	if len(seq2s) == 0 {
		return nil, errorhelper.CallerError(ErrEmptyInput)
	}
	for _, seq2 := range seq2s {
		if seq2 == nil {
			return nil, errorhelper.CallerError(ErrNilInput)
		}
	}
	return func(yield func(K, V) bool) {
			for _, seq2 := range seq2s {
				for k, v := range seq2 {
					if !yield(k, v) {
						return
					}
				}
			}
		},
		nil
}
