package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

// Take returns a specified number of contiguous pairs of values from the start of a [sequence].
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func Take[K, V any](seq2 iter.Seq2[K, V], count int) (iter.Seq2[K, V], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if count <= 0 {
		return iterhelper.Empty2[K, V](), nil
	}
	return func(yield func(K, V) bool) {
			i := 0
			for k, v := range seq2 {
				if !yield(k, v) {
					return
				}
				i++
				if i >= count {
					return
				}
			}
		},
		nil
}

// TakeLast returns a specified number of contiguous pairs of values from the end of a [sequence].
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func TakeLast[K, V any](seq2 iter.Seq2[K, V], count int) (iter.Seq2[K, V], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if count <= 0 {
		return iterhelper.Empty2[K, V](), nil
	}
	vv := iterhelper.Collect2(seq2)
	if count >= len(vv)/2 {
		return seq2, nil
	}
	r, err := iterhelper.Var2[K, V](vv[len(vv)-count*2:]...)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// TakeWhile returns pairs of values from a [sequence] as long as a specified condition is true.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func TakeWhile[K, V any](seq2 iter.Seq2[K, V], pred func(K, V) bool) (iter.Seq2[K, V], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if pred == nil {
		return nil, errorhelper.CallerError(ErrNilPredicate)
	}
	return func(yield func(K, V) bool) {
			for k, v := range seq2 {
				if !pred(k, v) {
					return
				}
				if !yield(k, v) {
					return
				}
			}
		},
		nil
}

// TakeWhileIdx returns pairs of values from a [sequence] as long as a specified condition is true.
// The pair's index is used in the logic of the predicate function.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func TakeWhileIdx[K, V any](seq2 iter.Seq2[K, V], pred func(K, V, int) bool) (iter.Seq2[K, V], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if pred == nil {
		return nil, errorhelper.CallerError(ErrNilPredicate)
	}
	return func(yield func(K, V) bool) {
			i := 0
			for k, v := range seq2 {
				if !pred(k, v, i) {
					return
				}
				if !yield(k, v) {
					return
				}
				i++
			}
		},
		nil
}
