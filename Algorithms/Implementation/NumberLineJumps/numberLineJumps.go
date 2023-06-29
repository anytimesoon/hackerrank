package NumberLineJumps

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if (x1 > x2 && v1 > v2) || (x1 < x2 && v1 < v2) {
		return "NO"
	}

	if x1 == x2 && v1 != v2 {
		return "NO"
	}

	var i int
	for i < 10000 {
		if x1 == x2 {
			return "YES"
		}

		x1 = x1 + v1
		x2 = x2 + v2
		i++
	}
	return "NO"
}
