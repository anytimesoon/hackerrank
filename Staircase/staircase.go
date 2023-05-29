package Staircase

import (
	"fmt"
	"strings"
)

func Staircase(n int32) {
	hashCounter := 1
	nAsInt := int(n)
	for nAsInt >= hashCounter {
		fmt.Println(strings.Repeat(" ", nAsInt-hashCounter) + strings.Repeat("#", hashCounter))
		hashCounter++
	}
}
