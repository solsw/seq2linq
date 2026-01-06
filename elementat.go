package seq2linq

import (
	"errors"
	"iter"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
)

// ElementAt returns the element (a pair of values) at a specified index in a [sequence].
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func ElementAt[K, V any](seq2 iter.Seq2[K, V], index int) (K, V, error) {
	if seq2 == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilInput)
	}
	if index < 0 {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrIndexOutOfRange)
	}
	i := 0
	for k, v := range seq2 {
		if i == index {
			return k, v, nil
		}
		i++
	}
	return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrIndexOutOfRange)
}

// ElementAtOrDefault returns the element (a pair of values) at a specified index in
// a [sequence] or a pair of corresponding [zero values] if the index is out of range.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
// [zero values]: https://go.dev/ref/spec#The_zero_value
func ElementAtOrDefault[K, V any](seq2 iter.Seq2[K, V], index int) (K, V, error) {
	if seq2 == nil {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), errorhelper.CallerError(ErrNilInput)
	}
	k, v, err := ElementAt(seq2, index)
	if errors.Is(err, ErrIndexOutOfRange) {
		return generichelper.ZeroValue[K](), generichelper.ZeroValue[V](), nil
	}
	return k, v, nil
}
