package test

import (
	"testing"

	"github.com/FdeFabricio/leetcode-in-go/problems"
	"github.com/FdeFabricio/leetcode-in-go/utils"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{
			name:  "original test 1",
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want:  [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{name: "original test 2", input: []string{""}, want: [][]string{{""}}},
		{name: "original test 3", input: []string{"a"}, want: [][]string{{"a"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problems.GroupAnagrams(tt.input); !utils.Compare2DStringArray(got, tt.want) {
				t.Errorf("GroupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
