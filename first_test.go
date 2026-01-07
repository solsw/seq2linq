package seq2linq

import (
	"errors"
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
	"github.com/solsw/iterhelper"
)

func TestFirst(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		wantK       int
		wantV       string
		expectedErr error
	}{
		{name: "NilInput",
			seq2:        nil,
			expectedErr: ErrNilInput,
		},
		{name: "EmptyInput",
			seq2:        iterhelper.Empty2[int, string](),
			expectedErr: ErrEmptyInput,
		},
		{name: "First1",
			seq2:  errorhelper.Must(iterhelper.Var2[int, string](9, "nine")),
			wantK: 9,
			wantV: "nine",
		},
		{name: "First",
			seq2:  sec2_int_word(),
			wantK: 0,
			wantV: "zero",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotV, gotErr := First(tt.seq2)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("First() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("First() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("First() succeeded unexpectedly")
			}
			if (gotK != tt.wantK) || (gotV != tt.wantV) {
				t.Errorf("First(): %v, want: %v",
					generichelper.Tuple2[int, string]{Item1: gotK, Item2: gotV},
					generichelper.Tuple2[int, string]{Item1: tt.wantK, Item2: tt.wantV})
			}
		})
	}
}

func TestFirstPred(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		pred        func(int, string) bool
		wantK       int
		wantV       string
		expectedErr error
	}{
		{name: "NilInput",
			seq2:        nil,
			pred:        func(int, string) bool { return true },
			expectedErr: ErrNilInput,
		},
		{name: "NilPredicate",
			seq2:        sec2_int_word(),
			expectedErr: ErrNilPredicate,
		},
		{name: "EmptyInput",
			seq2:        iterhelper.Empty2[int, string](),
			pred:        func(int, string) bool { return true },
			expectedErr: ErrEmptyInput,
		},
		{name: "NoMatch",
			seq2:        sec2_int_word(),
			pred:        func(_ int, v string) bool { return len(v) > 44 },
			expectedErr: ErrNoMatch,
		},
		{name: "FirstPred",
			seq2:  sec2_int_word(),
			pred:  func(_ int, v string) bool { return len(v) > 4 },
			wantK: 3,
			wantV: "three",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotV, gotErr := FirstPred(tt.seq2, tt.pred)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("FirstPred() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("FirstPred() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("FirstPred() succeeded unexpectedly")
			}
			if (gotK != tt.wantK) || (gotV != tt.wantV) {
				t.Errorf("FirstPred(): %v, want: %v",
					generichelper.Tuple2[int, string]{Item1: gotK, Item2: gotV},
					generichelper.Tuple2[int, string]{Item1: tt.wantK, Item2: tt.wantV})
			}
		})
	}
}

func TestFirstOrDefault(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		wantK       int
		wantV       string
		expectedErr error
	}{
		{name: "EmptySequence",
			seq2:  iterhelper.Empty2[int, string](),
			wantK: generichelper.ZeroValue[int](),
			wantV: generichelper.ZeroValue[string](),
		},
		{name: "SingleElementSequence",
			seq2:  errorhelper.Must(iterhelper.Var2[int, string](9, "nine")),
			wantK: 9,
			wantV: "nine",
		},
		{name: "MultipleElementSequence",
			seq2:  sec2_int_word(),
			wantK: 0,
			wantV: "zero",
		},
		{name: "EarlyOutAfterFirstElement",
			seq2: errorhelper.Must(Select(
				errorhelper.Must(iterhelper.Var2[int, string](6, "6", 4, "4", 0, "0", 2, "2")),
				func(i int, s string) (int, string) { return 12 / i, s },
			)),
			wantK: 2,
			wantV: "6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotV, gotErr := FirstOrDefault(tt.seq2)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("FirstOrDefault() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("FirstOrDefault() succeeded unexpectedly")
			}
			if (gotK != tt.wantK) || (gotV != tt.wantV) {
				t.Errorf("FirstOrDefault(): %v, want: %v",
					generichelper.Tuple2[int, string]{Item1: gotK, Item2: gotV},
					generichelper.Tuple2[int, string]{Item1: tt.wantK, Item2: tt.wantV})
			}
		})
	}
}

func TestFirstOrDefaultPred(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		pred        func(int, string) bool
		wantK       int
		wantV       string
		expectedErr error
	}{
		{name: "NilInput",
			seq2:        nil,
			pred:        func(int, string) bool { return true },
			expectedErr: ErrNilInput,
		},
		{name: "NilPredicate",
			seq2:        sec2_int_word(),
			pred:        nil,
			expectedErr: ErrNilPredicate,
		},
		{name: "EmptySequence",
			seq2:  iterhelper.Empty2[int, string](),
			pred:  func(int, string) bool { return true },
			wantK: generichelper.ZeroValue[int](),
			wantV: generichelper.ZeroValue[string](),
		},
		{name: "SingleElementSequenceWithMatchingPredicate",
			seq2:  errorhelper.Must(iterhelper.Var2[int, string](3, "three")),
			pred:  func(_ int, v string) bool { return len(v) > 4 },
			wantK: 3,
			wantV: "three",
		},
		{name: "SingleElementSequenceWithNonMatchingPredicate",
			seq2:  errorhelper.Must(iterhelper.Var2[int, string](3, "three")),
			pred:  func(_ int, v string) bool { return len(v) < 4 },
			wantK: generichelper.ZeroValue[int](),
			wantV: generichelper.ZeroValue[string](),
		},
		{name: "MultipleElementSequenceWithNoPredicateMatches",
			seq2:  errorhelper.Must(iterhelper.Var2[int, string](3, "three", 5, "five", 7, "seven", 9, "nine")),
			pred:  func(_ int, v string) bool { return len(v) > 44 },
			wantK: generichelper.ZeroValue[int](),
			wantV: generichelper.ZeroValue[string](),
		},
		{name: "MultipleElementSequenceWithSinglePredicateMatch",
			seq2:  errorhelper.Must(iterhelper.Var2[int, string](3, "three", 5, "five", 7, "seven", 9, "nine")),
			pred:  func(k int, _ string) bool { return k == 5 },
			wantK: 5,
			wantV: "five",
		},
		{name: "MultipleElementSequenceWithMultiplePredicateMatches",
			seq2:  errorhelper.Must(iterhelper.Var2[int, string](3, "three", 5, "five", 7, "seven", 9, "nine")),
			pred:  func(_ int, v string) bool { return len(v) == 4 },
			wantK: 5,
			wantV: "five",
		},
		{name: "EarlyOutAfterFirstElementWithPredicate",
			seq2: errorhelper.Must(Select(
				errorhelper.Must(iterhelper.Var2[int, string](6, "6", 4, "4", 0, "0", 2, "2")),
				func(i int, s string) (int, string) { return 12 / i, s },
			)),
			pred:  func(k int, _ string) bool { return k == 3 },
			wantK: 3,
			wantV: "4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotV, gotErr := FirstOrDefaultPred(tt.seq2, tt.pred)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("FirstOrDefaultPred() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("FirstOrDefaultPred() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("FirstOrDefaultPred() succeeded unexpectedly")
			}
			if (gotK != tt.wantK) || (gotV != tt.wantV) {
				t.Errorf("FirstOrDefaultPred(): %v, want: %v",
					generichelper.Tuple2[int, string]{Item1: gotK, Item2: gotV},
					generichelper.Tuple2[int, string]{Item1: tt.wantK, Item2: tt.wantV})
			}
		})
	}
}
