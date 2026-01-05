package seq2linq

import (
	"iter"
	"strings"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func TestDistinct(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "Distinct",
			seq2: errorhelper.Must(Concat(sec2_int_word(), sec2_int_word())),
			want: sec2_int_word(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Distinct(tt.seq2)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("Distinct() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("Distinct() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("Distinct(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}

func TestDistinctEq(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		eq          func(int, string, int, string) bool
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "DistinctEq",
			seq2: errorhelper.Must(Concat(
				sec2_int_word(),
				errorhelper.Must(iterhelper.Var2[int, string](1, "ONE", 3, "THREE", 5, "Five", 9, "nine")),
				sec2_int_word(),
			)),
			eq: func(_ int, s1 string, _ int, s2 string) bool {
				return strings.EqualFold(s1, s2)
			},
			want: sec2_int_word(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := DistinctEq(tt.seq2, tt.eq)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("DistinctEq() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("DistinctEq() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("DistinctEq(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}
