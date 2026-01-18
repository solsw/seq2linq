package seq2linq

import (
	"errors"
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func TestOrderByLs(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		less        func(int, string, int, string) bool
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "NilInput",
			seq2:        nil,
			less:        func(int, string, int, string) bool { return true },
			expectedErr: ErrNilInput,
		},
		{name: "NilLess",
			seq2:        sec2_int_word(),
			less:        nil,
			expectedErr: ErrNilLess,
		},
		{name: "EmptyInput",
			seq2: iterhelper.Empty2[int, string](),
			less: func(int, string, int, string) bool { return true },
			want: iterhelper.Empty2[int, string](),
		},
		{name: "OrderByLs",
			seq2: sec2_int_word(),
			less: func(_ int, v1 string, _ int, v2 string) bool { return v1 < v2 },
			want: errorhelper.Must(iterhelper.Var2[int, string](8, "eight", 5, "five",
				4, "four", 9, "nine", 1, "one", 7, "seven", 6, "six", 3, "three", 2, "two", 0, "zero")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := OrderByLs(tt.seq2, tt.less)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("OrderByLs() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("OrderByLs() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("OrderByLs() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("OrderByLs(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}
