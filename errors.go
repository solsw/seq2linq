package seq2linq

import (
	"errors"
)

var (
	ErrNilInput     = errors.New("nil input sequence")
	ErrNilNext      = errors.New("nil next")
	ErrNilPredicate = errors.New("nil predicate")
	ErrNilSelector  = errors.New("nil selector")
)
