package seq2linq

import (
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func TestAppend(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		k           int
		v           string
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "EmptyIn",
			seq2: iterhelper.Empty2[int, string](),
			k:    1,
			v:    "one",
			want: errorhelper.Must(iterhelper.Var2[int, string](1, "one")),
		},
		{name: "Append",
			seq2: errorhelper.Must(iterhelper.Var2[int, string](2, "two", 1, "one")),
			k:    2,
			v:    "two",
			want: errorhelper.Must(iterhelper.Var2[int, string](2, "two", 1, "one", 2, "two")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Append(tt.seq2, tt.k, tt.v)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("Append() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("Append() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("Append(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}
