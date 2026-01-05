package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// DistinctBy returns distinct elements from a [sequence] according to
// a specified key selector function and using [generichelper.DeepEqual] to compare keys.
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

// DistinctByEq returns distinct elements from a sequence according to
// a specified key selector function and using a specified 'keyEqual' to compare keys.
//
// [DistinctByEq]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.distinctby
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

// // [DistinctByCmp] returns distinct elements from a sequence according to a specified key selector function
// // and using a specified 'compare' to compare keys. (See [DistinctCmp].)
// //
// // [DistinctByCmp]: https://learn.microsoft.com/dotnet/api/system.linq.enumerable.distinctby
// func DistinctByCmp[K, V, Key any](source iter.Seq[K, V],
// 	keySelector func(K, V) Key, compare func(Key, Key) int) (iter.Seq[K, V], error) {
// 	if source == nil {
// 		return nil, errorhelper.CallerError(ErrNilSource)
// 	}
// 	if keySelector == nil {
// 		return nil, errorhelper.CallerError(ErrNilSelector)
// 	}
// 	if compare == nil {
// 		return nil, errorhelper.CallerError(ErrNilCompare)
// 	}
// 	return func(yield func(K, V) bool) {
// 			seen := make([]Key, 0)
// 			for s := range source {
// 				k := keySelector(s)
// 				i := elIdxInElelCmp(k, seen, compare)
// 				if i == len(seen) || compare(k, seen[i]) != 0 {
// 					elIntoElelAtIdx(k, &seen, i)
// 					if !yield(s) {
// 						return
// 					}
// 				}
// 			}
// 		},
// 		nil
// }
