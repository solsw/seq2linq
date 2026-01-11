package seq2linq

import (
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func TestSelectMany(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		sel         func(int, string) iter.Seq2[int, string]
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "Projection0",
			seq2: sec2_int_word(),
			sel: func(i int, s string) iter.Seq2[int, string] {
				return errorhelper.Must(iterhelper.Var2[int, string](i, s))
			},
			want: sec2_int_word(),
		},
		{name: "Projection1",
			seq2: errorhelper.Must(iterhelper.Var2[int, string](0, "0", 1, "1")),
			sel: func(i int, s string) iter.Seq2[int, string] {
				return errorhelper.Must(iterhelper.Var2[int, string](i, s, i, s))
			},
			want: errorhelper.Must(iterhelper.Var2[int, string](0, "0", 0, "0", 1, "1", 1, "1")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := SelectMany(tt.seq2, tt.sel)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("SelectMany() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("SelectMany() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("SelectMany(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}

func TestSelectManyIdx(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		sel         func(int, string, int) iter.Seq2[int, string]
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "Projection0",
			seq2: sec2_int_word(),
			sel: func(i int, s string, _ int) iter.Seq2[int, string] {
				return errorhelper.Must(iterhelper.Var2[int, string](i, s))
			},
			want: sec2_int_word(),
		},
		{name: "Projection1",
			seq2: errorhelper.Must(iterhelper.Var2[int, string](1, "1", 2, "2")),
			sel: func(i int, s string, idx int) iter.Seq2[int, string] {
				return errorhelper.Must(iterhelper.Var2[int, string](i+idx, s, i-idx, s))
			},
			want: errorhelper.Must(iterhelper.Var2[int, string](1, "1", 1, "1", 3, "2", 1, "2")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := SelectManyIdx(tt.seq2, tt.sel)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("SelectManyIdx() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("SelectManyIdx() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("SelectManyIdx(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}
