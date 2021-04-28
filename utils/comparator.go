package utils

import (
	"reflect"
	"sort"
)

// Returns true if two 2D-string arrays are equals independent of the order
func Compare2DStringArray(a, b [][]string) bool {
	for _, arr := range a {
		sort.Strings(arr)
	}
	for _, arr := range b {
		sort.Strings(arr)
	}
	sort.Slice(a, func(i, j int) bool { return len(a[i]) < len(a[j]) })
	sort.Slice(b, func(i, j int) bool { return len(b[i]) < len(b[j]) })
	return reflect.DeepEqual(a, b)
}
