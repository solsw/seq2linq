package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// Append appends a pair of values to the end of the [sequence].
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func Append[K, V any](seq2 iter.Seq2[K, V], k K, v V) (iter.Seq2[K, V], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	repeat1, _ := Repeat(k, v, 1)
	r, err := Concat(seq2, repeat1)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}
