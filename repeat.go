package seq2linq

import (
	"iter"

	"github.com/solsw/errorhelper"
)

// Repeat returnes an [iterator] that yields one repeated pair of values.
// 'count' is the number of times to repeat the pair.
//
// [iterator]: https://pkg.go.dev/iter#Seq2
func Repeat[K, V any](k K, v V, count int) (iter.Seq2[K, V], error) {
	if count < 0 {
		return nil, errorhelper.CallerError(ErrNegativeCount)
	}
	return func(yield func(K, V) bool) {
			for range count {
				if !yield(k, v) {
					return
				}
			}
		},
		nil
}
