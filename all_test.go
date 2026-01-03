package seq2linq

import (
	"errors"
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func TestAll(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		pred        func(int, string) bool
		want        bool
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
		{name: "EmptyAlwaysReturnsTrue",
			seq2: iterhelper.Empty2[int, string](),
			pred: func(int, string) bool { return false },
			want: true,
		},
		{name: "AlwaysTruePredReturnsTrue",
			seq2: sec2_int_word(),
			pred: func(int, string) bool { return true },
			want: true,
		},
		{name: "AlwaysFalsePredReturnsFalse",
			seq2: sec2_int_word(),
			pred: func(int, string) bool { return false },
			want: false,
		},
		{name: "PredMatchingNoElements",
			seq2: sec2_int_word(),
			pred: func(_ int, s string) bool { return len(s) < 3 },
			want: false,
		},
		{name: "PredMatchingSomeElements",
			seq2: sec2_int_word(),
			pred: func(_ int, s string) bool { return len(s) == 3 },
			want: false,
		},
		{name: "PredMatchingAllElements",
			seq2: sec2_int_word(),
			pred: func(_ int, s string) bool { return len(s) >= 3 },
			want: true,
		},
		{name: "SequenceIsNotEvaluatedAfterFirstNonMatch",
			seq2: errorhelper.Must(Select(
				errorhelper.Must(iterhelper.Var2[int, string](4, "4", 6, "6", 0, "0", 1, "1", 2, "2")),
				func(i int, s string) (int, string) { return 12 / i, s },
			)),
			pred: func(i int, _ string) bool { return i > 2 },
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := All(tt.seq2, tt.pred)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("All() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("All() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("All() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}
