package seq2linq

import (
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "EmptyInput",
			seq2: iterhelper.Empty2[int, string](),
			want: iterhelper.Empty2[int, string](),
		},
		{name: "Reverse",
			seq2: errorhelper.Must(iterhelper.Var2[int, string](1, "one", 2, "two", 3, "three")),
			want: errorhelper.Must(iterhelper.Var2[int, string](3, "three", 2, "two", 1, "one")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Reverse(tt.seq2)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("Reverse() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("Reverse() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("Reverse(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}
