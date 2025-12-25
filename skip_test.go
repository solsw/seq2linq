package seq2linq

import (
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func TestSkip(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		count       int
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "NegativeOrZeroCount",
			seq2:  sec2_int_word(),
			count: -5,
			want:  sec2_int_word(),
		},
		{name: "CountShorterThanInputLength",
			seq2:  sec2_int_word(),
			count: 7,
			want:  errorhelper.Must(iterhelper.Var2[int, string](7, "seven", 8, "eight", 9, "nine")),
		},
		{name: "CountEqualToInputLength",
			seq2:  sec2_int_word(),
			count: 10,
			want:  iterhelper.Empty2[int, string](),
		},
		{name: "CountGreaterThanInputLength",
			seq2:  sec2_int_word(),
			count: 100,
			want:  iterhelper.Empty2[int, string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Skip(tt.seq2, tt.count)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("Skip() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("Skip() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("Skip(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}

func TestSkipLast(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		count       int
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "NegativeOrZeroCount",
			seq2:  sec2_int_word(),
			count: 0,
			want:  sec2_int_word(),
		},
		{name: "CountShorterThanInputLength",
			seq2:  sec2_int_word(),
			count: 7,
			want:  errorhelper.Must(iterhelper.Var2[int, string](0, "zero", 1, "one", 2, "two")),
		},
		{name: "CountEqualToInputLength",
			seq2:  sec2_int_word(),
			count: 10,
			want:  iterhelper.Empty2[int, string](),
		},
		{name: "CountGreaterThanInputLength",
			seq2:  sec2_int_word(),
			count: 100,
			want:  iterhelper.Empty2[int, string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := SkipLast(tt.seq2, tt.count)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("SkipLast() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("SkipLast() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("SkipLast(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}

func TestSkipWhile(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		pred        func(int, string) bool
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "PredicateFailingFirstElement",
			seq2: sec2_int_word(),
			pred: func(_ int, s string) bool { return len(s) > 4 },
			want: sec2_int_word(),
		},
		{name: "PredicateMatchingSomeElements",
			seq2: sec2_int_word(),
			pred: func(_ int, s string) bool { return len(s) < 5 },
			want: errorhelper.Must(iterhelper.Var2[int, string](3, "three", 4, "four", 5, "five", 6, "six", 7, "seven", 8, "eight", 9, "nine")),
		},
		{name: "PredicateMatchingAllElements",
			seq2: sec2_int_word(),
			pred: func(_ int, s string) bool { return len(s) < 100 },
			want: iterhelper.Empty2[int, string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := SkipWhile(tt.seq2, tt.pred)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("SkipWhile() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("SkipWhile() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("SkipWhile(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}

func TestSkipWhileIdx(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		pred        func(int, string, int) bool
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "PredicateFailingFirstElement",
			seq2: sec2_int_word(),
			pred: func(_ int, s string, idx int) bool { return idx+len(s) > 4 },
			want: sec2_int_word(),
		},
		{name: "PredicateMatchingSomeElements",
			seq2: sec2_int_word(),
			pred: func(_ int, s string, idx int) bool { return len(s) != idx },
			want: errorhelper.Must(iterhelper.Var2[int, string](4, "four", 5, "five", 6, "six", 7, "seven", 8, "eight", 9, "nine")),
		},
		{name: "PredicateMatchingAllElements",
			seq2: sec2_int_word(),
			pred: func(int, string, int) bool { return true },
			want: iterhelper.Empty2[int, string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := SkipWhileIdx(tt.seq2, tt.pred)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("SkipWhileIdx() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("SkipWhileIdx() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("SkipWhileIdx(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}
