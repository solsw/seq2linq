package seq2linq

import (
	"errors"
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/generichelper"
	"github.com/solsw/iterhelper"
)

func TestSingle(t *testing.T) {
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
		{name: "MultiElements",
			seq2:        sec2_int_word(),
			expectedErr: ErrMultiElements,
		},
		{name: "SingleElementSequence",
			seq2:  errorhelper.Must(iterhelper.Var2[int, string](3, "three")),
			wantK: 3,
			wantV: "three",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotV, gotErr := Single(tt.seq2)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("Single() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("Single() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("Single() succeeded unexpectedly")
			}
			if (gotK != tt.wantK) || (gotV != tt.wantV) {
				t.Errorf("Single(): %v, want: %v",
					generichelper.Tuple2[int, string]{Item1: gotK, Item2: gotV},
					generichelper.Tuple2[int, string]{Item1: tt.wantK, Item2: tt.wantV})
			}
		})
	}
}

func TestSinglePred(t *testing.T) {
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
			expectedErr: ErrNilInput,
		},
		{name: "NilPredicate",
			seq2:        sec2_int_word(),
			pred:        nil,
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
		{name: "MultiMatch",
			seq2:        sec2_int_word(),
			pred:        func(_ int, v string) bool { return len(v) > 3 },
			expectedErr: ErrMultiMatch,
		},
		{name: "SinglePred1",
			seq2:  errorhelper.Must(iterhelper.Var2[int, string](3, "three")),
			pred:  func(k int, v string) bool { return k == 3 && v == "three" },
			wantK: 3,
			wantV: "three",
		},
		{name: "SinglePred2",
			seq2:  sec2_int_word(),
			pred:  func(_ int, v string) bool { return v == "three" },
			wantK: 3,
			wantV: "three",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotV, gotErr := SinglePred(tt.seq2, tt.pred)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("SinglePred() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("SinglePred() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("SinglePred() succeeded unexpectedly")
			}
			if (gotK != tt.wantK) || (gotV != tt.wantV) {
				t.Errorf("SinglePred(): %v, want: %v",
					generichelper.Tuple2[int, string]{Item1: gotK, Item2: gotV},
					generichelper.Tuple2[int, string]{Item1: tt.wantK, Item2: tt.wantV})
			}
		})
	}
}

func TestSingleOrDefault(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		defK        int
		defV        string
		wantK       int
		wantV       string
		expectedErr error
	}{
		{name: "NilInput",
			seq2:        nil,
			expectedErr: ErrNilInput,
		},
		{name: "MultiElements",
			seq2:        sec2_int_word(),
			expectedErr: ErrMultiElements,
		},
		{name: "EmptyInput0",
			seq2:  iterhelper.Empty2[int, string](),
			wantK: generichelper.ZeroValue[int](),
			wantV: generichelper.ZeroValue[string](),
		},
		{name: "EmptyInput1",
			seq2:  iterhelper.Empty2[int, string](),
			defK:  6,
			defV:  "six",
			wantK: 6,
			wantV: "six",
		},
		{name: "SingleElementSequence",
			seq2:  errorhelper.Must(iterhelper.Var2[int, string](3, "three")),
			defK:  6,
			defV:  "six",
			wantK: 3,
			wantV: "three",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotV, gotErr := SingleOrDefault(tt.seq2, tt.defK, tt.defV)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("SingleOrDefault() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("Single() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("SingleOrDefault() succeeded unexpectedly")
			}
			if (gotK != tt.wantK) || (gotV != tt.wantV) {
				t.Errorf("SingleOrDefault(): %v, want: %v",
					generichelper.Tuple2[int, string]{Item1: gotK, Item2: gotV},
					generichelper.Tuple2[int, string]{Item1: tt.wantK, Item2: tt.wantV})
			}
		})
	}
}
