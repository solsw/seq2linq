package seq2linq

import (
	"errors"
)

var (
	ErrEmptyInput    = errors.New("empty input")
	ErrNegativeCount = errors.New("negative count")
	ErrNilEqual      = errors.New("nil equal")
	ErrNilInput      = errors.New("nil input sequence")
	ErrNilNext       = errors.New("nil next")
	ErrNilPredicate  = errors.New("nil predicate")
	ErrNilSelector   = errors.New("nil selector")
)
