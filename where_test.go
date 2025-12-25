package seq2linq

import (
	"errors"
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func TestWhere(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		pred        func(int, string) bool
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "NilInThrows",
			seq2:        nil,
			pred:        func(i int, _ string) bool { return i > 5 },
			expectedErr: ErrNilInput,
		},
		{name: "NilPredThrows",
			seq2:        sec2_int_word(),
			pred:        nil,
			expectedErr: ErrNilPredicate,
		},
		{name: "EmptyIn",
			seq2: iterhelper.Empty2[int, string](),
			pred: func(i int, _ string) bool { return i > 5 },
			want: iterhelper.Empty2[int, string](),
		},
		{name: "FalsePredicate",
			seq2: sec2_int_word(),
			pred: func(int, string) bool { return false },
			want: iterhelper.Empty2[int, string](),
		},
		{name: "TruePredicate",
			seq2: sec2_int_word(),
			pred: func(int, string) bool { return true },
			want: sec2_int_word(),
		},
		{name: "Filtering1",
			seq2: sec2_int_word(),
			pred: func(i int, s string) bool { return i > 5 && len(s) == 3 },
			want: errorhelper.Must(iterhelper.Var2[int, string](6, "six")),
		},
		{name: "Filtering2",
			seq2: sec2_int_word(),
			pred: func(i int, s string) bool { return i%2 == 1 && s[len(s)-1] == 'e' },
			want: errorhelper.Must(iterhelper.Var2[int, string](1, "one", 3, "three", 5, "five", 9, "nine")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Where(tt.seq2, tt.pred)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("Where() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("Where() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("Where() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("Where(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}

func TestWhereIdx(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		pred        func(int, string, int) bool
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "EmptyIn",
			seq2: iterhelper.Empty2[int, string](),
			pred: func(i int, _ string, _ int) bool { return i > 5 },
			want: iterhelper.Empty2[int, string](),
		},
		{name: "Filtering1",
			seq2: sec2_int_word(),
			pred: func(i int, s string, idx int) bool { return len(s) == idx },
			want: errorhelper.Must(iterhelper.Var2[int, string](4, "four")),
		},
		{name: "Filtering2",
			seq2: sec2_int_word(),
			pred: func(i int, s string, idx int) bool { return len(s) == (i+idx)/2 },
			want: errorhelper.Must(iterhelper.Var2[int, string](4, "four")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := WhereIdx(tt.seq2, tt.pred)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("WhereIdx() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("WhereIdx() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("WhereIdx() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("WhereIdx(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}
