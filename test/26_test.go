package test

import (
	"testing"

	"github.com/FdeFabricio/leetcode-in-go/problems"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "original test 1", input: []int{1, 1, 2}, want: 2},
		{name: "original test 2", input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, want: 5},

		{name: "no duplicates", input: []int{2, 3, 4, 5}, want: 4},
		{name: "duplicate in the end", input: []int{2, 3, 5, 5}, want: 3},
		{name: "empty list", input: []int{}, want: 0},
		{name: "one item list", input: []int{1}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problems.RemoveDuplicates(tt.input); got != tt.want {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}

}
