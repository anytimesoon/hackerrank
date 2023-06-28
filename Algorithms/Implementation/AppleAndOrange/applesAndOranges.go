package AppleAndOrange

func CountApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) []int32 {
	var applesInRange int32
	var orangesInRange int32

	for _, throw := range apples {
		if checkIfInRange(s, t, a+throw) {
			applesInRange++
		}
	}

	for _, throw := range oranges {
		if checkIfInRange(s, t, b+throw) {
			orangesInRange++
		}
	}

	return []int32{applesInRange, orangesInRange}
}

func checkIfInRange(s int32, t int32, dist int32) bool {
	if dist >= s && dist <= t {
		return true
	}

	return false
}
