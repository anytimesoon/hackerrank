package GradingStudents

func gradingStudents(grades []int32) []int32 {
	res := make([]int32, 0)

	for _, grade := range grades {
		if grade < 38 ||
			grade%5 == 0 ||
			(grade-1)%5 == 0 ||
			(grade-2)%5 == 0 {
			res = append(res, grade)
			continue
		}

		res = append(res, roundUp(grade+1))
	}
	return res
}

func roundUp(grade int32) int32 {
	if grade%5 == 0 {
		return grade
	}

	return roundUp(grade + 1)
}
