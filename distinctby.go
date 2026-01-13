package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// DistinctBy returns distinct elements (pairs of values) from a [sequence] according to
// a specified key selector function and using [reflect.DeepEqual] to compare keys.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func DistinctBy[K, V, Key any](seq2 iter.Seq2[K, V], keySel func(K, V) Key) (iter.Seq2[K, V], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if keySel == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	r, err := DistinctByEq(seq2, keySel, generichelper.DeepEqual[Key])
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// elInElelEq determines (using 'eq') whether 'ee' contains 'el'
func elInElelEq[T any](el T, ee []T, eq func(T, T) bool) bool {
	for _, e := range ee {
		if eq(e, el) {
			return true
		}
	}
	return false
}

// DistinctByEq returns distinct elements (pairs of values) from a [sequence] according to
// a specified key selector function and using a specified equality function to compare keys.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func DistinctByEq[K, V, Key any](seq2 iter.Seq2[K, V],
	keySel func(K, V) Key, keyEq func(Key, Key) bool) (iter.Seq2[K, V], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if keySel == nil {
		return nil, errorhelper.CallerError(ErrNilSelector)
	}
	if keyEq == nil {
		return nil, errorhelper.CallerError(ErrNilEqual)
	}
	return func(yield func(K, V) bool) {
			var seen []Key
			for k, v := range seq2 {
				key := keySel(k, v)
				if !elInElelEq(key, seen, keyEq) {
					seen = append(seen, key)
					if !yield(k, v) {
						return
					}
				}
			}
		},
		nil
}
