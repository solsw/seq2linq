package seq2linq

import (
	"iter"
	"sort"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func orderByLsPrim[K, V any](seq2 iter.Seq2[K, V], less func(K, V, K, V) bool) iter.Seq2[K, V] {
	tt := iterhelper.Collect2Tuple(seq2)
	sort.SliceStable(tt, func(i, j int) bool {
		return less(tt[i].Item1, tt[i].Item2, tt[j].Item1, tt[j].Item2)
	})
	return iterhelper.Var2Tuple(tt...)
}

// OrderByLs sorts the elements (pairs of values) of a [sequence]
// in ascending order using a specified less function.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func OrderByLs[K, V any](seq2 iter.Seq2[K, V],
	less func(k1 K, v1 V, k2 K, v2 V) bool) (iter.Seq2[K, V], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if less == nil {
		return nil, errorhelper.CallerError(ErrNilLess)
	}
	return orderByLsPrim(seq2, less), nil
}
