package PlusMinus

import "fmt"

func PlusMinus(arr []int32) {
	size := float32(len(arr))
	counter := make(map[string]float32)

	for _, val := range arr {

		if val > 0 {
			counter["+"] += 1
		} else if val < 0 {
			counter["-"] += 1
		} else {
			counter["0"] += 1
		}

	}

	fmt.Printf("%6f\n", counter["+"]/size)
	fmt.Printf("%6f\n", counter["-"]/size)
	fmt.Printf("%6f\n", counter["0"]/size)
}
