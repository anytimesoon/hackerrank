package TimeConversion

import (
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	if strings.Contains(s, "AM") {
		if s[:2] == "12" {
			return "00" + s[2:len(s)-2]
		}

		return s[:len(s)-2]
	}

	i, err := strconv.Atoi(s[:2])
	if err != nil {
		panic(err)
	}

	if i != 12 {
		return strconv.Itoa(i+12) + s[2:len(s)-2]
	} else {
		return s[:len(s)-2]
	}
}
