package seq2linq

import (
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func TestConcat(t *testing.T) {
	inst2 := errorhelper.Must(iterhelper.Var2[int, string](1, "one", 2, "two"))
	tests := []struct {
		name        string
		in1         iter.Seq2[int, string]
		in2         iter.Seq2[int, string]
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "Empty",
			in1:  iterhelper.Empty2[int, string](),
			in2:  iterhelper.Empty2[int, string](),
			want: iterhelper.Empty2[int, string](),
		},
		{name: "Empty1",
			in1:  iterhelper.Empty2[int, string](),
			in2:  errorhelper.Must(iterhelper.Var2[int, string](1, "one", 3, "three", 5, "five", 9, "nine")),
			want: errorhelper.Must(iterhelper.Var2[int, string](1, "one", 3, "three", 5, "five", 9, "nine")),
		},
		{name: "Empty2",
			in1:  errorhelper.Must(iterhelper.Var2[int, string](1, "one", 3, "three", 5, "five", 9, "nine")),
			in2:  iterhelper.Empty2[int, string](),
			want: errorhelper.Must(iterhelper.Var2[int, string](1, "one", 3, "three", 5, "five", 9, "nine")),
		},
		{name: "Concatenation",
			in1:  errorhelper.Must(iterhelper.Var2[int, string](1, "one", 2, "two")),
			in2:  errorhelper.Must(iterhelper.Var2[int, string](3, "three", 4, "four")),
			want: errorhelper.Must(iterhelper.Var2[int, string](1, "one", 2, "two", 3, "three", 4, "four")),
		},
		{name: "Concatenation2",
			in1:  errorhelper.Must(Repeat(1, "one", 1)),
			in2:  errorhelper.Must(Repeat(2, "two", 2)),
			want: errorhelper.Must(iterhelper.Var2[int, string](1, "one", 2, "two", 2, "two")),
		},
		{name: "SameSequence",
			in1:  inst2,
			in2:  inst2,
			want: errorhelper.Must(iterhelper.Var2[int, string](1, "one", 2, "two", 1, "one", 2, "two")),
		},
		{name: "SameSequence2",
			in1:  errorhelper.Must(Take(sec2_int_word(), 2)),
			in2:  errorhelper.Must(Skip(sec2_int_word(), 2)),
			want: sec2_int_word(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Concat(tt.in1, tt.in2)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("Concat() failed: %v", gotErr)
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
