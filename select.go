package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// Select projects each pair of values yielded by the [iterator] into a new form.
//
// [iterator]: https://pkg.go.dev/iter#Seq2
func Select[InK, InV, OutK, OutV any](in iter.Seq2[InK, InV], sel func(InK, InV) (OutK, OutV)) (iter.Seq2[OutK, OutV], error) {
	if in == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if sel == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	return func(yield func(OutK, OutV) bool) {
			for k, v := range in {
				if !yield(sel(k, v)) {
					return
				}
			}
		},
		nil
}

// SelectIdx projects each pair of values yielded by the [iterator] into a new form.
// Each pair's index is used in the logic of the selector function.
//
// [iterator]: https://pkg.go.dev/iter#Seq2
func SelectIdx[InK, InV, OutK, OutV any](in iter.Seq2[InK, InV],
	sel func(InK, InV, int) (OutK, OutV)) (iter.Seq2[OutK, OutV], error) {
	if in == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if sel == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	return func(yield func(OutK, OutV) bool) {
			i := 0
			for k, v := range in {
				if !yield(sel(k, v, i)) {
					return
				}
				i++
			}
		},
		nil
}
