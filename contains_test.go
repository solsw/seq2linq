package seq2linq

import (
	"errors"
	"iter"
	"strings"
	"testing"

	"github.com/solsw/iterhelper"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		valK        int
		valV        string
		want        bool
		expectedErr error
	}{
		{name: "EmptyIn",
			seq2: iterhelper.Empty2[int, string](),
			valK: 1,
			valV: "one",
			want: false,
		},
		{name: "NoMatch",
			seq2: sec2_int_word(),
			valK: 3,
			valV: "one",
			want: false,
		},
		{name: "Match",
			seq2: sec2_int_word(),
			valK: 6,
			valV: "six",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Contains(tt.seq2, tt.valK, tt.valV)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("Contains() failed: %v", gotErr)
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("Contains() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsEq(t *testing.T) {
	tests := []struct {
		name        string
		seq2        iter.Seq2[int, string]
		valK        int
		valV        string
		eq          func(int, string, int, string) bool
		want        bool
		expectedErr error
	}{
		{name: "NilInThrows",
			seq2:        nil,
			eq:          func(int, string, int, string) bool { return true },
			expectedErr: ErrNilInput,
		},
		{name: "NilEqThrows",
			seq2:        sec2_int_word(),
			eq:          nil,
			expectedErr: ErrNilEqual,
		},
		{name: "EmptyIn",
			seq2: iterhelper.Empty2[int, string](),
			valK: 1,
			valV: "one",
			eq:   func(int, string, int, string) bool { return true },
			want: false,
		},
		{name: "NoMatch",
			seq2: sec2_int_word(),
			valK: 6,
			valV: "SIX",
			eq:   func(i1 int, s1 string, i2 int, s2 string) bool { return i1 == i2 && s1 == s2 },
			want: false,
		},
		{name: "Match",
			seq2: sec2_int_word(),
			valK: 12,
			valV: "SIX",
			eq: func(i1 int, s1 string, i2 int, s2 string) bool {
				return i1 == i2/2 && strings.EqualFold(s1, s2)
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := ContainsEq(tt.seq2, tt.valK, tt.valV, tt.eq)
			if gotErr != nil {
				if tt.expectedErr == nil {
					t.Errorf("ContainsEq() failed: %v", gotErr)
				} else {
					if !errors.Is(gotErr, tt.expectedErr) {
						t.Errorf("ContainsEq() error: %v, expected: %v", gotErr, tt.expectedErr)
					}
				}
				return
			}
			if tt.expectedErr != nil {
				t.Fatal("ContainsEq() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("ContainsEq() = %v, want %v", got, tt.want)
			}
		})
	}
}
