package test

import (
	"testing"

	"github.com/FdeFabricio/leetcode-in-go/utils"
)

func TestCompare2DStringArray(t *testing.T) {
	tests := []struct {
		name string
		a    [][]string
		b    [][]string
		want bool
	}{
		{
			name: "true - identical",
			a:    [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
			b:    [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
			want: true,
		},
		{
			name: "true - different order",
			a:    [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
			b:    [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
			want: true,
		},
		{
			name: "false - not equal",
			a:    [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
			b:    [][]string{{"eat", "tea", "ate"}, {"tan"}, {"nat", "bat"}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.Compare2DStringArray(tt.a, tt.b); got != tt.want {
				t.Errorf("Compare2DStringArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
