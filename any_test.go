package seq2linq

import (
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func TestAny(t *testing.T) {
	tests := []struct {
		name        string
		in          iter.Seq2[int, string]
		want        bool
		expectedErr error
	}{
		{name: "EmptyReturnsFalse",
			in:   iterhelper.Empty2[int, string](),
			want: false,
		},
		{name: "NonEmptyReturnsTrue",
			in:   sec2_int_word(),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Any(tt.in)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("Any() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("Any() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyPred(t *testing.T) {
	tests := []struct {
		name        string
		in          iter.Seq2[int, string]
		pred        func(int, string) bool
		want        bool
		expectedErr error
	}{
		{name: "EmptyReturnsFalse",
			in:   iterhelper.Empty2[int, string](),
			pred: func(int, string) bool { return true },
			want: false,
		},
		{name: "PredMatchesPair",
			in:   sec2_int_word(),
			pred: func(i int, s string) bool { return i == 2 && s == "two" },
			want: true,
		},
		{name: "PredMatchesNoPairs",
			in:   sec2_int_word(),
			pred: func(i int, s string) bool { return i == 1 && s == "two" },
			want: false,
		},
		{name: "SequenceIsNotEvaluatedAfterFirstMatch",
			in: errorhelper.Must(Select(
				errorhelper.Must(iterhelper.Var2[int, string](6, "6", 4, "4", 0, "0", 2, "2")),
				func(i int, s string) (int, string) { return 12 / i, s },
			)),
			pred: func(i int, _ string) bool { return i > 2 },
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := AnyPred(tt.in, tt.pred)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("AnyPred() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("AnyPred() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("AnyPred() = %v, want %v", got, tt.want)
			}
		})
	}
}
