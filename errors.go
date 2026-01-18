package seq2linq

import (
	"errors"
)

var (
	ErrEmptyInput      = errors.New("empty input sequence")
	ErrIndexOutOfRange = errors.New("index out of range")
	ErrNegativeCount   = errors.New("negative count")
	ErrNilEqual        = errors.New("nil equal function")
	ErrNilInput        = errors.New("nil input sequence")
	ErrNilLess         = errors.New("nil less function")
	ErrNilNext         = errors.New("nil next function")
	ErrNilPredicate    = errors.New("nil predicate function")
	ErrNilSelector     = errors.New("nil selector function")
	ErrNoMatch         = errors.New("no predicate match")
)
