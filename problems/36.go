package problems

// 36. Valid Sudoku
// https://leetcode.com/problems/valid-sudoku/
//
// This is the naive approach that iterate over board three times.
// One could iterate only once and hold 3 maps to track collisions,
// but that would not change the asymptotic complexity.
//
// Time Complexity O(n) // where n is the size of board
// Space Complexity O(1)
func IsValidSudoku(board [][]byte) bool {
	// check row-wise
	for x := 0; x < 9; x++ {
		digits := make(map[byte]struct{})
		for y := 0; y < 9; y++ {
			cell := board[x][y]
			if cell != '.' {
				if _, ok := digits[cell]; ok {
					return false
				}
				digits[cell] = struct{}{}
			}
		}
	}
	// check column-wise
	for y := 0; y < 9; y++ {
		digits := make(map[byte]struct{})
		for x := 0; x < 9; x++ {
			cell := board[x][y]
			if cell != '.' {
				if _, ok := digits[cell]; ok {
					return false
				}
				digits[cell] = struct{}{}
			}
		}
	}
	// check grid-wise
	for h := 0; h < 3; h++ {
		for v := 0; v < 3; v++ {
			digits := make(map[byte]struct{})
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					cell := board[3*h+x][3*v+y]
					if cell != '.' {
						if _, ok := digits[cell]; ok {
							return false
						}
						digits[cell] = struct{}{}
					}
				}
			}

		}
	}
	return true
}
