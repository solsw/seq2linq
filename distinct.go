package seq2linq

import (
	"iter"
	"reflect"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// Distinct returns distinct elements (pairs of values) from a [sequence]
// by using [reflect.DeepEqual] to determine equality of corresponding values of pairs.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func Distinct[K, V any](seq2 iter.Seq2[K, V]) (iter.Seq2[K, V], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	r, err := DistinctEq(seq2, func(k1 K, v1 V, k2 K, v2 V) bool {
		return reflect.DeepEqual(k1, k2) && reflect.DeepEqual(v1, v2)
	})
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// DistinctEq returns distinct elements (pairs of values) from a [sequence]
// by using a specified function to determine equality of corresponding values of pairs.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func DistinctEq[K, V any](seq2 iter.Seq2[K, V], eq func(K, V, K, V) bool) (iter.Seq2[K, V], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if eq == nil {
		return nil, errorhelper.CallerError(ErrNilEqual)
	}
	type Key generichelper.Tuple2[K, V]
	r, err := DistinctByEq(seq2,
		func(k K, v V) Key { return Key{Item1: k, Item2: v} },
		func(key1, key2 Key) bool { return eq(key1.Item1, key1.Item2, key2.Item1, key2.Item2) },
	)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}
