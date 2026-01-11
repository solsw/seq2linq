package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// Zip applies a specified function to the corresponding pairs of values
// of two [sequence]s, producing a [sequence] of the resulting pairs of values.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func Zip[InK1, InV1, InK2, InV2, OutK, OutV any](
	first iter.Seq2[InK1, InV1], second iter.Seq2[InK2, InV2],
	sel func(InK1, InV1, InK2, InV2) (OutK, OutV)) (iter.Seq2[OutK, OutV], error) {
	if first == nil || second == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if sel == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	return func(yield func(OutK, OutV) bool) {
			next1, stop1 := iter.Pull2(first)
			defer stop1()
			next2, stop2 := iter.Pull2(second)
			defer stop2()
			for {
				k1, v1, ok1 := next1()
				k2, v2, ok2 := next2()
				if !ok1 || !ok2 {
					return
				}
				if !yield(sel(k1, v1, k2, v2)) {
					return
				}
			}
		},
		nil
}
