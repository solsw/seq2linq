package seq2linq

import (
	"errors"
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func TestRepeat(t *testing.T) {
	tests := []struct {
		name        string
		k           int
		v           string
		count       int
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "NegativeCount",
			k:           1,
			v:           "one",
			count:       -1,
			expectedErr: ErrNegativeCount,
		},
		{name: "EmptyRepeat",
			k:     1,
			v:     "one",
			count: 0,
			want:  iterhelper.Empty2[int, string](),
		},
		{name: "SimpleRepeat",
			k:     1,
			v:     "one",
			count: 2,
			want:  errorhelper.Must(iterhelper.Var2[int, string](1, "one", 1, "one")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Repeat(tt.k, tt.v, tt.count)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("Repeat() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("Repeat() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("Repeat() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("Repeat(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}
