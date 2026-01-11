package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// SelectMany projects each pair of values of a [sequence] to another [sequence]
// and flattens the resulting sequences into one sequence.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func SelectMany[InK, InV, OutK, OutV any](seq2 iter.Seq2[InK, InV],
	sel func(InK, InV) iter.Seq2[OutK, OutV]) (iter.Seq2[OutK, OutV], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if sel == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	return func(yield func(OutK, OutV) bool) {
			for inK, inV := range seq2 {
				for outK, outV := range sel(inK, inV) {
					if !yield(outK, outV) {
						return
					}
				}
			}
		},
		nil
}

// SelectManyIdx projects each pair of values of a [sequence] and its index to another [sequence]
// and flattens the resulting sequences into one sequence.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func SelectManyIdx[InK, InV, OutK, OutV any](seq2 iter.Seq2[InK, InV],
	sel func(InK, InV, int) iter.Seq2[OutK, OutV]) (iter.Seq2[OutK, OutV], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if sel == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	return func(yield func(OutK, OutV) bool) {
			i := 0
			for inK, inV := range seq2 {
				for outK, outV := range sel(inK, inV, i) {
					if !yield(outK, outV) {
						return
					}
				}
				i++
			}
		},
		nil
}
