package seq2linq

import (
	"fmt"
	"iter"
)

func sec2_int_word() iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		if !yield(0, "zero") {
			return
		}
		if !yield(1, "one") {
			return
		}
		if !yield(2, "two") {
			return
		}
		if !yield(3, "three") {
			return
		}
		if !yield(4, "four") {
			return
		}
		if !yield(5, "five") {
			return
		}
		if !yield(6, "six") {
			return
		}
		if !yield(7, "seven") {
			return
		}
		if !yield(8, "eight") {
			return
		}
		if !yield(9, "nine") {
			return
		}
	}
}

func sec2_int_string(n int) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for i := range n {
			if !yield(i, fmt.Sprint(i)) {
				return
			}
		}
	}
}

func infinite_int_string() iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		i := 0
		for {
			if !yield(i, fmt.Sprint(i)) {
				return
			}
			i++
		}
	}
}
