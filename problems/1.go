package problems

// 1. Two Sum
// https://leetcode.com/problems/two-sum/
//
// Solution using hash table
// Time complexity: O(n)
// Space complexity: O(n)
func TwoSum(nums []int, target int) []int {
	var myMap = map[int]int{}
	for i, num := range nums {
		if j, ok := myMap[target-num]; ok {
			return []int{j, i}
		}
		myMap[num] = i
	}
	return nil
}

// Brute force
// Time complexity: O(n^2)
// Space complexity: O(1)
func TwoSumBruteForce(nums []int, target int) []int {
	for i, a := range nums {
		for j, b := range nums[i+1:] {
			if a+b == target {
				return []int{i, j + i + 1}
			}
		}
	}
	return nil
}
