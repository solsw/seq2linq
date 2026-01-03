package seq2linq

import (
	"errors"
	"iter"
	"strconv"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func TestSelect(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		sel         func(int, string) (int, string)
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "NilInput",
			seq2:        nil,
			sel:         func(i int, s string) (int, string) { return i, s },
			expectedErr: ErrNilInput,
		},
		{name: "NilSelector",
			seq2:        sec2_int_word(),
			sel:         nil,
			expectedErr: ErrNilSelector,
		},
		{name: "EmptyInput",
			seq2: iterhelper.Empty2[int, string](),
			sel:  func(i int, s string) (int, string) { return i, s },
			want: iterhelper.Empty2[int, string](),
		},
		{name: "Projection1",
			seq2: sec2_int_word(),
			sel:  func(i int, s string) (int, string) { return i, s },
			want: sec2_int_word(),
		},
		{name: "Projection2",
			seq2: sec2_int_string(2),
			sel:  func(i int, s string) (int, string) { return i + i, s + s },
			want: errorhelper.Must(iterhelper.Var2[int, string](0, "00", 2, "11")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Select(tt.seq2, tt.sel)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("Select() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("Select() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("Select() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("Select(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}

func TestSelectIdx(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		sel         func(int, string, int) (string, int)
		want        iter.Seq2[string, int]
		expectedErr error
	}{
		{name: "EmptyInput",
			seq2: iterhelper.Empty2[int, string](),
			sel:  func(int, string, int) (string, int) { return "", 0 },
			want: iterhelper.Empty2[string, int](),
		},
		{name: "Projection",
			seq2: errorhelper.Must(Where(sec2_int_word(), func(i int, _ string) bool { return i < 5 && i%2 == 0 })),
			sel: func(i int, s string, idx int) (string, int) {
				return strconv.Itoa(idx) + ":" + s, i * i
			},
			want: errorhelper.Must(iterhelper.Var2[string, int]("0:zero", 0, "1:two", 4, "2:four", 16)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := SelectIdx(tt.seq2, tt.sel)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("SelectIdx() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("SelectIdx() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("SelectIdx() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("SelectIdx(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}
