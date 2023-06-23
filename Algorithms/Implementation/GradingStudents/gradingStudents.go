package GradingStudents

import "strconv"

func gradingStudents(grades []int32) []int32 {
	res := make([]int32, 0)
	m := map[string]bool{
		"3": true,
		"4": true,
		"8": true,
		"9": true,
	}

	for _, grade := range grades {
		iString := strconv.Itoa(int(grade))
		if grade < 38 || !m[iString[1:]] {
			res = append(res, grade)
			continue
		}

		if iString[1:] == "3" || iString[1:] == "8" {
			res = append(res, grade+2)
		} else {
			res = append(res, grade+1)
		}
	}
	return res
}
