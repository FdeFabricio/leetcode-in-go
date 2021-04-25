package utils

// Returns true if two int arrays are identical
func CompareIntArrays(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, element := range a {
		if element != b[i] {
			return false
		}
	}

	return true
}
