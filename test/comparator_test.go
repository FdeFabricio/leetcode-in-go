package test

import (
	"testing"

	"github.com/FdeFabricio/leetcode-in-go/utils"
)

func TestCompareIntArrays(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want bool
	}{
		{name: "true - identical", a: []int{1, 1, 2}, b: []int{1, 1, 2}, want: true},
		{name: "true - empty", a: []int{}, b: []int{}, want: true},
		{name: "true - nil", a: nil, b: nil, want: true},

		{name: "false - not identical", a: []int{1, 1, 2}, b: []int{1, 2, 2}, want: false},
		{name: "false - different length", a: []int{1, 1, 2}, b: []int{1, 1}, want: false},
		{name: "false - empty", a: []int{1}, b: []int{}, want: false},
		{name: "false - nil", a: nil, b: []int{1}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.CompareIntArrays(tt.a, tt.b); got != tt.want {
				t.Errorf("CompareIntArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
