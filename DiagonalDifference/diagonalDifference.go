package DiagonalDifference

func DiagonalDifference(arr [][]int32) int32 {
	j := len(arr) - 1
	i := 0
	primary := int32(0)
	secondary := int32(0)
	// iterate over first array with one marker at first i and one at last i
	for _, row := range arr {
		primary += row[i]
		secondary += row[j]
		i++
		j--
	}

	return sanitize(primary - secondary)
}

func sanitize(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
