package seq2linq

import (
	"errors"
	"iter"
	"testing"

	"github.com/solsw/errorhelper"
	"github.com/solsw/iterhelper"
)

func TestDistinctBy(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		keySel      func(int, string) int
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "NilInput",
			seq2:        nil,
			keySel:      func(_ int, s string) int { return len(s) },
			expectedErr: ErrNilInput,
		},
		{name: "NilSelector",
			seq2:        sec2_int_word(),
			keySel:      nil,
			expectedErr: ErrNilSelector,
		},
		{name: "DistinctBy",
			seq2:   sec2_int_word(),
			keySel: func(_ int, s string) int { return len(s) },
			want:   errorhelper.Must(iterhelper.Var2[int, string](0, "zero", 1, "one", 3, "three")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := DistinctBy(tt.seq2, tt.keySel)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("DistinctBy() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("DistinctBy() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("DistinctBy() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("DistinctBy(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}

func TestDistinctByEq(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		keySel      func(int, string) int
		keyEq       func(int, int) bool
		want        iter.Seq2[int, string]
		expectedErr error
	}{
		{name: "DistinctByEq",
			seq2:   sec2_int_word(),
			keySel: func(_ int, s string) int { return len(s) % 2 },
			keyEq:  func(a, b int) bool { return a == b },
			want:   errorhelper.Must(iterhelper.Var2[int, string](0, "zero", 1, "one")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := DistinctByEq(tt.seq2, tt.keySel, tt.keyEq)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("DistinctByEq() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("DistinctByEq() succeeded unexpectedly")
			}
			equal, _ := iterhelper.Equal2(got, tt.want)
			if !equal {
				t.Errorf("DistinctByEq(): %v, want: %v", iterhelper.StringDef2(got), iterhelper.StringDef2(tt.want))
			}
		})
	}
}
