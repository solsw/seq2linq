package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// OfType filters the elements (pairs of values) of a [sequence] based on specified out types.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func OfType[InK, InV, OutK, OutV any](seq2 iter.Seq2[InK, InV]) (iter.Seq2[OutK, OutV], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	return func(yield func(OutK, OutV) bool) {
			for inK, inV := range seq2 {
				var anyK any = inK
				outK, ok := anyK.(OutK)
				if !ok {
					continue
				}
				var anyV any = inV
				outV, ok := anyV.(OutV)
				if !ok {
					continue
				}
				if !yield(outK, outV) {
					return
				}
			}
		},
		nil
}
