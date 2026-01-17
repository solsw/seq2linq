package seq2linq

import (
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func TestOfType(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[any, any]
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "EmptyInput",
			seq2: iterhelper.Empty2[any, any](),
			want: iterhelper.Empty2[int, string](),
		},
		{name: "OfType1",
			seq2: errorhelper.Must(iterhelper.Var2[any, any](0, "1", 2, "3")),
			want: errorhelper.Must(iterhelper.Var2[int, string](0, "1", 2, "3")),
		},
		{name: "OfType2",
			seq2: errorhelper.Must(iterhelper.Var2[any, any](0, "1", 2, 3, 4, "5", "6", 7)),
			want: errorhelper.Must(iterhelper.Var2[int, string](0, "1", 4, "5")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := OfType[any, any, int, string](tt.seq2)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("OfType() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("OfType() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("OfType(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}
