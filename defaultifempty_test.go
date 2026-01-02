package seq2linq

import (
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func TestDefaultIfEmpty(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "EmptyInput",
			seq2: iterhelper.Empty2[int, string](),
			want: errorhelper.Must(iterhelper.Var2[int, string](0, "")),
		},
		{name: "NonEmptyInput",
			seq2: sec2_int_word(),
			want: sec2_int_word(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := DefaultIfEmpty(tt.seq2)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("DefaultIfEmpty() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("DefaultIfEmpty() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("DefaultIfEmpty(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}

func TestDefaultIfEmptyDef(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		defK        int
		defV        string
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "EmptyInput",
			seq2: iterhelper.Empty2[int, string](),
			defK: 5,
			defV: "five",
			want: errorhelper.Must(iterhelper.Var2[int, string](5, "five")),
		},
		{name: "NonEmptyInput",
			seq2: sec2_int_word(),
			defK: 5,
			defV: "five",
			want: sec2_int_word(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := DefaultIfEmptyDef(tt.seq2, tt.defK, tt.defV)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("DefaultIfEmptyDef() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("DefaultIfEmptyDef() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("DefaultIfEmptyDef(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}
