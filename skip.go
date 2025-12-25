package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

// Skip bypasses a specified number of pairs of values in a [sequence]
// and then returns the remaining ones.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func Skip[K, V any](seq2 iter.Seq2[K, V], count int) (iter.Seq2[K, V], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if count <= 0 {
		return seq2, nil
	}
	return func(yield func(K, V) bool) {
			i := 0
			for k, v := range seq2 {
				if i < count {
					i++
					continue
				}
				if !yield(k, v) {
					return
				}
			}
		},
		nil
}

// SkipLast returns a [sequence] that contains the pairs of values from 'seq2'
// with the last 'count' ones omitted.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func SkipLast[K, V any](seq2 iter.Seq2[K, V], count int) (iter.Seq2[K, V], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if count <= 0 {
		return seq2, nil
	}
	vv := iterhelper.Collect2(seq2)
	if count >= len(vv)/2 {
		return iterhelper.Empty2[K, V](), nil
	}
	r, err := iterhelper.Var2[K, V](vv[:len(vv)-count*2]...)
	if err != nil {
		return nil, errorhelper.CallerError(err)
	}
	return r, nil
}

// SkipWhile bypasses pairs of values in a [sequence] as long as a specified condition is true
// and then returns the remaining ones.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func SkipWhile[K, V any](seq2 iter.Seq2[K, V], pred func(K, V) bool) (iter.Seq2[K, V], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if pred == nil {
		return nil, errorhelper.CallerError(ErrNilPredicate)
	}
	return func(yield func(K, V) bool) {
			rest := false
			for k, v := range seq2 {
				if !rest {
					if pred(k, v) {
						continue
					} else {
						rest = true
					}
				}
				if !yield(k, v) {
					return
				}
			}
		},
		nil
}

// SkipWhileIdx bypasses pairs of values in a [sequence] as long as a specified condition is true
// and then returns the remaining ones.
// The pair's index is used in the logic of the predicate function.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func SkipWhileIdx[K, V any](seq2 iter.Seq2[K, V], pred func(K, V, int) bool) (iter.Seq2[K, V], error) {
	if seq2 == nil {
		return nil, errorhelper.CallerError(ErrNilInput)
	}
	if pred == nil {
		return nil, errorhelper.CallerError(ErrNilPredicate)
	}
	return func(yield func(K, V) bool) {
			rest := false
			i := 0
			for k, v := range seq2 {
				if !rest {
					if pred(k, v, i) {
						i++
						continue
					} else {
						rest = true
					}
				}
				if !yield(k, v) {
					return
				}
			}
		},
		nil
}
