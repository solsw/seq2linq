package seq2linq

import (
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func TestZip(t *testing.T) {
	tests := []struct {
		name        string
		first       iter.Seq2[int, string]
		second      iter.Seq2[int, string]
		sel         func(int, string, int, string) (int, string)
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "EmptyInput",
			first:  iterhelper.Empty2[int, string](),
			second: iterhelper.Empty2[int, string](),
			sel:    func(int, string, int, string) (int, string) { return 0, "" },
			want:   iterhelper.Empty2[int, string](),
		},
		{name: "Zip",
			first:  errorhelper.Must(iterhelper.Var2[int, string](1, "1", 2, "2")),
			second: sec2_int_word(),
			sel:    func(k1 int, v1 string, k2 int, v2 string) (int, string) { return k1 + k2, v1 + v2 },
			want:   errorhelper.Must(iterhelper.Var2[int, string](1, "1zero", 3, "2one")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Zip(tt.first, tt.second, tt.sel)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("Zip() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("Zip() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("Zip(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}
