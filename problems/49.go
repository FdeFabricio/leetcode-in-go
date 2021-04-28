package problems

// 49. Group Anagrams
// https://leetcode.com/problems/group-anagrams/
//
// For each string, it encodes into an array of 26 positions with
// the count of each letter, and saved it to a map to group them.
// Another option would be to encode by sorting the characters of
// the string.
// Time Complexity:  O(nl)  // where n is the size of the array
// Space Complexity: O(nl)  // and l is the length of the string
func GroupAnagrams(strs []string) [][]string {
	myMap := make(map[[26]int][]string)
	for _, str := range strs {
		encodedString := encodeString(str)
		myMap[encodedString] = append(myMap[encodedString], str)
	}
	var result [][]string
	for _, v := range myMap {
		result = append(result, v)
	}
	return result
}

func encodeString(str string) [26]int {
	var result [26]int
	for _, char := range str {
		result[char-'a']++
	}
	return result
}
