package seq2linq

import (
	"iter"
	"reflect"

	"github.com/solsw/errorhelper"
)

// Contains determines whether a [sequence] contains a specified pair of values
// by using [reflect.DeepEqual] to determine equality of corresponding values.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func Contains[K, V any](seq2 iter.Seq2[K, V], valK K, valV V) (bool, error) {
	if seq2 == nil {
		return false, errorhelper.CallerError(ErrNilInput)
	}
	r, err := ContainsEq(seq2, valK, valV, func(k1 K, v1 V, k2 K, v2 V) bool {
		return reflect.DeepEqual(k1, k2) && reflect.DeepEqual(v1, v2)
	})
	if err != nil {
		return false, errorhelper.CallerError(err)
	}
	return r, nil
}

// ContainsEq determines whether a [sequence] contains a specified pair of values
// by using a specified function to determine equality of pairs of values.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func ContainsEq[K, V any](seq2 iter.Seq2[K, V], valK K, valV V,
	eq func(k K, v V, valK K, valV V) bool) (bool, error) {
	if seq2 == nil {
		return false, errorhelper.CallerError(ErrNilInput)
	}
	if eq == nil {
		return false, errorhelper.CallerError(ErrNilEqual)
	}
	for k, v := range seq2 {
		if eq(k, v, valK, valV) {
			return true, nil
		}
	}
	return false, nil
}
