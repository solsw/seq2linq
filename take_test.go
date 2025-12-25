package seq2linq

import (
	"errors"
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func TestTake(t *testing.T) {
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
			want:  iterhelper.Empty2[int, string](),
		},
		{name: "CountShorterThanInputLength",
			seq2:  sec2_int_word(),
			count: 3,
			want:  errorhelper.Must(iterhelper.Var2[int, string](0, "zero", 1, "one", 2, "two")),
		},
		{name: "CountEqualToInputLength",
			seq2:  sec2_int_word(),
			count: 10,
			want:  sec2_int_word(),
		},
		{name: "CountGreaterThanInputLength",
			seq2:  sec2_int_word(),
			count: 100,
			want:  sec2_int_word(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Take(tt.seq2, tt.count)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("Take() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("Take() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("Take() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("Take(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}

func TestTakeLast(t *testing.T) {
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
			want:  iterhelper.Empty2[int, string](),
		},
		{name: "CountShorterThanInputLength",
			seq2:  sec2_int_word(),
			count: 3,
			want:  errorhelper.Must(iterhelper.Var2[int, string](7, "seven", 8, "eight", 9, "nine")),
		},
		{name: "CountEqualToInputLength",
			seq2:  sec2_int_word(),
			count: 10,
			want:  sec2_int_word(),
		},
		{name: "CountGreaterThanInputLength",
			seq2:  sec2_int_word(),
			count: 100,
			want:  sec2_int_word(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := TakeLast(tt.seq2, tt.count)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("TakeLast() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("TakeLast() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("TakeLast(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}

func TestTakeWhile(t *testing.T) {
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
			want: iterhelper.Empty2[int, string](),
		},
		{name: "PredicateMatchingSomeElements",
			seq2: sec2_int_word(),
			pred: func(_ int, s string) bool { return len(s) < 5 },
			want: errorhelper.Must(iterhelper.Var2[int, string](0, "zero", 1, "one", 2, "two")),
		},
		{name: "PredicateMatchingAllElements",
			seq2: sec2_int_word(),
			pred: func(_ int, s string) bool { return len(s) < 10 },
			want: sec2_int_word(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := TakeWhile(tt.seq2, tt.pred)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("TakeWhile() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("TakeWhile() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("TakeWhile(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}

func TestTakeWhileIdx(t *testing.T) {
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
			want: iterhelper.Empty2[int, string](),
		},
		{name: "PredicateMatchingSomeElements",
			seq2: sec2_int_word(),
			pred: func(_ int, s string, idx int) bool { return len(s) != idx },
			want: errorhelper.Must(iterhelper.Var2[int, string](0, "zero", 1, "one", 2, "two", 3, "three")),
		},
		{name: "PredicateMatchingAllElements",
			seq2: sec2_int_word(),
			pred: func(int, string, int) bool { return true },
			want: sec2_int_word(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := TakeWhileIdx(tt.seq2, tt.pred)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("TakeWhileIdx() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("TakeWhileIdx() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("TakeWhileIdx(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}
