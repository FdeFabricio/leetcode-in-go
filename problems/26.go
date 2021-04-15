package problems

// 26. Remove Duplicates from Sorted Array
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
//
// To pass by reference we would have to change the function's signature. Therefore,
// to accomplish the goal without allocating extra space, we use a slice.
// More info: https://github.com/golang/go/wiki/SliceTricks#filtering-without-allocating
//
// Time complexity: O(n)
// Space complexity: O(1)
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	resultSlice := nums[:1]
	for _, num := range nums[1:] {
		if num != resultSlice[len(resultSlice)-1] {
			resultSlice = append(resultSlice, num)
		}
	}

	return len(resultSlice)
}
