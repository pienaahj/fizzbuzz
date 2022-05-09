package fizzbuzz

import (
	"fmt"
)

func Run(ints []int) []string {
	strings := []string{}
	for _, i := range ints {
		strings = append(strings, fmt.Sprint(i))
	}
	return strings
}
