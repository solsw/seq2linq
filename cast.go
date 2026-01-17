package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// Cast casts the elements (pairs of values) of a [sequence] to specified out types.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func Cast[InK, InV, OutK, OutV any](seq2 iter.Seq2[InK, InV]) (iter.Seq2[OutK, OutV], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	return func(yield func(OutK, OutV) bool) {
			for k, v := range seq2 {
				var anyK, anyV any = k, v
				if !yield(anyK.(OutK), anyV.(OutV)) {
					return
				}
			}
		},
		nil
}
