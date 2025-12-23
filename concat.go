package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// Concat concatenates two [sequences].
//
// [sequences]: https://pkg.go.dev/iter#Seq2
func Concat[K, V any](in1, in2 iter.Seq2[K, V]) (iter.Seq2[K, V], error) {
	if in1 == nil || in2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	return func(yield func(K, V) bool) {
			for k1, v1 := range in1 {
				if !yield(k1, v1) {
					return
				}
			}
			for k2, v2 := range in2 {
				if !yield(k2, v2) {
					return
				}
			}
		},
		nil
}
