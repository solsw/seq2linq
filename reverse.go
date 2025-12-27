package seq2linq

import (
	"iter"
	"slices"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

// Reverse inverts the order of the elements in a [sequence].
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func Reverse[K, V any](seq2 iter.Seq2[K, V]) (iter.Seq2[K, V], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	tt := iterhelper.Collect2Tuple(seq2)
	slices.Reverse(tt)
	return iterhelper.Var2Tuple(tt...), nil
}
