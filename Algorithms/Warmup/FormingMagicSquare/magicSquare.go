package FormingMagicSquare

import (
	"math"
)

// list of all possible magic squares
var posSqs = [][][]int32{
	{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
	{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
	{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
	{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
	{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
	{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
	{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
	{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
}

func min(x, y int32) int32 {
	if x > y {
		return y
	}
	return x
}

// sanitize turns a negative number positive
func sanitize(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func FormingMagicSquare(s [][]int32) int32 {
	res := int32(math.MaxInt32)

	// iterate over the list of squares
	for _, sq := range posSqs {
		cost := int32(0)
		// iterate of over the first row
		for i, row := range sq {
			// iterate over each integer
			for j := range row {
				cost += sanitize(row[j] - s[i][j])
			}
		}
		res = min(res, cost)
	}

	return res
}
