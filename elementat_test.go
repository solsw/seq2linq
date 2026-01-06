package seq2linq

import (
	"errors"
	"iter"
	"testing"

	"github.com/solsw/generichelper"
)

func TestElementAt(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		index       int
		wantK       int
		wantV       string
		expectedErr error
	}{
		{name: "IndexOutOfRange1",
			seq2:        sec2_int_word(),
			index:       -1,
			expectedErr: ErrIndexOutOfRange,
		},
		{name: "IndexOutOfRange2",
			seq2:        sec2_int_word(),
			index:       100,
			expectedErr: ErrIndexOutOfRange,
		},
		{name: "ElementAt",
			seq2:  sec2_int_word(),
			index: 1,
			wantK: 1,
			wantV: "one",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotV, gotErr := ElementAt(tt.seq2, tt.index)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("ElementAt() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("ElementAt() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("ElementAt() succeeded unexpectedly")
			}
			if (gotK != tt.wantK) || (gotV != tt.wantV) {
				t.Errorf("ElementAt(): %v, want: %v",
					generichelper.Tuple2[int, string]{Item1: gotK, Item2: gotV},
					generichelper.Tuple2[int, string]{Item1: tt.wantK, Item2: tt.wantV})
			}
		})
	}
}

func TestElementAtOrDefault(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		index       int
		wantK       int
		wantV       string
		expectedErr error
	}{
		{name: "IndexOutOfRange1",
			seq2:  sec2_int_word(),
			index: -1,
			wantK: generichelper.ZeroValue[int](),
			wantV: generichelper.ZeroValue[string](),
		},
		{name: "IndexOutOfRange2",
			seq2:  sec2_int_word(),
			index: 100,
			wantK: generichelper.ZeroValue[int](),
			wantV: generichelper.ZeroValue[string](),
		},
		{name: "ElementAtOrDefault",
			seq2:  sec2_int_word(),
			index: 2,
			wantK: 2,
			wantV: "two",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, gotV, gotErr := ElementAtOrDefault(tt.seq2, tt.index)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("ElementAtOrDefault() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("ElementAtOrDefault() succeeded unexpectedly")
			}
			if (gotK != tt.wantK) || (gotV != tt.wantV) {
				t.Errorf("ElementAtOrDefault(): %v, want: %v",
					generichelper.Tuple2[int, string]{Item1: gotK, Item2: gotV},
					generichelper.Tuple2[int, string]{Item1: tt.wantK, Item2: tt.wantV})
			}
		})
	}
}
