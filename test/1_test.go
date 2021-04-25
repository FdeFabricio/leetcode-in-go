package test

import (
	"testing"

	"github.com/FdeFabricio/leetcode-in-go/problems"
	"github.com/FdeFabricio/leetcode-in-go/utils"
)

var tests = []struct {
	name   string
	nums   []int
	target int
	want   []int
}{
	{name: "original test 1", nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
	{name: "original test 2", nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
	{name: "original test 3", nums: []int{3, 3}, target: 6, want: []int{0, 1}},

	{name: "negative numbers", nums: []int{-3, -4, -2}, target: -5, want: []int{0, 2}},
	{name: "negative and positive numbers", nums: []int{-3, 3, -2}, target: 1, want: []int{1, 2}},

	{name: "no solution", nums: []int{-3, 3, -2}, target: 3, want: nil},
}

func TestTwoSum(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problems.TwoSum(tt.nums, tt.target); !utils.CompareIntArrays(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoSumBruteForce(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problems.TwoSumBruteForce(tt.nums, tt.target); !utils.CompareIntArrays(got, tt.want) {
				t.Errorf("TwoSumBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
