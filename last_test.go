package seq2linq

import (
	"errors"
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
	"github.com/solsw/iterhelper"
)

func TestLast(t *testing.T) {
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
		{name: "Last1",
			seq2:  errorhelper.Must(iterhelper.Var2[int, string](9, "nine")),
			wantK: 9,
			wantV: "nine",
		},
		{name: "Last",
			seq2:  sec2_int_word(),
			wantK: 9,
			wantV: "nine",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotV, gotErr := Last(tt.seq2)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("Last() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("Last() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("Last() succeeded unexpectedly")
			}
			if (gotK != tt.wantK) || (gotV != tt.wantV) {
				t.Errorf("Last(): %v, want: %v",
					generichelper.Tuple2[int, string]{Item1: gotK, Item2: gotV},
					generichelper.Tuple2[int, string]{Item1: tt.wantK, Item2: tt.wantV})
			}
		})
	}
}

func TestLastPred(t *testing.T) {
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
		{name: "LastPred",
			seq2:  sec2_int_word(),
			pred:  func(_ int, v string) bool { return len(v) > 4 },
			wantK: 8,
			wantV: "eight",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotV, gotErr := LastPred(tt.seq2, tt.pred)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("LastPred() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("LastPred() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("LastPred() succeeded unexpectedly")
			}
			if (gotK != tt.wantK) || (gotV != tt.wantV) {
				t.Errorf("Last(): %v, want: %v",
					generichelper.Tuple2[int, string]{Item1: gotK, Item2: gotV},
					generichelper.Tuple2[int, string]{Item1: tt.wantK, Item2: tt.wantV})
			}
		})
	}
}

func TestLastOrDefault(t *testing.T) {
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
			wantK: 9,
			wantV: "nine",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotV, gotErr := LastOrDefault(tt.seq2)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("LastOrDefault() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("LastOrDefault() succeeded unexpectedly")
			}
			if (gotK != tt.wantK) || (gotV != tt.wantV) {
				t.Errorf("LastOrDefault(): %v, want: %v",
					generichelper.Tuple2[int, string]{Item1: gotK, Item2: gotV},
					generichelper.Tuple2[int, string]{Item1: tt.wantK, Item2: tt.wantV})
			}
		})
	}
}

func TestLastOrDefaultPred(t *testing.T) {
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
			wantK: 9,
			wantV: "nine",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotV, gotErr := LastOrDefaultPred(tt.seq2, tt.pred)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("LastOrDefaultPred() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("LastOrDefaultPred() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("LastOrDefaultPred() succeeded unexpectedly")
			}
			if (gotK != tt.wantK) || (gotV != tt.wantV) {
				t.Errorf("LastOrDefaultPred(): %v, want: %v",
					generichelper.Tuple2[int, string]{Item1: gotK, Item2: gotV},
					generichelper.Tuple2[int, string]{Item1: tt.wantK, Item2: tt.wantV})
			}
		})
	}
}
