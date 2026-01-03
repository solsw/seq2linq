package seq2linq

import (
	"errors"
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func TestConcat(t *testing.T) {
	inst2 := errorhelper.Must(iterhelper.Var2[int, string](1, "one", 2, "two"))
	tests := []struct {
		name        string
		seq2s       []iter.Seq2[int, string]
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "EmptyInput",
			seq2s:       []iter.Seq2[int, string]{},
			expectedErr: ErrEmptyInput,
		},
		{name: "NilInput",
			seq2s: []iter.Seq2[int, string]{
				iterhelper.Empty2[int, string](),
				nil,
				iterhelper.Empty2[int, string](),
			},
			expectedErr: ErrNilInput,
		},
		{name: "Empty",
			seq2s: []iter.Seq2[int, string]{
				iterhelper.Empty2[int, string](),
				iterhelper.Empty2[int, string](),
				iterhelper.Empty2[int, string](),
			},
			want: iterhelper.Empty2[int, string](),
		},
		{name: "Empty1",
			seq2s: []iter.Seq2[int, string]{
				iterhelper.Empty2[int, string](),
				errorhelper.Must(iterhelper.Var2[int, string](1, "one", 3, "three", 5, "five", 9, "nine")),
			},
			want: errorhelper.Must(iterhelper.Var2[int, string](1, "one", 3, "three", 5, "five", 9, "nine")),
		},
		{name: "Empty2",
			seq2s: []iter.Seq2[int, string]{
				errorhelper.Must(iterhelper.Var2[int, string](5, "five", 9, "nine")),
				iterhelper.Empty2[int, string](),
				inst2,
			},
			want: errorhelper.Must(iterhelper.Var2[int, string](5, "five", 9, "nine", 1, "one", 2, "two")),
		},
		{name: "Concatenation",
			seq2s: []iter.Seq2[int, string]{
				errorhelper.Must(iterhelper.Var2[int, string](1, "one", 2, "two")),
				errorhelper.Must(iterhelper.Var2[int, string](3, "three", 4, "four")),
			},
			want: errorhelper.Must(iterhelper.Var2[int, string](1, "one", 2, "two", 3, "three", 4, "four")),
		},
		{name: "Concatenation2",
			seq2s: []iter.Seq2[int, string]{
				errorhelper.Must(Repeat(1, "one", 1)),
				errorhelper.Must(Repeat(2, "two", 2)),
			},
			want: errorhelper.Must(iterhelper.Var2[int, string](1, "one", 2, "two", 2, "two")),
		},
		{name: "SameSequence",
			seq2s: []iter.Seq2[int, string]{inst2, inst2, inst2},
			want:  errorhelper.Must(iterhelper.Var2[int, string](1, "one", 2, "two", 1, "one", 2, "two", 1, "one", 2, "two")),
		},
		{name: "SameSequence2",
			seq2s: []iter.Seq2[int, string]{
				errorhelper.Must(Take(sec2_int_word(), 2)),
				errorhelper.Must(Skip(sec2_int_word(), 2)),
			},
			want: sec2_int_word(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Concat(tt.seq2s...)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("Concat() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("Concat() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("Concat() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("Concat(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}
